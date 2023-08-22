package main

import (
	"fmt"
)

type NativeDictionary[T any] struct {
	size   int
	slots  []string
	values []T
}

// создание экземпляра словаря
func Init[T any](sz int) NativeDictionary[T] {
	nd := NativeDictionary[T]{size: sz, slots: nil, values: nil}
	nd.slots = make([]string, sz)
	nd.values = make([]T, sz)
	return nd
}

func (nd *NativeDictionary[T]) HashFun(value string) int {
	// всегда возвращает корректный индекс слота
	sum := 0
	for _, b := range value {
		sum += int(b)
	}
	return sum % nd.size
}

func (nd *NativeDictionary[T]) IsKey(key string) bool {
	// возвращает true если ключ имеется
	ind := nd.HashFun(key)
	if nd.slots[ind] == key {
		return true
	}

	col := nd.collisionFind(key)

	return col != -1
}

func (nd *NativeDictionary[T]) Get(key string) (T, error) {
	var result T
	if key == "" {
		return result, fmt.Errorf("not found")
	}
	// возвращает value для key,
	// или error если ключ не найден

	ind := nd.HashFun(key)

	if nd.slots[ind] != key {
		ind = nd.collisionFind(key)
		if ind == -1 {
			return result, fmt.Errorf("not found")
		}
	}
	result = nd.values[ind]
	return result, nil
}

func (nd *NativeDictionary[T]) Put(key string, value T) {
	if key == "" {
		return
	}
	// гарантированно записываем
	// значение value по ключу key
	i := nd.HashFun(key)
	if nd.slots[i] == key {
		nd.values[i] = value
		return
	}

	col := nd.collisionFind(key)
	if col > -1 && nd.slots[col] == key {
		nd.values[col] = value
		return
	}

	ind := nd.seekSlot(key)
	if ind == -1 {
		return
	}

	nd.slots[ind] = key
	nd.values[ind] = value
}

func (nd *NativeDictionary[T]) seekSlot(key string) int {
	ind := nd.HashFun(key)
	if nd.slots[ind] == "" {
		return ind
	}

	for k, v := range nd.slots {
		if v == "" {
			return k
		}
	}

	return -1
}

func (nd *NativeDictionary[T]) collisionFind(key string) int {
	for i, k := range nd.slots {
		if k == key {
			return i
		}
	}

	return -1
}
