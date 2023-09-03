package main

import "graph_with_edges/graph"

func main() {
	g := graph.New()

	for i := 0; i < 8; i++ {
		g.AddVertex(i)
	}

	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 5, 3)
	g.AddEdge(0, 7, 2)
	g.AddEdge(1, 5, 1)
	g.AddEdge(1, 4, 8)
	g.AddEdge(4, 5, 4)
	g.AddEdge(2, 3, 6)
	g.AddEdge(2, 6, 3)
	g.AddEdge(3, 6, 4)
	g.AddEdge(3, 4, 2)
	g.AddEdge(4, 6, 2)

	graph.DFS(&g, 0)

	pr := graph.Primes(&g)
	graph.DFS(pr, 0)

}
