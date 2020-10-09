package main

import (
	"fmt"
	"github.com/go-gremlin/gremlin"
	"log"
)

func main() {
	if err := gremlin.NewCluster("ws://localhost:8182/gremlin"); err != nil {
		log.Fatal(err)
	}

	data, err := gremlin.Query(`g.V()`).Exec()
	if err != nil  {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}