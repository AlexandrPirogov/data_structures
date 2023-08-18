package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	sut := Queue[int]{}

	assert.Equal(t, 0, sut.Size())
}

func TestCountInNil(t *testing.T) {
	var sut *Queue[int]

	assert.Equal(t, 0, sut.Size())
}

func TestEnqueueInEmpty(t *testing.T) {
	sut := Queue[int]{}

	item := 1
	sut.Enqueue(item)

	assert.Equal(t, item, sut.count)
}

func TestDequeueFromEmpty(t *testing.T) {
	var sut *Queue[int]

	_, err := sut.Dequeue()
	assert.NotNil(t, err)
}

func TestDequeueEnqueue(t *testing.T) {
	sut := Queue[int]{}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, num := range nums {
		sut.Enqueue(num)
	}

	for _, test := range nums {
		t.Run("", func(t *testing.T) {
			old := sut.count
			res, err := sut.Dequeue()
			new := sut.count

			assert.Nil(t, err)
			assert.Equal(t, old-1, new)
			assert.Equal(t, test, res)
		})
	}

	_, res := sut.Dequeue()
	assert.NotNil(t, res)
}

func TestQueueSpin(t *testing.T) {
	sut := Queue[int]{}
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, num := range nums {
		sut.Enqueue(num)
	}

	sut.Spin(5)
	sut.Spin(5)

	res, _ := sut.Dequeue()

	assert.Equal(t, nums[0], res)
}
