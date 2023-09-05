package graph

import (
	"container/heap"
	"log"
	"math"
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
	Val         int
	Edges       []*Edge
	Status      string
	ShortDist   int
	ShortParent *Vertex
}

func newVertex(val int) *Vertex {
	return &Vertex{
		Val:    val,
		Edges:  make([]*Edge, 0),
		Status: new,
	}
}

func (from *Vertex) addEdge(e *Edge) {
	from.Edges = append(from.Edges, e)
	//to.Edges = append(to.Edges, newEdge(w, to, from))
}

type Graph struct {
	Vertexes []*Vertex
	Edges    []*Edge
}

func New() Graph {
	return Graph{
		Vertexes: make([]*Vertex, 0),
		Edges:    make([]*Edge, 0),
	}
}

func (g *Graph) AddVertex(v int) {
	vertex := newVertex(v)
	g.Vertexes = append(g.Vertexes, vertex)
}

func (g *Graph) AddEdge(vert1, vert2, weight int) {
	v1 := g.findByVal(vert1)
	v2 := g.findByVal(vert2)
	e := newEdge(weight, v1, v2)
	v1.addEdge(e)
	g.Edges = append(g.Edges, e)
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
			log.Printf("from %d to %d weight %d short %d", e.From.Val, e.To.Val, e.Weight, e.To.ShortDist)
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

func (g Graph) ShortestDFS(from int, to int) {
	g.reset()
	v1 := g.findByVal(from)
	//v2 := g.findByVal(to)
	topSort := []*Vertex{}
	g.shortDFS(v1, &topSort)
	for _, v := range topSort {
		log.Printf("%d", v.Val)
		if v == v1 {
			v.ShortDist = 0
		} else {
			v.ShortDist = math.MaxInt
		}
	}

	for _, u := range topSort {
		for _, v := range u.Edges {
			if v.To.ShortDist > u.ShortDist+v.Weight {
				v.To.ShortDist = u.ShortDist + v.Weight
				v.To.ShortParent = u
			}

		}
	}

	v5 := g.findByVal(5)
	for v5.ShortParent != nil {
		log.Printf("Vert %d", v5.Val)
		v5 = v5.ShortParent
	}
}

func (g *Graph) shortDFS(v *Vertex, top *[]*Vertex) {
	v.Status = active
	for _, e := range v.Edges {
		if e.To.Status == new {
			g.shortDFS(e.To, top)
		}
	}

	v.Status = finished
	*top = append([]*Vertex{v}, *top...)
}

func (g Graph) Dijkstra(from int) {
	g.reset()
	v1 := g.findByVal(from)
	for _, v := range g.Vertexes {
		if v == v1 {
			v.ShortDist = 0
		} else {
			v.ShortDist = math.MaxInt
		}
	}

	h := NewVertexHeap()
	heap.Push(h, v1)
	g.dijkstra(v1, h)

	for _, v := range g.Vertexes {
		v5 := g.findByVal(v.Val)
		for v5.ShortParent != nil {
			log.Printf("Vert %d %d", v5.Val, v5.ShortDist)
			v5 = v5.ShortParent
		}
		log.Printf("Vert %d 0", v1.Val)
		log.Println()

	}

}

func (g *Graph) dijkstra(v *Vertex, h *VertexHeap) {
	for h.Len() > 0 {
		u := heap.Pop(h).(*Vertex)
		for _, v := range u.Edges {
			if v.To.ShortDist > u.ShortDist+v.Weight {
				v.To.ShortDist = u.ShortDist + v.Weight
				v.To.ShortParent = u
				heap.Push(h, v.To)
			}

		}
	}
}

func (g Graph) BellmanFord(from int) {
	g.reset()
	v1 := g.findByVal(from)
	for _, v := range g.Vertexes {
		if v == v1 {
			v.ShortDist = 0
		} else {
			v.ShortDist = 1000
		}
	}

	g.bellmanFord(v1)
	for _, v := range g.Vertexes {
		v5 := g.findByVal(v.Val)
		for v5.ShortParent != nil {
			log.Printf("Vert %d %d", v5.Val, v5.ShortDist)
			v5 = v5.ShortParent
		}
		log.Printf("Vert %d 0", v1.Val)
		log.Println()

	}

}

func (g *Graph) bellmanFord(v *Vertex) {

	for i := 0; i < len(g.Vertexes)-1; i++ {
		for _, e := range g.Edges {
			//log.Printf("u %d-->v %d, dist(u) %d dist(v) %d w %d", e.From.Val, e.To.Val, e.From.ShortDist, e.To.ShortDist, e.Weight)
			if e.To.ShortDist > e.From.ShortDist+e.Weight {
				e.To.ShortDist = e.From.ShortDist + e.Weight
				e.To.ShortParent = e.From
			}
			//log.Printf("u %d-->v %d, dist(u) %d dist(v) %d w %d", e.From.Val, e.To.Val, e.From.ShortDist, e.To.ShortDist, e.Weight)
			//log.Println()
		}
	}

}
