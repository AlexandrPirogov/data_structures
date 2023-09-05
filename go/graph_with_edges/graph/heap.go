package graph

import (
	"container/heap"
)

func NewHeap() *EdgeHeap {
	e := &EdgeHeap{}
	heap.Init(e)
	return e
}

// An EdgeHeap is a min-heap of ints.
type EdgeHeap []*Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Edge))
}

func (h *EdgeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewVertexHeap() *VertexHeap {
	e := &VertexHeap{}
	heap.Init(e)
	return e
}

// An EdgeHeap is a min-heap of ints.
type VertexHeap []*Vertex

func (h VertexHeap) Len() int           { return len(h) }
func (h VertexHeap) Less(i, j int) bool { return h[i].ShortDist < h[j].ShortDist }
func (h VertexHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *VertexHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Vertex))
}

func (h VertexHeap) InHeap(v *Vertex) bool {
	for _, vert := range h {
		if vert == v {
			return true
		}
	}
	return false
}

func (h *VertexHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
