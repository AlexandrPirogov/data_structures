package main

import (
	"constraints"
	//      "fmt"
"	os"
	"strconv"
)

type PowerSet[T constraints.Ordered] struct {
	// ваша реализация хранилища
	slots map[T]bool
}

// создание экземпляра множества
func Init[T constraints.Ordered]() PowerSet[T] {
	return PowerSet[T]{
		slots: map[T]bool{},
	}
}

func (p *PowerSet[T]) Size() int {
	// количество элементов в множестве
	return len(p.slots)
}

func (p *PowerSet[T]) Put(value T) {
	// всегда срабатывает
	p.slots[value] = true

}

func (p *PowerSet[T]) Get(value T) bool {
	// возвращает true если value имеется в множестве
	_, ok := p.slots[value]
	return ok
}

func (p *PowerSet[T]) Remove(value T) bool {
	// возвращает true если value удалено
	_, ok := p.slots[value]
	if !ok {
		return false
	}

	delete(p.slots, value)
	return true
}

func (p *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	// пересечение текущего множества и set2
	result := Init[T]()
	for k := range p.slots {
		if set2.slots[k] {
			result.Put(k)
		}
	}
	return result
}

func (p *PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	// объединение текущего множества и set2
	result := Init[T]()
	for k := range p.slots {
		result.Put(k)
	}

	for k := range set2.slots {
		result.Put(k)
	}
	// ...
	return result
}

func (p *PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	// разница текущего множества и set2
	result := Init[T]()

	for k := range p.slots {
		if _, ok := set2.slots[k]; !ok {
			result.Put(k)
		}
	}
	return result
}

func (p *PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {

	for k := range set2.slots {
		if _, ok := p.slots[k]; !ok {
			return false
		}
	}
	return true
}
