package main

import (
	"log"
	"simplegraph/graph"
)

func main() {
	g := graph.New[int]()
	for i := 0; i < 2; i++ {
		g.AddVertex(i)
	}

	g.AddEdge(0, 1)

	//	g.DFS(0)

	log.Println(g.IsStrongConnected())
}
