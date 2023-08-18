package main

import (
	"os"
	"fmt"
)

//  "fmt" включите если используете

type Queue[T any] struct {
	count int
	elems []T
}

func (q *Queue[T]) Size() int {
	q = ifNil(q)
	return q.count
}

func (q *Queue[T]) Dequeue() (T, error) {
	q = ifNil(q)
	var result T

	if q.count == 0 {
		return result, fmt.Errorf("queue is empty")
	}

	result = q.elems[0]
	q.elems = q.elems[1:]
	q.count--
	return result, nil
}

func (q *Queue[T]) Enqueue(itm T) {
	q = ifNil(q)
	q.elems = append(q.elems, itm)
	q.count++
}

func (q *Queue[T]) Spin(n int) {
	head := q.elems[n:]
	tail := q.elems[0:n]
	head = append(head, tail...)
	q.elems = head
}

func ifNil[T any](q *Queue[T]) *Queue[T] {
	if q == nil {
		return &Queue[T]{
			count: 0,
			elems: make([]T, 0),
		}
	}

	return q
}
