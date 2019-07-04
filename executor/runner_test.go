
package executor

import (
	"log"
	"sync"
	"testing"
	"encoding/json"
)


func TestPythonRunner(t *testing.T) {
	exec, err := StartLocalExecutor("./grip-exec-python/grip-exec.py")
	if err != nil {
		t.Errorf("error: %s", err)
	}
	log.Printf("Port %d", exec.Port)
	code := &Code{ Code: `
def add(x):
	return x + 1
`,
		Function: "add",
	}
	codeRes, err := exec.Compile(code)
	log.Printf("%#v", codeRes)

	in := make(chan *Input, 10)
	out, err := exec.Process(in)
	if err != nil {
		t.Error(err)
	}
	w := &sync.WaitGroup{}
	w.Add(1)
	go func() {
		count := 0
		for o := range out {
			log.Printf("%#v", o)
			count++
		}
		if count != 3 {
			t.Error("Wrong return count")
		}
		w.Done()
	}()
	s1, _ := json.Marshal([]int{1})
	in <- &Input{Code:codeRes.Id, Data:string(s1)}
	s2, _ := json.Marshal([]int{2})
	in <- &Input{Code:codeRes.Id, Data:string(s2)}
	s3, _ := json.Marshal([]int{3})
	in <- &Input{Code:codeRes.Id, Data:string(s3)}
	close(in)

	w.Wait()
	exec.Close()
}
