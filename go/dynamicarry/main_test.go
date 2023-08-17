package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmpty(t *testing.T) {
	sut := DynArray[int]{}
	sut.MakeArray(16)

	assert.NotNil(t, sut.array)
	assert.Equal(t, 16, sut.capacity)
	assert.Equal(t, 0, sut.count)

}

func TestCreateEmptyWithMin(t *testing.T) {
	sut := DynArray[int]{}
	sut.MakeArray(0)

	assert.NotNil(t, sut.array)
	assert.Equal(t, 16, sut.capacity)
	assert.Equal(t, 0, sut.count)
}

func TestAppendInEmpty(t *testing.T) {
	sut := DynArray[int]{}
	sut.MakeArray(0)

	sut.Append(16)
	res, err := sut.GetItem(0)

	assert.Nil(t, err)
	assert.Equal(t, res, 16)
}

func TestAppend(t *testing.T) {
	sut := DynArray[int]{}
	sut.MakeArray(9)
	tests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for k, test := range tests {
		t.Run("", func(t *testing.T) {
			oldC := sut.count
			sut.Append(test)
			newC := sut.count

			res, err := sut.GetItem(k)

			assert.Nil(t, err)
			assert.Equal(t, oldC+1, newC)
			assert.Equal(t, res, test)
		})
	}
}

func TestGetItemError(t *testing.T) {
	sut := DynArray[int]{}
	sut.MakeArray(16)
	tests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, test := range tests {
		sut.Append(test)
	}

	max := len(tests)

	_, negativeErr := sut.GetItem(-1)
	assert.NotNil(t, negativeErr)

	_, greaterCountErr := sut.GetItem(max + 1)
	assert.NotNil(t, greaterCountErr)
}

func TestIncreaseCapacity(t *testing.T) {
	init := 16
	sut := DynArray[int]{}
	sut.MakeArray(init)

	for i := 0; i < 666; i++ {
		sut.Append(i)
		res, err := sut.GetItem(i)

		assert.Nil(t, err)
		assert.Equal(t, i, res)
		assert.Equal(t, i+1, sut.count)
		if i == sut.capacity {
			assert.Equal(t, sut.capacity, init)
			init *= 2
		}
	}

}

func TestInsertIntoTail(t *testing.T) {
	init := 16
	sut := DynArray[int]{}
	sut.MakeArray(init)

	for i := 0; i < init; i++ {
		sut.Insert(i, i)
		res, err := sut.GetItem(sut.count - 1)
		assert.Nil(t, err)
		assert.Equal(t, i, res)
	}

	assert.Equal(t, 16, sut.count)
}

func TestInsertIntoTailWithIncrease(t *testing.T) {
	init := 16
	sut := DynArray[int]{}
	sut.MakeArray(init)

	for i := 0; i < init; i++ {
		sut.Insert(i, i)
	}

	sut.Insert(16, 16)
	res, err := sut.GetItem(16)

	assert.Nil(t, err)
	assert.Equal(t, init*2, sut.capacity)
	assert.Equal(t, 16, res)

}

func TestInsertIntoHead(t *testing.T) {
	init := 16
	sut := DynArray[int]{}
	sut.MakeArray(init)

	for i := 0; i < init; i++ {
		sut.Insert(sut.count, 0)
		res, err := sut.GetItem(0)

		assert.Nil(t, err)
		assert.Equal(t, i, res)
	}

	assert.Equal(t, 16, sut.count)
}

func TestInsertIntoHeadWithIncrease(t *testing.T) {
	init := 16
	sut := DynArray[int]{}
	sut.MakeArray(init)
	for i := 0; i < init; i++ {
		sut.Insert(i, i)
	}

	item := -1
	sut.Insert(item, 0)
	res, err := sut.GetItem(0)

	assert.Nil(t, err)
	assert.Equal(t, init*2, sut.capacity)
	assert.Equal(t, item, res)
}

func TestInsertIntoBody(t *testing.T) {
	init := 10
	sut := DynArray[int]{}
	sut.MakeArray(init)

	for i := 0; i < init; i++ {
		sut.Insert(sut.count, 0)
	}

	for i := 0; i < 3; i++ {
		old, _ := sut.GetItem(i * 2)
		sut.Insert(i*-1, i*2)

		new, _ := sut.GetItem(i*2 + 1)
		res, err := sut.GetItem(i * 2)

		assert.Nil(t, err)
		assert.Equal(t, i*-1, res)
		assert.Equal(t, old, new)
	}
}

func TestInsertIntoBodyWithIncrease(t *testing.T) {
	init := 16
	sut := DynArray[int]{}
	sut.MakeArray(init)

	for i := 0; i < init; i++ {
		sut.Insert(sut.count, 0)
	}

	item, index := -1, 10

	old, _ := sut.GetItem(index)
	sut.Insert(item, index)
	new, _ := sut.GetItem(index + 1)
	res, err := sut.GetItem(index)

	assert.Nil(t, err)
	assert.Equal(t, item, res)
	assert.Equal(t, old, new)
	assert.Equal(t, init*2, sut.capacity)
}

func TestRemoveIncorrectIndex(t *testing.T) {
	sut := DynArray[int]{}
	sut.Init()

	negative := sut.Remove(-1)
	greaterCount := sut.Remove(sut.count + 1)

	assert.NotNil(t, negative)
	assert.NotNil(t, greaterCount)
}

func TestRemoveFromEmpty(t *testing.T) {
	sut := DynArray[int]{}
	sut.Init()

	old := sut.count
	err := sut.Remove(0)
	new := sut.count

	assert.NotNil(t, err)
	assert.Equal(t, old, new)
}

func TestRemoveHeadFromFilled(t *testing.T) {
	init := 16
	sut := DynArray[int]{}
	sut.Init()

	for i := 0; i < init; i++ {
		sut.Insert(i, i)
	}

	for i := 0; i < init; i++ {
		old := sut.count
		err := sut.Remove(0)
		new := sut.count
		assert.Nil(t, err)
		assert.Equal(t, old-1, new)
	}

	assert.Equal(t, 0, sut.count)
}

func TestRemoveTailFromFilled(t *testing.T) {
	init := 16
	sut := DynArray[int]{}
	sut.Init()

	for i := 0; i < init; i++ {
		sut.Insert(i, i)
	}

	for i := 0; i < init; i++ {
		old := sut.count
		err := sut.Remove(sut.count - 1)
		new := sut.count
		assert.Nil(t, err)
		assert.Equal(t, old-1, new)
	}

	assert.Equal(t, 0, sut.count)
}

func TestRemoveAndDecreaseToMinCapacity(t *testing.T) {
	init := 32
	sut := DynArray[int]{}
	sut.MakeArray(init)

	for i := 0; i < init; i++ {
		sut.Insert(i, i)
	}

	for i := 0; i < init; i++ {
		old := sut.count
		err := sut.Remove(sut.count - 1)
		new := sut.count
		assert.Nil(t, err)
		assert.Equal(t, old-1, new)
	}

	assert.Equal(t, 16, sut.capacity)
}

func TestMakeCopy(t *testing.T) {
	init := 16
	sut := DynArray[int]{}
	sut.MakeArray(init)

	for i := 0; i < init; i++ {
		sut.Insert(i, i)
	}

	clone := sut
	sut.MakeArray(32)

	for i := 0; i < init; i++ {
		t.Run("", func(t *testing.T) {
			r1, e1 := clone.GetItem(i)
			r2, e2 := sut.GetItem(i)

			assert.Nil(t, e1)
			assert.Nil(t, e2)

			assert.Equal(t, r1, r2)
		})
	}
}
