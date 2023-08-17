package main

import (
	"errors"
	"fmt"

	"os"
)

type DynArray[T any] struct {
	count    int
	capacity int
	array    []T
}

func (da *DynArray[T]) Init() {
	da.count = 0
	da.MakeArray(16)
}

func (da *DynArray[T]) MakeArray(sz int) {
	//  копируем содержимое array в arr ...
	if sz < 16 {
		sz = 16
	}

	da.capacity = sz

	var arr = make([]T, da.capacity)
	copy(arr, da.array)
	da.array = arr //
}

func (da *DynArray[T]) Insert(itm T, index int) error {
	if index < 0 || index > da.count {
		return fmt.Errorf("bad index '%d'", index)
	}

	da.increaseCap()
	for i := da.count; i > index; i-- {
		da.array[i] = da.array[i-1]
	}

	da.array[index] = itm
	da.count++
	return nil
}

func (da *DynArray[T]) Remove(index int) error {
	if index < 0 || index >= da.count {
		return fmt.Errorf("bad index '%d'", index)
	}

	da.decreaseCap()
	da.array = append(da.array[:index], da.array[index+1:]...)
	var defauilt T
	da.array = append(da.array, defauilt)

	if da.count > 0 {
		da.count--
	}
	return nil
}

func (da *DynArray[T]) Append(itm T) {
	da.increaseCap()
	da.array[da.count] = itm
	da.count++
}

func (da *DynArray[T]) GetItem(index int) (T, error) {
	var result T
	if index < 0 || index > da.count {
		return result, errors.New("incorrect index")
	}
	// ...

	result = da.array[index]
	return result, nil
}

func (da *DynArray[T]) increaseCap() {
	if da.count == da.capacity {
		da.capacity *= 2
		tmp := da.array
		da.array = make([]T, da.capacity)
		copy(da.array, tmp)
	}
}

func (da *DynArray[T]) decreaseCap() {
	if float64(float64(da.count)/float64(da.capacity)) < 0.5 {
		da.capacity = int(float64(da.capacity) / 1.5)
		if da.capacity < 16 {
			da.capacity = 16
		}
		tmp := da.array
		da.array = make([]T, da.capacity)
		copy(da.array, tmp)
	}

}
