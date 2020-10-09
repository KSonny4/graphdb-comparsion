# Graph database comparison

Our goal is to compare graph databases and their ability to create relations between nodes since Dgraph is not capable of doing it. The implemented solutions should be in Golang.

Each database should be tested as follows:

1. Have own folder with docker-compose.yml making it easy to run the solution
2. Golang script that 
    1. Add nodes
    2. Make a request that should link nodes with some relation (e.g. create edges between nodes where attribute x = y)

Databases of our interest:

* [] [dgraph](dgraph)
* [] [arangodb](arangodb)
* [] [orientdb](orientdb)
* [] [Amazon Neptune](amazon_neptune)
* [] Azure CosmosDB
* [] ~flockDB~ (unmaintained now)
* [] cayley
* [] titan
* [] neo4j


## what is tincker-pop stack (gremlin)