package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// assert that created new ll is empty
func TestCreateList(t *testing.T) {
	//arrange
	sut := LinkedList{}

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.head)
	assert.Equal(t, sut.Count(), 0)
}

func TestAddInEmptyList(t *testing.T) {
	sut := LinkedList{}
	insert := Node{nil, 10}

	sut.AddInTail(insert)

	assert.Equal(t, sut.Count(), 1)
	assert.Equal(t, sut.head, sut.tail)
	assert.Equal(t, sut.head, &insert)
}

// Adding many different notes
func TestAddManyList(t *testing.T) {
	tests := []Node{}

	for i := 0; i < 10; i++ {
		tests = append(tests, Node{nil, i})
	}

	sut := LinkedList{}

	for _, edge := range tests {
		t.Run(fmt.Sprintf("%d", edge.value), func(t *testing.T) {
			sut.AddInTail(edge)
			assert.Equal(t, sut.tail, &edge)
		})
	}
}

// checks find method with list
// with distinct values
func TestFindInSetList(t *testing.T) {
	sut := LinkedList{}
	for i := 0; i < 10; i++ {
		sut.AddInTail(Node{nil, i})
	}

	for i := 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			res, err := sut.Find(i)
			assert.Nil(t, err)
			assert.Equal(t, i, res.value)
		})
	}
}

// Checks find method in empty list
func TestFindInEmptyList(t *testing.T) {
	sut := LinkedList{}
	res, err := sut.Find(10)
	assert.NotNil(t, err)
	assert.Equal(t, -1, res.value)
}

// Checks find method in list with duplicates
func TestFindInList(t *testing.T) {
	duplicate := 10
	sut := LinkedList{}
	for i := 0; i < 10; i++ {
		sut.AddInTail(Node{nil, duplicate})
	}

	for i := 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			res, err := sut.Find(duplicate)
			assert.Nil(t, err)
			assert.Equal(t, duplicate, res.value)
		})
	}
}

// Checks that findall returns empty slice for empty list
func TestFindAllInEmptyList(t *testing.T) {
	sut := LinkedList{}

	res := sut.FindAll(10)
	assert.Equal(t, len(res), 0)

}

// checks that findall in set list returns exactly slice with single elem
func TestFindAllInSetList(t *testing.T) {
	sut := LinkedList{}
	for i := 0; i < 10; i++ {
		sut.AddInTail(Node{nil, i})
	}

	for i := 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			res := sut.FindAll(i)
			assert.Equal(t, len(res), 1)
		})
	}
}

// checks that findall in list with duplicates
// returns slice with all elems in lsit
func TestFindAllInList(t *testing.T) {
	duplicate := 10
	sut := LinkedList{}
	for i := 0; i < 10; i++ {
		sut.AddInTail(Node{nil, duplicate})
	}

	for i := 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			res := sut.FindAll(duplicate)
			assert.NotEmpty(t, res)
		})
	}
}

func TestCleanEmptyList(t *testing.T) {
	sut := LinkedList{}

	sut.Clean()

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)

	assert.Equal(t, 0, sut.Count())
}

func TestCleanList(t *testing.T) {
	sut := LinkedList{}
	for i := 0; i < 10; i++ {
		sut.AddInTail(Node{nil, i})
	}

	sut.Clean()

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)

	assert.Equal(t, 0, sut.Count())
}

// checks that insert first make head equal inserted elem
func TestInsertFirstInEmptyList(t *testing.T) {
	sut := LinkedList{}

	insert := Node{nil, 1}
	sut.InsertFirst(insert)

	assert.Equal(t, sut.head, &insert)
	assert.Nil(t, sut.head.next)
}

func TestInsertFirstInList(t *testing.T) {
	sut := LinkedList{}
	for i := 0; i < 10; i++ {
		sut.AddInTail(Node{nil, i})
	}

	insert := Node{nil, 20}
	sut.InsertFirst(insert)

	assert.Equal(t, sut.head.value, insert.value)
}

