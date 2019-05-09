package server

import (
	"fmt"

	"github.com/bmeg/grip/gdbi"
	"github.com/bmeg/grip/gripql"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

func buildField(x string) (*graphql.Field, error) {
	var o *graphql.Field
	switch x {
	case "NUMERIC":
		o = &graphql.Field{Type: graphql.Float}
	case "STRING":
		o = &graphql.Field{Type: graphql.String}
	case "BOOL":
		o = &graphql.Field{Type: graphql.Boolean}
	default:
		return nil, fmt.Errorf("%s does not map to a GraphQL type", x)
	}
	return o, nil
}

func buildSliceField(name string, s []interface{}) (*graphql.Field, error) {
	var f *graphql.Field
	var err error

	if len(s) > 0 {
		val := s[0]

		if x, ok := val.(map[string]interface{}); ok {
			f, err = buildObjectField(name, x)

		} else if x, ok := val.([]interface{}); ok {
			f, err = buildSliceField(name, x)

		} else if x, ok := val.(string); ok {
			f, err = buildField(x)

		} else {
			err = fmt.Errorf("unhandled type: %T %v", val, val)
		}

	} else {
		err = fmt.Errorf("slice is empty")
	}

	if err != nil {
		return nil, fmt.Errorf("buildSliceField error: %v", err)
	}

	return &graphql.Field{Type: graphql.NewList(f.Type)}, nil
}

func buildObjectField(name string, obj map[string]interface{}) (*graphql.Field, error) {
	o, err := buildObject(name, obj)
	if err != nil {
		return nil, err
	}

	return &graphql.Field{Type: o}, nil
}

func buildObject(name string, obj map[string]interface{}) (*graphql.Object, error) {
	objFields := graphql.Fields{}

	for key, val := range obj {
		var err error

		// handle map
		if x, ok := val.(map[string]interface{}); ok {
			objFields[key], err = buildObjectField(key, x)

			// handle slice
		} else if x, ok := val.([]interface{}); ok {
			objFields[key], err = buildSliceField(key, x)

			// handle string
		} else if x, ok := val.(string); ok {
			objFields[key], err = buildField(x)

			// handle other cases
		} else {
			err = fmt.Errorf("unhandled type: %T %v", val, val)
		}

		if err != nil {
			log.WithFields(log.Fields{"object": name, "field": key, "error": err}).Error("graphql: buildObject")
			// return nil, fmt.Errorf("object: %s: field: %s: error: %v", name, key, err)
		}
	}

	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   name,
			Fields: objFields,
		},
	), nil
}

func buildFieldConfigArument(obj *graphql.Object) graphql.FieldConfigArgument {
	args := graphql.FieldConfigArgument{
		"id":     &graphql.ArgumentConfig{Type: graphql.String},
		"ids":    &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
		"limit":  &graphql.ArgumentConfig{Type: graphql.Int},
		"offset": &graphql.ArgumentConfig{Type: graphql.Int},
	}
	if obj == nil {
		return args
	}
	for k, v := range obj.Fields() {
		switch v.Type {
		case graphql.String, graphql.Int, graphql.Float, graphql.Boolean:
			args[k] = &graphql.ArgumentConfig{Type: v.Type}
		default:
			continue
		}
	}
	return args
}

func buildEdgeField(objectMap map[string]*graphql.Object, edgeMap map[string]map[string]*graphql.Object, e *gripql.Edge) *graphql.Field {
	objs := []*graphql.Object{}
	for _, o := range edgeMap[e.Label] {
		objs = append(objs, o)
	}
	if len(objs) == 0 {
		return nil
	} else if len(objs) == 1 {
		o := objs[0]
		return &graphql.Field{
			Name: e.To,
			Type: graphql.NewList(o),
			Args: buildFieldConfigArument(o),
		}
	} else {
		return &graphql.Field{
			Name: e.To,
			Type: graphql.NewList(graphql.NewUnion(graphql.UnionConfig{
				Name:  fmt.Sprintf("%s_to_%s", e.From, e.Label),
				Types: objs,
				ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
					if mapVal, ok := p.Value.(map[string]interface{}); ok {
						vLabel := mapVal["label"]
						return objectMap[vLabel.(string)]
					}
					return nil
				},
			})),
			Args: buildFieldConfigArument(nil),
		}
	}
}

func buildGraphQLSchema(db gdbi.GraphDB, workdir string, graph string, schema *gripql.Graph) (*graphql.Schema, error) {
	if schema == nil {
		return nil, fmt.Errorf("graphql.NewSchema error: nil gripql.Graph for graph: %s", graph)
	}

	objects := map[string]*graphql.Object{}

	for _, obj := range schema.Vertices {
		props := obj.GetDataMap()
		if props == nil {
			props = make(map[string]interface{})
		}
		props["id"] = "STRING"
		props["label"] = "STRING"
		gqlObj, err := buildObject(obj.Label, props)
		if err != nil {
			return nil, err
		}
		objects[obj.Label] = gqlObj
	}

	// Setup outgoing edge fields
	// Note: edge properties are not accessible in this model
	edgeLabels := []string{}
	polyEdgeTo := make(map[string]map[string]*graphql.Object)
	for _, e := range schema.Edges {
		edgeLabels = append(edgeLabels, e.Label)
		to := objects[e.To]
		if _, ok := polyEdgeTo[e.Label]; ok {
			polyEdgeTo[e.Label][to.Name()] = to
		} else {
			polyEdgeTo[e.Label] = map[string]*graphql.Object{to.Name(): to}
		}
	}

	for _, e := range schema.Edges {
		f := buildEdgeField(objects, polyEdgeTo, e)
		if f != nil {
			objects[e.From].AddFieldConfig(e.Label, f)
		}
	}

	resolverConf := &GraphQLResolverConfig{
		DB:         db,
		WorkDir:    workdir,
		Graph:      graph,
		EdgeLabels: edgeLabels,
	}

	queryFields := graphql.Fields{}
	for objName, obj := range objects {
		label := obj.Name()
		args := buildFieldConfigArument(obj)
		f := &graphql.Field{
			Name: objName,
			Type: graphql.NewList(obj),
			Args: args,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return ResolveGraphQL(resolverConf, label, params)
			},
		}
		queryFields[objName] = f
	}

	query := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: queryFields,
		},
	)

	schemaConfig := graphql.SchemaConfig{
		Query: query,
	}

	gqlSchema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, fmt.Errorf("graphql.NewSchema error: %v", err)
	}

	return &gqlSchema, nil
}
