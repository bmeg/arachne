package tabular

import (
  "log"
  "fmt"
  "context"
  "path/filepath"
  "github.com/bmeg/grip/gdbi"
  "github.com/bmeg/grip/gripql"
)

type TabularGDB struct {
  graph *TabularGraph
}

func NewGDB(conf *GraphConfig, indexPath string) *TabularGDB {
  out := TabularGraph{}
  out.idx, _ = NewTablularIndex(indexPath)
  out.vertices = []*Table{}

  for _, t := range conf.Tables {
    log.Printf("Table: %s", t)
    fPath := filepath.Join( filepath.Dir(conf.path), t.Path )
    log.Printf("Loading: %s with primaryKey %s", fPath, t.PrimaryKey)
    tix := out.idx.IndexTSV(fPath, t.PrimaryKey, []string{})
    if t.Label != "" {
      out.vertices = append(out.vertices, &Table{data:tix, prefix:t.Prefix, label:t.Label})
    }
  }
  return &TabularGDB{&out}
}


func (g *TabularGDB) AddGraph(string) error {
  return fmt.Errorf("AddGraph not implemented")
}

func (g *TabularGDB) DeleteGraph(string) error {
  return fmt.Errorf("AddGraph not implemented")
}


func (g *TabularGDB) ListGraphs() []string {
  return []string{"main"}
}

func (g *TabularGDB) Graph(graphID string) (gdbi.GraphInterface, error) {
  return g.graph, nil
}

func (g *TabularGDB) BuildSchema(ctx context.Context, graphID string, sampleN uint32, random bool) (*gripql.Graph, error) {
  return nil, fmt.Errorf("AddGraph not implemented")
}

func (g *TabularGDB) Close() error {
    return g.graph.Close()
}
