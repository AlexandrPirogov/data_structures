package main

import "graph_with_edges/graph"

func main() {
	g := graph.New()

	for i := 0; i < 7; i++ {
		g.AddVertex(i)
	}

	g.AddEdge(0, 2, 3)
	g.AddEdge(0, 3, 4)
	g.AddEdge(1, 0, 4)
	g.AddEdge(1, 3, 8)
	g.AddEdge(2, 3, 7)
	g.AddEdge(2, 5, -12)
	g.AddEdge(3, 4, 0)
	g.AddEdge(3, 5, 5)
	g.AddEdge(4, 1, 10)
	g.AddEdge(4, 6, 3)
	g.AddEdge(5, 4, 1)
	g.AddEdge(5, 6, -2)

	g.Dijkstra(0)

}
