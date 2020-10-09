from pprint import pprint
from gremlin_python import statics
from gremlin_python.structure.graph import Graph
from gremlin_python.process.graph_traversal import __
from gremlin_python.process.strategies import *
from gremlin_python.driver.driver_remote_connection import DriverRemoteConnection
from gremlin_python.process.traversal import Cardinality

statics.load_statics(globals())
# Initialize the graph and the traversal
graph = Graph()
g = graph.traversal().withRemote(DriverRemoteConnection('ws://localhost:8182/gremlin','g'))

# Remove all
print('Removing the nodes...')
g.V().drop().iterate()

# print count of Nodes and Edges
print('Number of nodes {}, number of edges {}.'.format(g.V().count().next(),g.E().count().next()))

# Create nodes
print('Writing two nodes...')
node1 = (g.addV('label1').property('file_type','zip')
		.property('description','First node')
		.property('Node number',1))
node1.next()

node2 = (g.addV('label2').property('file_type','zip')
		 .property('description','Second node')
		 .property('Node number',2))
node2.next()

node2 = (g.addV('label3').property('file_type','zip')
		 .property('description','Third node')
		 .property('Node number',3))
node2.next()

print('Number of nodes {}, number of edges {}.'.format(g.V().count().next(),g.E().count().next()))

nodes_data = g.V().valueMap().toList()
pprint(f"Nodes in database: {nodes_data}")

# create edges between nodes with same "zip" property
g.V().has("file_type","zip").as_("a").V().has("file_type","zip").as_("b").where("a", eq("b")).by("file_type").addE("link").from_("a").to("b").next()

# Showing data on the edge:
edges_data = g.E().toList()
print('Edge and the data associated:')
print(edges_data)

print('Number of nodes {}, number of edges {}.'.format(g.V().count().next(),g.E().count().next()))

# Remove all
print('Removing the nodes...')
g.V().drop().iterate()