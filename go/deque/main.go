package main

import (
	"os"
	"fmt"
)

//  "fmt"

type Deque[T any] struct {
	count int
	elems []T
}

func (d *Deque[T]) Size() int {
	return d.count
}

func (d *Deque[T]) AddFront(itm T) {
	d.elems = append(d.elems, itm)
	d.count++
}

func (d *Deque[T]) AddTail(itm T) {
	d.elems = append([]T{itm}, d.elems...)
	d.count++
}

func (d *Deque[T]) RemoveFront() (T, error) {
	var result T
	if d.count == 0 {
		return result, fmt.Errorf("deque is empty")
	}

	result = d.elems[d.count-1]
	d.elems = d.elems[:d.count-1]
	d.count--
	return result, nil
}

func (d *Deque[T]) RemoveTail() (T, error) {
	var result T
	if d.count == 0 {
		return result, fmt.Errorf("deque is empty")
	}

	result = d.elems[0]
	d.elems = d.elems[1:]
	d.count--
	return result, nil
}

func ifNil[T any](d *Deque[T]) *Deque[T] {
	if d == nil || d.elems == nil {
		return &Deque[T]{
			count: 0,
			elems: make([]T, 0),
		}
	}
	return d
}
