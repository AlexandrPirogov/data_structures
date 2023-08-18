package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	sut := Stack[int]{}
	assert.Equal(t, 0, sut.Size())
}

func TestPushInNil(t *testing.T) {
	var sut *Stack[int]
	sut.Push(1)
}

func TestPushCountChanges(t *testing.T) {
	sut := Stack[int]{}

	for i := 0; i < 10; i++ {
		old := sut.Size()
		sut.Push(i)
		new := sut.Size()

		assert.Equal(t, old+1, new)
	}
}

func TestPopFromNil(t *testing.T) {
	var sut *Stack[int]
	_, err := sut.Pop()
	assert.NotNil(t, err)
}

func TestPopEmpty(t *testing.T) {
	sut := Stack[int]{}

	_, err := sut.Pop()
	assert.NotNil(t, err)
}

func TestPopFilled(t *testing.T) {
	sut := Stack[int]{}
	tests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, num := range tests {
		sut.Push(num)
	}

	for i, j := 0, len(tests)-1; i < j; i, j = i+1, j-1 {
		tests[i], tests[j] = tests[j], tests[i]
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res, err := sut.Pop()

			assert.Nil(t, err)
			assert.Equal(t, test, res)
		})
	}

	_, err := sut.Pop()
	assert.NotNil(t, err)
}

func TestPeekFromNil(t *testing.T) {
	var sut *Stack[int]

	_, err := sut.Peek()

	assert.NotNil(t, err)
}

func TestPeekFromEmpty(t *testing.T) {
	sut := Stack[int]{}

	_, err := sut.Peek()

	assert.NotNil(t, err)
}

func TestPeekFromFilled(t *testing.T) {
	sut := Stack[int]{}
	tests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, num := range tests {
		sut.Push(num)
	}

	for i, j := 0, len(tests)-1; i < j; i, j = i+1, j-1 {
		tests[i], tests[j] = tests[j], tests[i]
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res, err := sut.Peek()

			assert.Nil(t, err)
			assert.Equal(t, test, res)

			sut.Pop()
		})
	}

	_, err := sut.Peek()

	assert.NotNil(t, err)

}
