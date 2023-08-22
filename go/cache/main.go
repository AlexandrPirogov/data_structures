// very simple implementation
package main

import "fmt"

type NativeCache[T any] struct {
	size  int
	slots []string
	value []T
	hits  []int
	// ...
}

func Init[T any](sz int) NativeCache[T] {
	return NativeCache[T]{
		size:  sz,
		slots: make([]string, sz),
		value: make([]T, sz),
		hits:  make([]int, sz),
	}
}

func (n *NativeCache[T]) hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % n.size
}

func (n *NativeCache[T]) Put(key string, val T) {
	index := n.hash(key)
	if n.slots[index] == key {
		n.hits[index]++
		return
	}
	if n.slots[index] != "" {
		for k, v := range n.slots {
			if v == "" {
				index = k
				break
			}
		}
		index = n.delete()
	}
	n.slots[index] = key
	n.value[index] = val
	n.hits[index]++
}

func (n *NativeCache[T]) Get(key string) (T, error) {
	var res T
	i := n.hash(key)

	if n.slots[i] != key {
		for k := range n.slots {
			if n.slots[k] == key {
				return n.value[k], nil
			}
		}
		return res, fmt.Errorf("not found")
	}

	return n.value[i], nil
}

func (n *NativeCache[T]) delete() int {
	min := n.hits[0]
	pos := 0
	for k, v := range n.hits {
		if v < min {
			pos = k
		}
	}

	var res T
	n.hits[pos] = 0
	n.slots[pos] = ""
	n.value[pos] = res
	return pos
}
