package example

import (
	"fmt"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/gripql/example"

	"github.com/bmeg/grip/log"
	"github.com/bmeg/grip/util/rpc"
	"github.com/spf13/cobra"
)

var host = "localhost:8202"
var graph = "example-graph"
var exampleSet = "starwars"

func found(set []string, val string) bool {
	for _, i := range set {
		if i == val {
			return true
		}
	}
	return false
}

// Cmd is the example loader command line definition
var Cmd = &cobra.Command{
	Use:   "example-graph",
	Short: "Load an example graph",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := gripql.Connect(rpc.ConfigWithDefaults(host), true)
		if err != nil {
			return err
		}

		resp, err := conn.ListGraphs()
		if err != nil {
			return err
		}
		if found(resp.Graphs, graph) {
			return fmt.Errorf("grip already contains a graph called %s", graph)
		}

		err = conn.AddGraph(graph)
		if err != nil {
			return err
		}

		elemChan := make(chan *gripql.GraphElement)
		wait := make(chan bool)
		go func() {
			if err := conn.BulkAdd(elemChan); err != nil {
				log.WithFields(log.Fields{"error": err}).Error("bulk add error")
			}
			wait <- false
		}()

		log.WithFields(log.Fields{"graph": graph}).Info("Loading example data into graph")
		for _, v := range example.SWVertices {
			elemChan <- &gripql.GraphElement{Graph: graph, Vertex: v}
		}
		for _, e := range example.SWEdges {
			elemChan <- &gripql.GraphElement{Graph: graph, Edge: e}
		}

		close(elemChan)
		<-wait

		return nil
	},
}

func init() {
	flags := Cmd.Flags()
	flags.StringVar(&host, "host", host, "grip server url")
}
