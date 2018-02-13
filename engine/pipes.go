package engine

type inPipe <-chan *traveler
type outPipe chan<- *traveler

type traveler struct {
  id string
  label string
  from, to string
  data map[string]interface{}
  marks map[string]*traveler
  count int64
  groupCounts map[string]int64
  row []*traveler
  value interface{}
  dataType
}

type dataType uint8
const (
  noData dataType = iota
  vertexData
  edgeData
  countData
  groupCountData
  valueData
  rowData
)

func run(procs []processor, in inPipe, out outPipe, bufsize int) {
  if len(procs) == 0 {
    close(out)
    return
  }
  if len(procs) == 1 {
    procs[0].process(in, out)
    close(out)
    return
  }

  for i := 0; i < len(procs) - 1; i++ {
    glue := make(chan *traveler, bufsize)
    go func(i int, in inPipe, out outPipe) {
      procs[i].process(in, out)
      close(out)
    }(i, in, glue)
    in = glue
  }
  procs[len(procs)-1].process(in, out)
  close(out)
}

/*

func mapPipe() {
	mfunc, err := jsengine.NewJSEngine(source, pengine.imports)
	if err != nil {
		log.Printf("Script Error: %s", err)
	}

	for i := range pipe.Travelers {
		out := mfunc.Call(i.GetCurrent())
		if out != nil {
			a := i.AddCurrent(*out)
			o <- a
		}
	}
}

func foldPipe() {
	mfunc, err := jsengine.NewJSEngine(source, pengine.imports)
	if err != nil {
		log.Printf("Script Error: %s", err)
	}

	var last *aql.QueryResult
	first := true
	for i := range pipe.Travelers {
		if first {
			last = i.GetCurrent()
			first = false
		} else {
			last = mfunc.Call(last, i.GetCurrent())
		}
	}
	if last != nil {
		i := Traveler{}
		a := i.AddCurrent(*last)
		o <- a
	}
}

func filterPipe() {
	mfunc, err := jsengine.NewJSEngine(source, pengine.imports)
	if err != nil {
		log.Printf("Script Error: %s", err)
	}
	for i := range pipe.Travelers {
		out := mfunc.CallBool(i.GetCurrent())
		if out {
			o <- i
		}
	}
}

func filterValuesPipe() {
  // TODO only create JS engine once?
	mfunc, err := jsengine.NewJSEngine(source, pengine.imports)
	if err != nil {
		log.Printf("Script Error: %s", err)
	}
	for i := range pipe.Travelers {
		out := mfunc.CallValueMapBool(i.State)
		if out {
			o <- i
		}
	}
}

func vertexFromValuesPipe() {
	mfunc, err := jsengine.NewJSEngine(source, pengine.imports)
	if err != nil {
		log.Printf("Script Error: %s", err)
	}
	for i := range pipe.Travelers {

		t.startTimer("javascript")
		out := mfunc.CallValueToVertex(i.State)
		t.endTimer("javascript")

		for _, j := range out {
			v := db.GetVertex(j, load)
			if v != nil {
				o <- i.AddCurrent(aql.QueryResult{Result: &aql.QueryResult_Vertex{Vertex: v}})
			}
		}
	}
}

*/

func contains(a []string, v string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}
