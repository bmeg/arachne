
package executor

import (
	"log"
	"sync"
	"testing"
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

	in <- &Input{Code:codeRes.Id}
	in <- &Input{Code:codeRes.Id}
	in <- &Input{Code:codeRes.Id}
	close(in)

	w.Wait()
	exec.Close()
}
