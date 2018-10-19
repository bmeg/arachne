package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/kvgraph"
	_ "github.com/bmeg/grip/kvi/badgerdb" // import so badger will register itself
	"github.com/bmeg/grip/util"
	"github.com/bmeg/grip/util/rpc"
)

func TestBasicAuthFail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := testConfig()
	conf.BasicAuth = []BasicCredential{{User: "testuser", Password: "abc123"}}
	defer os.RemoveAll(conf.WorkDir)

	srv, err := NewGripServer(nil, conf)
	if err != nil {
		t.Fatal(err)
	}

	go srv.Serve(ctx)

	cli, err := gripql.Connect(rpc.Config{ServerAddress: conf.RPCAddress(), Timeout: 5 * time.Second}, true)
	if err != nil {
		t.Fatal(err)
	}

	_, err = cli.Traversal(&gripql.GraphQuery{Graph: "test", Query: gripql.NewQuery().V().Statements})
	if err == nil || !strings.Contains(err.Error(), "PermissionDenied") {
		t.Error("expected error")
	}

	_, err = cli.ListGraphs()
	if err == nil || !strings.Contains(err.Error(), "PermissionDenied") {
		t.Error("expected error")
	}

	_, err = cli.GetVertex("test", "1")
	if err == nil || !strings.Contains(err.Error(), "PermissionDenied") {
		t.Error("expected error")
	}

	resp, _ := http.Get(fmt.Sprintf("http://localhost:%s/v1/graph", conf.HTTPPort))
	if resp.StatusCode != 401 {
		t.Error("expected http 401 error")
	}
}

func TestBasicAuth(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := testConfig()
	conf.BasicAuth = []BasicCredential{{User: "testuser", Password: "abc123"}}
	defer os.RemoveAll(conf.WorkDir)

	os.Setenv("GRIP_USER", "testuser")
	os.Setenv("GRIP_PASSWORD", "abc123")
	defer os.Unsetenv("GRIP_USER")
	defer os.Unsetenv("GRIP_PASSWORD")

	tmpDB := "grip.db." + util.RandomString(6)
	db, err := kvgraph.NewKVGraphDB("badger", tmpDB)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDB)

	srv, err := NewGripServer(db, conf)
	if err != nil {
		t.Fatal(err)
	}

	go srv.Serve(ctx)

	cli, err := gripql.Connect(rpc.ConfigWithDefaults(conf.RPCAddress()), true)
	if err != nil {
		t.Fatal(err)
	}

	err = cli.AddGraph("test")
	if err != nil {
		t.Fatal(err)
	}

	err = cli.AddVertex("test", &gripql.Vertex{Gid: "1", Label: "test"})
	if err != nil {
		t.Fatal(err)
	}

	_, err = cli.Traversal(&gripql.GraphQuery{Graph: "test", Query: gripql.NewQuery().V().Statements})
	if err != nil {
		t.Error("unexpected error", err)
	}

	_, err = cli.ListGraphs()
	if err != nil {
		t.Error("unexpected error", err)
	}

	_, err = cli.GetVertex("test", "1")
	if err != nil {
		t.Error("unexpected error", err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%s/v1/graph", conf.HTTPPort), nil)
	req.SetBasicAuth("testuser", "abc123")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		t.Error("basic auth error error")
	}
	returnString := `{"result":{"graph":"test"}}
`
	bodyText, err := ioutil.ReadAll(resp.Body)
	if string(bodyText) != returnString {
		t.Error("incorrect http return value")
	}
}
