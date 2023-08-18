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

func TestInsertFirstInEmptyOnce(t *testing.T) {
	sut := LinkedList2{}

	node := Node{nil, nil, 1}
	sut.InsertFirst(node)

	assert.NotNil(t, sut.head)
	assert.Equal(t, sut.head, sut.tail)
	assert.Equal(t, sut.head.value, node.value)
}

func TestInsertFirstHeadTail(t *testing.T) {
	sut := LinkedList2{}

	node := Node{nil, nil, 1}
	node1 := Node{nil, nil, 2}

	sut.InsertFirst(node)
	sut.InsertFirst(node1)

	assert.Equal(t, sut.head.next, sut.tail)
	assert.Equal(t, sut.tail.prev, sut.head)
	assert.Equal(t, 2, sut.Count())
}

func TestInsertFirst(t *testing.T) {
	sut := LinkedList2{}

	for i := 1; i <= 10; i++ {
		node := Node{nil, nil, i}
		sut.InsertFirst(node)

		assert.NotNil(t, sut.head)
		assert.NotNil(t, sut.tail)
		assert.Equal(t, sut.head.value, node.value)
	}

	assert.Equal(t, 1, sut.tail.value)
}

func TestInsertInTail(t *testing.T) {
	sut := LinkedList2{}

	nodes := make([]Node, 0)
	for i := 0; i < 4; i++ {
		node := Node{nil, nil, i}
		nodes = append(nodes, node)
		sut.AddInTail(node)
	}

	for _, test := range nodes {
		t.Run("", func(t *testing.T) {
			sut.Insert(sut.tail, test)
			res := sut.FindAll(test.value)

			assert.Equal(t, 2, len(res))
			assert.Equal(t, test.value, sut.tail.value)
		})
	}

}

func TestInsertInHead(t *testing.T) {
	sut := LinkedList2{}

	nodes := make([]Node, 0)
	for i := 0; i < 10; i++ {
		node := Node{nil, nil, i}
		nodes = append(nodes, node)
		sut.AddInTail(node)
	}

	for _, test := range nodes {
		t.Run("", func(t *testing.T) {
			sut.Insert(sut.head, test)
			res := sut.FindAll(test.value)

			assert.Equal(t, 2, len(res))
			assert.Equal(t, sut.head.next.value, test.value)
		})
	}

}

func TestAfterHead(t *testing.T) {
	sut := LinkedList2{}

	head := Node{nil, nil, 1}
	sut.AddInTail(head)
	sut.AddInTail(head)
	after := sut.head

	node := Node{nil, nil, 2}

	sut.Insert(after, node)
	actual, _ := sut.Find(node.value)

	assert.Equal(t, sut.head, after)
	assert.Equal(t, sut.head.next.value, actual.value)
}

func TestInsertGeneral(t *testing.T) {
	sut := LinkedList2{}

	nodes := make([]Node, 0)
	for i := 0; i < 10; i++ {
		node := Node{nil, nil, i}
		nodes = append(nodes, node)
		sut.AddInTail(node)
	}

	tmp := sut.head
	for _, test := range nodes {
		t.Run("", func(t *testing.T) {
			sut.Insert(tmp, test)

			result := sut.FindAll(tmp.value)
			assert.Equal(t, 2, len(result))
			assert.Equal(t, tmp.value, tmp.next.value)

			tmp = tmp.next.next
		})
	}

}

func TestDeleteOneEmpty(t *testing.T) {
	sut := LinkedList2{}
	sut.Delete(1, false)
	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
}

func TestDeleteOneHeadWithTail(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 1})

	sut.Delete(1, false)
	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
}

func TestDeleteOneHeadWithoutTail(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 2})

	sut.Delete(1, false)
	assert.Equal(t, sut.head, sut.tail)
	assert.Equal(t, sut.head.value, 2)
}

func TestDeleteOneTail(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 2})

	sut.Delete(2, false)
	assert.Equal(t, sut.head, sut.tail)
	assert.Equal(t, sut.tail.value, 1)
}

func TestDeleteOnetTail(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 2})

	sut.Delete(2, false)
	assert.Equal(t, sut.head, sut.tail)
	assert.Equal(t, sut.tail.value, 1)
}

func TestDeleteOneGeneral(t *testing.T) {
	sut := LinkedList2{}

	nodes := make([]Node, 0)
	for i := 0; i < 10; i++ {
		node := Node{nil, nil, i}
		nodes = append(nodes, node)
		sut.AddInTail(node)
	}

	nodes = nodes[1 : len(nodes)-1]

	for _, node := range nodes {
		t.Run("", func(t *testing.T) {
			sut.Delete(node.value, false)

			res := sut.FindAll(node.value)

			assert.Empty(t, res)
		})
	}

	assert.Equal(t, sut.head.next, sut.tail)
	assert.Equal(t, sut.tail.prev, sut.head)
}