func TestInsertInList(t *testing.T) {
	sut := LinkedList{}
	nodes := make([]Node, 0)
	for i := 0; i < 10; i++ {
		add := Node{nil, i}
		nodes = append(nodes, add)
		sut.AddInTail(add)
	}

	for k, edge := range nodes {
		t.Run("", func(t *testing.T) {
			sut.Insert(&edge, Node{nil, k})
			res := sut.FindAll(k)

			assert.Equal(t, len(res), 2)
		})
	}

	for _, edge := range nodes {
		t.Run("", func(t *testing.T) {
			node, err := sut.Find(edge.value)

			assert.Nil(t, err)
			assert.Equal(t, node.next.value, edge.value)
		})
	}
}

func TestDeleteFromEmptyList(t *testing.T) {
	sut := LinkedList{}
	sut.Delete(10, false)

	assert.Nil(t, sut.head)
	assert.Equal(t, 0, sut.Count())
}

func TestDeleteAllFromEmptyList(t *testing.T) {
	sut := LinkedList{}
	sut.Delete(10, true)

	assert.Nil(t, sut.head)
	assert.Equal(t, 0, sut.Count())
}

func TestDeleteOneFromList(t *testing.T) {
	sut := LinkedList{}
	nodes := make([]Node, 0)
	for i := 0; i < 10; i++ {
		add := Node{nil, i}
		nodes = append(nodes, add)
		sut.AddInTail(add)
	}
	for _, edge := range nodes {
		t.Run("", func(t *testing.T) {
			oldcount := sut.Count()
			sut.Delete(edge.value, false)
			newcount := sut.Count()

			assert.Greater(t, oldcount, newcount)
		})
	}
}

func TestDeleteOneFromTail(t *testing.T) {
	sut := LinkedList{}
	nodes := make([]Node, 0)
	for i := 3; i >= 0; i-- {
		add := Node{nil, i}
		nodes = append(nodes, add)
		sut.AddInTail(add)
	}

	for k := range nodes {
		t.Run(fmt.Sprintf("deleting %d", k), func(t *testing.T) {
			old := sut.tail
			sut.Delete(k, false)
			new := sut.tail
			if new != nil {
				assert.Equal(t, new.value-1, old.value)
			}
		})
	}

	assert.Nil(t, sut.tail)
}

func TestDeleteOneFromHead(t *testing.T) {
	sut := LinkedList{}
	nodes := make([]Node, 0)
	for i := 0; i < 3; i++ {
		add := Node{nil, i}
		nodes = append(nodes, add)
		sut.AddInTail(add)
	}

	for _, edge := range nodes {
		t.Run(fmt.Sprintf("deleting %d", edge.value), func(t *testing.T) {
			old := sut.head
			sut.Delete(edge.value, false)
			new := sut.head
			if new != nil {
				assert.Equal(t, new.value-1, old.value)
			}
		})
	}

	assert.Nil(t, sut.head)
}

func TestDeleteAllEmpty(t *testing.T) {
	sut := LinkedList{}

	sut.Delete(-1, true)

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
	assert.Empty(t, sut.Count())

}

func TestDeleteAllFromList(t *testing.T) {
	sut := LinkedList{}
	nodes := []Node{
		{
			nil,
			2,
		},
		{
			nil,
			1,
		},
		{
			nil,
			2,
		},
		{
			nil,
			2,
		},
	}

	for _, node := range nodes {
		sut.AddInTail(node)
	}

	for _, edge := range nodes {
		t.Run(fmt.Sprintf("egdt %d", edge.value), func(t *testing.T) {
			sut.Delete(edge.value, true)
			res := sut.FindAll(edge.value)
			assert.Empty(t, res)
		})
	}

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
	assert.Equal(t, 0, sut.Count())
}

func TestCountInEmptyList(t *testing.T) {
	sut := LinkedList{}

	assert.Equal(t, 0, sut.Count())
}
