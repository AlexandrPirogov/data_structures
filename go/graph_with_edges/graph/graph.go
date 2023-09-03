package graph

import (
	"container/heap"
	"log"
)

type Edge struct {
	To     *Vertex
	From   *Vertex
	Weight int
}

func newEdge(w int, from *Vertex, to *Vertex) *Edge {
	return &Edge{
		To:     to,
		Weight: w,
		From:   from,
	}
}

const new = "new"
const active = "active"
const finished = "finished"

type Vertex struct {
	Val    int
	Edges  []*Edge
	Status string
}

func newVertex(val int) *Vertex {
	return &Vertex{
		Val:    val,
		Edges:  make([]*Edge, 0),
		Status: new,
	}
}

func (from *Vertex) addEdge(w int, to *Vertex) {
	from.Edges = append(from.Edges, newEdge(w, from, to))
	to.Edges = append(to.Edges, newEdge(w, to, from))
}

type Graph struct {
	Vertexes []*Vertex
}

func New() Graph {
	return Graph{Vertexes: make([]*Vertex, 0)}
}

func (g *Graph) AddVertex(v int) {
	vertex := newVertex(v)
	g.Vertexes = append(g.Vertexes, vertex)
}

func (g *Graph) AddEdge(vert1, vert2, weight int) {
	v1 := g.findByVal(vert1)
	v2 := g.findByVal(vert2)
	v1.addEdge(weight, v2)
}

func (g *Graph) findByVal(val int) *Vertex {
	for _, v := range g.Vertexes {
		if v.Val == val {
			return v
		}
	}
	return nil
}

func (g *Graph) reset() {
	for _, v := range g.Vertexes {
		v.Status = new
	}
}

func DFS(g *Graph, from int) {
	g.reset()
	v := g.findByVal(from)
	rDFS(v)
}

func rDFS(v *Vertex) {
	v.Status = active
	for _, e := range v.Edges {
		if e.To.Status == "new" {
			log.Printf("from %d to %d weight %d", e.From.Val, e.To.Val, e.Weight)
			rDFS(e.To)
		}
	}
	v.Status = finished
}

func Primes(g *Graph) *Graph {
	g.reset()
	acc := New()
	from := g.Vertexes[0]
	h := NewHeap()
	log.Printf("edges len %d", len(from.Edges))
	for _, e := range from.Edges {
		if e.To.Status == new {
			heap.Push(h, e)
		}
	}

	g.reset()

	for _, v := range g.Vertexes {
		vert := newVertex(v.Val)
		acc.Vertexes = append(acc.Vertexes, vert)

	}

	primeLoop(acc.Vertexes[0], h, g, &acc)
	return &acc
}

func primeLoop(v *Vertex, h *EdgeHeap, g *Graph, acc *Graph) {
	if h.Len() == 0 {
		return
	}
	e := heap.Pop(h).(*Edge)
	//log.Printf("e %d", e.Weight)
	add := acc.findByVal(e.To.Val)
	if add.Status == new {

		if e.To.Status == new {
			acc.AddEdge(e.From.Val, e.To.Val, e.Weight)
		}
		e.To.Status = active
		for _, ae := range e.To.Edges {
			if ae.To.Status == new {
				heap.Push(h, ae)
			}
		}

		primeLoop(e.To, h, g, acc)
	}
	v.Status = finished

}