func TestDeleteOneWithRepetition(t *testing.T) {
	sut := LinkedList2{}

	nodes := make([]Node, 0)
	for i := 0; i < 10; i++ {
		node := Node{nil, nil, 1}
		nodes = append(nodes, node)
		sut.AddInTail(node)
	}

	for _, node := range nodes {
		t.Run("", func(t *testing.T) {
			sut.Delete(node.value, false)

			res := sut.FindAll(node.value)
			if sut.head != nil {
				assert.NotEmpty(t, res)
			}
		})
	}

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
}

func TestDeleteOneCustom(t *testing.T) {
	sut := LinkedList2{}

	nodes := []Node{
		{
			nil, nil, 1,
		},
		{
			nil, nil, 2,
		},
		{
			nil, nil, 1,
		},
		{
			nil, nil, 3,
		},
	}

	for _, node := range nodes {
		sut.AddInTail(node)
	}

	sut.Delete(1, false)
	assert.Equal(t, 3, sut.tail.value)
	assert.Equal(t, 2, sut.head.value)

	sut.Delete(3, false)
	assert.Equal(t, 1, sut.tail.value)
	assert.Equal(t, 2, sut.head.value)
}

func TestDeleteAllEmpty(t *testing.T) {
	sut := LinkedList2{}

	sut.Delete(1, true)

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
}

func TestDeleteAllWithSingleHead(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 1})

	sut.Delete(1, true)

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
}

func TestDeleteAllWithManyEqualHead(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})

	sut.Delete(1, true)

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
}

func TestDeleteAllWithManyEqualHeadAndDifferentTail(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 2})

	sut.Delete(1, true)

	assert.Equal(t, sut.head, sut.tail)
	assert.Equal(t, sut.head.value, 2)
}

func TestDeleteAllWithManyAndDifferentBody(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 2})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 3})

	sut.Delete(1, true)

	assert.Equal(t, sut.head.next, sut.tail)
	assert.Equal(t, sut.tail.prev, sut.head)
	assert.Equal(t, sut.head.value, 2)
	assert.Equal(t, sut.tail.value, 3)
}

func TestDeleteAllWithManyTailsAndDifferentHead(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 2})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})

	sut.Delete(1, true)

	assert.NotNil(t, sut.head)
	assert.Equal(t, sut.head, sut.tail)
}

func TestDeleteAllWithTail(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 2})
	sut.AddInTail(Node{nil, nil, 3})
	sut.AddInTail(Node{nil, nil, 1})

	sut.Delete(1, true)

	assert.Equal(t, 2, sut.head.value)
	assert.Equal(t, 3, sut.tail.value)

	assert.Equal(t, sut.head.next, sut.tail)
	assert.Equal(t, sut.tail.prev, sut.head)
}

func TestDeleteAllGeneral(t *testing.T) {
	sut := LinkedList2{}
	sut.AddInTail(Node{nil, nil, 2})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 1})
	sut.AddInTail(Node{nil, nil, 2})

	sut.Delete(1, true)

	assert.Equal(t, 2, sut.head.value)
	assert.Equal(t, 2, sut.tail.value)

	assert.Equal(t, sut.head.next, sut.tail)
	assert.Equal(t, sut.tail.prev, sut.head)
}

func TestDeleteAllGeneralCustom(t *testing.T) {
	sut := LinkedList2{}
	nodes := []Node{
		{
			nil, nil, 1,
		},
		{
			nil, nil, 2,
		},
		{
			nil, nil, 1,
		},
		{
			nil, nil, 3,
		},
		{
			nil, nil, 1,
		},
		{
			nil, nil, 2,
		},
		{
			nil, nil, 1,
		},
		{
			nil, nil, 3,
		},
	}

	for _, node := range nodes {
		sut.AddInTail(node)
	}

	for _, node := range nodes {
		t.Run("", func(t *testing.T) {
			sut.Delete(node.value, true)

			res := sut.FindAll(node.value)

			assert.Empty(t, res)
		})
	}

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.tail)
}

func TestDeleteAllGeneralCustom1(t *testing.T) {
	sut := LinkedList2{}
	nodes := []Node{
		{
			nil, nil, 3,
		},
		{
			nil, nil, 2,
		},
		{
			nil, nil, 1,
		},
		{
			nil, nil, 3,
		},
		{
			nil, nil, 1,
		},
		{
			nil, nil, 2,
		},
		{
			nil, nil, 1,
		},
		{
			nil, nil, 3,
		},
	}

	for _, node := range nodes {
		sut.AddInTail(node)
	}

	sut.Delete(3, true)
	assert.Equal(t, 2, sut.head.value)
	assert.Equal(t, 1, sut.tail.value)
}
