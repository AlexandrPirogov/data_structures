package main

import (
	"fmt"
	"os"
)

//  "fmt" включите если используете

const error_stack_empty = "stack is empty"

type Stack[T any] struct {
	elems []T
	count int
}

func ifNil[T any](st *Stack[T]) *Stack[T] {
	if st == nil {
		return &Stack[T]{
			elems: make([]T, 0),
			count: 0,
		}
	}
	return st
}

func (st *Stack[T]) Size() int {
	st = ifNil(st)
	return st.count
}

func (st *Stack[T]) Peek() (T, error) {
	st = ifNil(st)
	var result T
	if st.count == 0 {
		return result, fmt.Errorf(error_stack_empty)
	}

	result = st.elems[st.count-1]
	return result, nil
}

func (st *Stack[T]) Pop() (T, error) {
	st = ifNil(st)
	var result T
	if st.count == 0 {
		return result, fmt.Errorf(error_stack_empty)
	}
	result = st.elems[st.count-1]
	st.count--
	return result, nil
}

func (st *Stack[T]) Push(itm T) {
	st = ifNil(st)
	st.elems = append(st.elems, itm)
	st.count++
}
