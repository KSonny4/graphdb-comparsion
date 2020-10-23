package main

import (
	"fmt"
	"log"
	"strings"
	"context"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type File struct {
	Name string `json:"_key"`
	FileType string `json:"filetype"`
}

type MyEdgeObject struct {
	From string `json:"_from"`
	To   string `json:"_to"`
}

func main() {
	fmt.Println("Hello World")

	// Create an HTTP connection to the database
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}
	// Create a client
	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})

	// Create database
	db, err := c.CreateDatabase(nil, "my_graph_db", nil)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//FIXME: I did not find a way how to create Graph without any of this garbage!!!	
	// define the edgeCollection to store the edges
	var edgeDefinition driver.EdgeDefinition
	edgeDefinition.Collection = "myEdgeCollection"
	// define a set of collections where an edge is going out...
	edgeDefinition.From = []string{"myCollection1"}

	// repeat this for the collections where an edge is going into
	edgeDefinition.To = []string{"myCollection1"}

	// A graph can contain additional vertex collections, defined in the set of orphan collections
	var options driver.CreateGraphOptions
	options.EdgeDefinitions = []driver.EdgeDefinition{edgeDefinition}
	//FIXME: I did not find a way how to create Graph without any of this garbage!!!
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	// now it's possible to create a graph
	graph, err := db.CreateGraph(nil, "myGraph", &options)
	if err != nil {
		log.Fatalf("Failed to create graph: %v", err)
	}

	// add vertex
	vertexCollection1, err := graph.VertexCollection(nil, "myCollection1")
	if err != nil {
		log.Fatalf("Failed to get vertex collection: %v", err)
	}

	myFiles := []File{
		File{
			"file1",
			"zip",
		},
		File{
			"file2",
			"zip",
		},
		File{
			"file3",
			"zip",
		},
	}

	// Add vertices to collection
	metas, _, err := vertexCollection1.CreateDocuments(nil, myFiles)
	if err != nil {
		log.Fatalf("Failed to create vertex documents: %v", err)
	}


	// Print vertices
	fmt.Printf("Created documents with keys '%s' in collection '%s' in database '%s'\n", strings.Join(metas.Keys(), ","), vertexCollection1.Name(), db.Name())


	// Print edges
	ctx := context.Background()
	//FIXME I dont belive this is correct, but the documentation is terrible
	query := "RETURN DOCUMENT('myCollection1/file1')"
	
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
    	log.Fatalf("Query for Print edges failed: %v", err) 
	}

	
	defer cursor.Close()
	// FIXME how to get output?
	fmt.Println(cursor)

	// Create edges between files
	// TODO https://www.arangodb.com/docs/stable/aql/tutorial-traversal.html#creating-the-edges

	// FOR rel in data
    // LET parentId = FIRST(
    //     FOR c IN Characters
    //         FILTER c.name == rel.parent.name
    //         FILTER c.surname == rel.parent.surname
    //         LIMIT 1
    //         RETURN c._id
    // )
    // LET childId = FIRST(
    //     FOR c IN Characters
    //         FILTER c.name == rel.child.name
    //         FILTER c.surname == rel.child.surname
    //         LIMIT 1
    //         RETURN c._id
    // )
    // FILTER parentId != null AND childId != null
    // INSERT { _from: childId, _to: parentId } INTO ChildOf
    // RETURN NEW

	// Print vertices connected with edge


	// delete graph
	graph.Remove(nil)
}

