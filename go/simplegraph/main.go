package main

import (
	"simplegraph/graph"
)

func main() {
	g := graph.New[int]()
	for i := 9; i >= 0; i-- {
		g.AddVertex(i)
	}

	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 3)

	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	g.AddEdge(8, 9)
	g.AddEdge(9, 7)

	//	g.DFS(0)

	g.Tarjan()
}
