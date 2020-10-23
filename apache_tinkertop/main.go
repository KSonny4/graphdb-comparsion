package main

import (
	"github.com/northwesternmutual/grammes/examples/exampleutil"
	"go.uber.org/zap"
	"log"

	"github.com/northwesternmutual/grammes"
)
func print_vertices_count(client *grammes.Client, logger *zap.Logger ) {
	// Count the vertices on the graph.
	count, err := client.VertexCount()
	if err != nil {
		logger.Fatal("Couldn't count vertices", zap.Error(err))
	}

	// Print out the number of vertices on the graph.
	// This should be 1.
	logger.Info("Counted Vertices", zap.Int64("count", count))
}

func call_query_string(client *grammes.Client, logger *zap.Logger, query string) {
	responses, err := client.ExecuteStringQuery(query)
	if err != nil {
		logger.Fatal("Error querying server", zap.Error(err))
	}

	for _, res := range responses {
		logger.Info("executed string query: [" + query + "]", zap.ByteString("result", res))
	}
}


func main() {
	logger := exampleutil.SetupLogger()
	defer logger.Sync()

	// Creates a new client with the localhost IP.
	client, err := grammes.DialWithWebSocket("ws://127.0.0.1:8182")
	if err != nil {
		log.Fatalf("Error while creating client: %s\n", err.Error())
	}

	// Drop all vertices on the graph currently.
	client.DropAll()

	// Drop the testing vertices when finished.
	defer client.DropAll()

	// Create a new traversal string to build your traverser.
	g := grammes.Traversal()

	// Create 3 vertices
	client.AddVertexByQuery(g.AddV("label1").Property("file_type", "zip"))
	client.AddVertexByQuery(g.AddV("label2").Property("file_type", "zip"))
	client.AddVertexByQuery(g.AddV("label3").Property("file_type", "zip"))

	print_vertices_count(client, logger)
	// print vertices
	call_query_string(client,logger,"g.V()")

	//create edges between vertices with same property file_type
	query_string := "g.V().has('file_type','zip').as('a').V().has('file_type','zip').as('b').where('a', eq('b')).by('file_type').addE('link').from('a').to('b')"
	call_query_string(client,logger,query_string)

	// print edges
	call_query_string(client,logger,"g.E()")

	print_vertices_count(client, logger)

}


