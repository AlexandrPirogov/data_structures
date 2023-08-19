package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzHashFun(f *testing.F) {
	sz := 19
	sut := Init(sz, 3)
	f.Fuzz(func(t *testing.T, str string) {
		ind := sut.HashFun(str)
		assert.Greater(t, ind, -1)
		assert.Less(t, ind, sz)
	})
}

func TestPutInEmpty(t *testing.T) {
	sz, step := 19, 3
	sut := Init(sz, step)

	val := "someval"
	ind := sut.HashFun(val)
	put := sut.Put(val)

	assert.Equal(t, ind, put)
	assert.Equal(t, val, sut.slots[ind])
	assert.True(t, sut.busy[ind])
}

func TestSeekSlotWithoutCollision(t *testing.T) {
	sz, step := 19, 3
	sut := Init(sz, step)

	val := "someval"
	res := sut.SeekSlot(val)
	put := sut.Put(val)

	assert.Equal(t, put, res)
}

func TestSeekSlotWithCollision(t *testing.T) {
	sz, step := 19, 3
	sut := Init(sz, step)
	val := "someval"
	val1 := "somevla"

	put := sut.Put(val)
	res := sut.SeekSlot(val1)

	assert.NotEqual(t, put, res)
	assert.NotEqual(t, -1, res)

}

func TestSeekSlotWithoutCollisionFilled(t *testing.T) {
	sz, step := 7, 3
	sut := Init(sz, step)
	strs := []string{"a", "b", "c", "d", "e", "f", "g"}

	for _, str := range strs {
		sut.Put(str)
	}

	ind := sut.SeekSlot("some")

	assert.Equal(t, -1, ind)
}

func TestSeekSlotWithoutCollisionNotFilled(t *testing.T) {
	sz, step := 7, 3
	sut := Init(sz, step)
	strs := []string{"a", "b", "c", "d", "e", "f"}

	for _, str := range strs {
		sut.Put(str)
	}

	ind := sut.SeekSlot("some")

	assert.NotEqual(t, -1, ind)
}

func TestSeekSlotWithCollisionFilled(t *testing.T) {
	sz, step := 7, 3
	sut := Init(sz, step)
	strs := []string{"a"}

	for _, str := range strs {
		sut.Put(str)
	}

	seek := sut.SeekSlot("a")
	ind := sut.Put("a")

	assert.Equal(t, -1, ind)
	assert.NotEqual(t, -1, seek)
}

func TestFindInEmpty(t *testing.T) {
	sz, step := 19, 3
	sut := Init(sz, step)

	res := sut.Find("not exists")
	assert.Equal(t, -1, res)
}

func TestFindInFilled(t *testing.T) {
	sz, step := 19, 3
	sut := Init(sz, step)

	strs := []string{"a", "b", "c", "d", "e", "f"}

	for _, str := range strs {
		sut.Put(str)
	}

	for _, str := range strs {
		t.Run("", func(t *testing.T) {
			res := sut.Find(str)
			assert.NotEqual(t, -1, res)
		})
	}
}

func TestFindNotExistingInFilled(t *testing.T) {
	sz, step := 19, 3
	sut := Init(sz, step)

	strs := []string{"a", "b", "c", "d", "e", "f"}

	for _, str := range strs {
		sut.Put(str)
	}

	for _, str := range strs {
		t.Run("", func(t *testing.T) {
			res := sut.Find(str + "suffix")
			assert.Equal(t, -1, res)
		})
	}
}
