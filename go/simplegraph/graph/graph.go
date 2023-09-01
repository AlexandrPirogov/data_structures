package graph

import (
	"log"
)

type Vertex[T comparable] struct {
	Val    T
	Hit    bool
	Status string
	Adj    []*Vertex[T]
}

type Graph[T comparable] struct {
	Vertexes []*Vertex[T]
}

func New[T comparable]() Graph[T] {
	return Graph[T]{
		Vertexes: make([]*Vertex[T], 0),
	}
}

func (g *Graph[T]) AddVertex(v T) {
	add := Vertex[T]{Val: v, Hit: false, Status: "new", Adj: make([]*Vertex[T], 0)}
	g.Vertexes = append(g.Vertexes, &add)
}

func (g *Graph[T]) AddEdge(v1 T, v2 T) {
	f := g.findByVal(v1)
	s := g.findByVal(v2)
	f.Adj = append(f.Adj, s)
}

func (g *Graph[T]) DFS(from T) {
	g.unsetAll()
	g.stackDFS(from)
}

func (g *Graph[T]) stackDFS(from T) {
	head := g.findByVal(from)

	s := []*Vertex[T]{head}
	res := []T{}

	for len(s) > 0 {
		curr := s[len(s)-1]
		s = s[:len(s)-1]
		res = append(res, curr.Val)
		curr.Hit = true
		for _, v := range curr.Adj {
			if !v.Hit {
				s = append(s, v)
			}
		}
	}
	log.Print(res)
}

func (g *Graph[T]) rDFS(from T) {
	curr := g.findByVal(from)
	if curr == nil {
		return
	}
	log.Printf("vertex :%v", from)
	curr.Status = "active"
	curr.Hit = true
	for _, v := range curr.Adj {
		if !v.Hit {
			g.rDFS(v.Val)
		}
	}
	curr.Status = "finished"
}

func (g *Graph[T]) Size() int {
	return len(g.Vertexes)
}

func (g *Graph[T]) unsetAll() {
	for _, v := range g.Vertexes {
		v.Hit = false
		v.Status = "new"
	}
}

func (g *Graph[T]) findByVal(val T) *Vertex[T] {
	for _, v := range g.Vertexes {
		if v.Val == val {
			return v
		}
	}
	return nil
}

func (g *Graph[T]) IsAcyclicDFS(from T) bool {
	g.unsetAll()
	curr := g.findByVal(from)
	if curr == nil {
		return true
	}

	return g.isAcyclic(from)
}

func (g *Graph[T]) isAcyclic(from T) bool {
	curr := g.findByVal(from)
	curr.Status = "active"
	for _, v := range curr.Adj {
		if v.Status == "active" {
			log.Printf("cycle in %v %v", from, v.Val)
			return false
		}

		if v.Status == "new" {
			if !g.isAcyclic(v.Val) {
				return false
			}
		}
	}
	curr.Status = "finished"
	return true
}

func (g *Graph[T]) TopologicalSort() {
	g.unsetAll()
	from := g.Vertexes[0]
	count := len(g.Vertexes) - 1
	res := g.topologicalSort(from.Val, make([]T, len(g.Vertexes)), &count)
	log.Println(res)
}

func (g *Graph[T]) topologicalSort(from T, acc []T, count *int) []T {
	curr := g.findByVal(from)
	if curr == nil {
		return acc
	}
	curr.Status = "active"
	curr.Hit = true
	for _, v := range curr.Adj {
		if !v.Hit {
			g.topologicalSort(v.Val, acc, count)
		}
	}

	acc[*count] = curr.Val
	*count -= 1
	curr.Status = "finished"
	return acc
}

func (g *Graph[T]) CountOfComponents() {
	g.unsetAll()

	count := 0
	for _, w := range g.Vertexes {
		if !w.Hit {
			g.rDFS(w.Val)
			count++
		}
	}

	log.Printf("Components count %d", count)
}

func (g *Graph[T]) Forest() {
	g.unsetAll()
	count := 0
	for _, w := range g.Vertexes {
		if !w.Hit {
			g.DFS(w.Val)
			count++
		}
	}
	log.Printf("Forest counts: %d", count)
}

func (g *Graph[T]) IsStrongConnected() bool {
	g1 := g.Transpose()

	vertex := g.Vertexes[0].Val

	g.DFS(vertex)
	log.Println(g.VisitedAll())
	g1.DFS(vertex)
	log.Println(g1.VisitedAll())
	return g.VisitedAll() && g1.VisitedAll()
}

func (g *Graph[T]) VisitedAll() bool {
	for _, v := range g.Vertexes {
		if !v.Hit {
			return false
		}
	}
	return true
}

func (g *Graph[T]) Transpose() *Graph[T] {
	tG := Graph[T]{
		make([]*Vertex[T], 0),
	}

	for _, v := range g.Vertexes {
		tG.AddVertex(v.Val)
	}

	for _, v := range g.Vertexes {
		for _, w := range v.Adj {
			tG.AddEdge(w.Val, v.Val)
		}
	}

	return &tG
}
