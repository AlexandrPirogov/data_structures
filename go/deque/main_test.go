package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmpty(t *testing.T) {
	var sut *Deque[int]

	assert.Equal(t, 0, sut.Size())
}

func TestSingleAddInFront(t *testing.T) {
	sut := Deque[int]{}

	old := sut.Size()
	sut.AddFront(1)
	new := sut.Size()

	assert.Equal(t, old+1, new)
}

func TestSingleAddFrontRemoveFront(t *testing.T) {
	sut := Deque[int]{}

	item := 1
	sut.AddFront(item)

	res, err := sut.RemoveFront()

	assert.Nil(t, err)
	assert.Equal(t, item, res)
}

func TestSingleAddFrontRemoveTail(t *testing.T) {
	sut := Deque[int]{}

	item := 1
	sut.AddFront(item)

	old := sut.Size()
	res, err := sut.RemoveTail()
	new := sut.Size()

	assert.Nil(t, err)
	assert.Equal(t, item, res)
	assert.Equal(t, old-1, new)
}

func TestSingleAddInRail(t *testing.T) {
	sut := Deque[int]{}

	old := sut.Size()
	sut.AddTail(1)
	new := sut.Size()

	assert.Equal(t, old+1, new)
}

func TestSingleAddTailRemoveFront(t *testing.T) {
	sut := Deque[int]{}

	item := 1
	sut.AddTail(item)

	old := sut.Size()
	res, err := sut.RemoveFront()
	new := sut.Size()

	assert.Nil(t, err)
	assert.Equal(t, item, res)
	assert.Equal(t, old-1, new)
}

func TestSingleAddTailRemoveTail(t *testing.T) {
	sut := Deque[int]{}

	item := 1
	sut.AddTail(item)

	res, err := sut.RemoveTail()

	assert.Nil(t, err)
	assert.Equal(t, item, res)
}

func TestRemoveFrontEmpty(t *testing.T) {
	sut := Deque[int]{}

	_, err := sut.RemoveFront()
	assert.NotNil(t, err)
}

func TestRemoveTailEmpty(t *testing.T) {
	sut := Deque[int]{}

	_, err := sut.RemoveTail()
	assert.NotNil(t, err)
}
