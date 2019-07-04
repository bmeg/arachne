
package executor

import (
  //"os"
  "io"
  "log"
  "fmt"
  //"bytes"
  "strings"
  "strconv"
  "context"
  "os/exec"
  grpc "google.golang.org/grpc"
)

type Runner struct {
  Port int
  Cmd *exec.Cmd
  Conn *grpc.ClientConn
  Client ExecutorClient
}

func StartLocalExecutor(path string) (Runner, error) {
  cmd := exec.Command(path)

  stdout, _ := cmd.StdoutPipe()
  err := cmd.Start()
  if err != nil {
    return Runner{}, err
  }

  m := make(chan int)
  err = nil

  go func() error {
    sent := false
    defer close(m)
    var out []byte
    buf := make([]byte, 1024, 1024)
    for {
      n, ierr := stdout.Read(buf)
      if !sent {
        if n > 0 {
          d := buf[:n]
          out = append(out, d...)
        }
        log.Printf("Read %d (%s) Buffer: %s", n, ierr, string(out))
        if strings.Contains(string(out), "\n") {
          t := strings.Split(string(out), "\n")
          log.Printf("Return port: %s", out)
          p, ierr := strconv.Atoi(string(t[0]))
          err = ierr
          m <- p
          sent = true
        }
      }
      if ierr != nil {
        if ierr == io.EOF {
          log.Printf("Executor closed")
          if cmd.ProcessState.ExitCode() != 0 {
            log.Printf("Executor error: %d", cmd.ProcessState.ExitCode())
          }
          ierr = nil
        }
        return ierr
      }
    }
  }()
  var port int
  port = <- m
  serverAddr := fmt.Sprintf("localhost:%d", port)
  conn, err := grpc.Dial(serverAddr,  grpc.WithInsecure())
  if err != nil {
    return Runner{}, err
  }
  client := NewExecutorClient(conn)
  return Runner{Port:port, Cmd:cmd, Conn:conn, Client:client}, err
}

func (run *Runner) Compile(code *Code) (*CompileResult, error) {
  return run.Client.Compile(context.Background(), code)
}

func (run *Runner) Process(in chan *Input) (chan *Result, error) {
  cl, err := run.Client.Process(context.Background())
  if err != nil {
    return nil, err
  }
  out := make(chan *Result, 100)
  go func() {
    for i := range in {
      cl.Send(i)
    }
    cl.CloseSend()
  }()
  go func() {
    for {
		    r, err := cl.Recv()
		    if err != nil {
          return
        }
        out <- r
    }
  }()
  return out, nil
}

func (run *Runner) Close() {
  run.Cmd.Process.Kill()
  run.Cmd.Wait()
}
