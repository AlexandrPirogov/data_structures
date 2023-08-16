package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmptyDoubleLinkedList(t *testing.T) {
	sut := LinkedList2{}

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
	assert.Equal(t, 0, sut.Count())
}

func TestAddInTail(t *testing.T) {
	sut := LinkedList2{}

	nodes := make([]Node, 0)
	for i := 0; i < 10; i++ {
		node := Node{nil, nil, i}
		nodes = append(nodes, node)
	}

	for _, node := range nodes {
		t.Run("", func(t *testing.T) {
			sut.AddInTail(node)

			assert.Equal(t, sut.tail.value, node.value)
		})
	}
}

func TestFindEmpty(t *testing.T) {
	sut := LinkedList2{}

	res, err := sut.Find(10)

	assert.NotNil(t, err)
	assert.Equal(t, -1, res.value)
}

func TestFindExisting(t *testing.T) {
	sut := LinkedList2{}

	nodes := make([]Node, 0)
	for i := 0; i < 10; i++ {
		node := Node{nil, nil, i}
		nodes = append(nodes, node)
		sut.AddInTail(node)
	}

	for _, node := range nodes {
		t.Run("", func(t *testing.T) {
			res, err := sut.Find(node.value)

			assert.Nil(t, err)
			assert.Equal(t, node.value, res.value)
		})
	}
}

func TestFindNotExisting(t *testing.T) {
	sut := LinkedList2{}

	nodes := make([]Node, 0)
	for i := 1; i <= 10; i++ {
		node := Node{nil, nil, i}
		nodes = append(nodes, node)
		sut.AddInTail(node)
	}

	for _, node := range nodes {
		t.Run("", func(t *testing.T) {
			res, err := sut.Find(node.value * -1)

			assert.NotNil(t, err)
			assert.Equal(t, -1, res.value)
		})
	}
}

func TestFindAllEmpty(t *testing.T) {
	sut := LinkedList2{}

	res := sut.FindAll(1)
	assert.Empty(t, res)
}

func TestFindAllExisting(t *testing.T) {
	sut := LinkedList2{}

	nodes := []Node{
		{
			prev:  nil,
			next:  nil,
			value: 1,
		},
		{
			prev:  nil,
			next:  nil,
			value: 2,
		},
		{
			prev:  nil,
			next:  nil,
			value: 3,
		},
		{
			prev:  nil,
			next:  nil,
			value: 3,
		},
		{
			prev:  nil,
			next:  nil,
			value: 2,
		},
		{
			prev:  nil,
			next:  nil,
			value: 1,
		},
	}

	for _, node := range nodes {
		sut.AddInTail(node)
	}

	for _, node := range nodes {
		t.Run("", func(t *testing.T) {
			res := sut.FindAll(node.value)

			assert.Greater(t, len(res), 0)
			for _, found := range res {
				assert.Equal(t, node.value, found.value)
			}
		})
	}
}

func TestFindAllNotExisting(t *testing.T) {
	sut := LinkedList2{}

	nodes := []Node{
		{
			prev:  nil,
			next:  nil,
			value: 1,
		},
		{
			prev:  nil,
			next:  nil,
			value: 2,
		},
		{
			prev:  nil,
			next:  nil,
			value: 3,
		},
		{
			prev:  nil,
			next:  nil,
			value: 3,
		},
		{
			prev:  nil,
			next:  nil,
			value: 2,
		},
		{
			prev:  nil,
			next:  nil,
			value: 1,
		},
	}

	for _, node := range nodes {
		sut.AddInTail(node)
	}

	for _, node := range nodes {
		t.Run("", func(t *testing.T) {
			res := sut.FindAll(node.value * -1)

			assert.Empty(t, res)
		})
	}
}

func TestCountEmpty(t *testing.T) {
	sut := LinkedList2{}
	assert.Empty(t, sut.Count())
}

func TestCountNotEmpty(t *testing.T) {
	sut := LinkedList2{}

	for i := 1; i <= 10; i++ {
		node := Node{nil, nil, i}
		sut.AddInTail(node)
	}

	assert.Equal(t, 10, sut.Count())
}

func TestCleanEmpty(t *testing.T) {
	sut := LinkedList2{}

	sut.Clean()

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
	assert.Empty(t, sut.Count())
}

func TestCleanNonEmpty(t *testing.T) {
	sut := LinkedList2{}

	for i := 1; i <= 10; i++ {
		node := Node{nil, nil, i}
		sut.AddInTail(node)
	}

	sut.Clean()

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
	assert.Empty(t, sut.Count())
}
