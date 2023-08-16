package main

import "errors"

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail
	}

	l.tail = &item
	l.tail.next = nil
}

func (l *LinkedList2) Count() int {
	tmp := l.head
	count := 0
	for tmp != nil {
		count++
		tmp = tmp.next
	}
	return count
}

// error не nil, если узел не найден
func (l *LinkedList2) Find(n int) (Node, error) {
	res := Node{value: -1, next: nil}
	found := false

	tmp := l.head
	for !found && tmp != nil {
		if tmp.value == n {
			res = *tmp
			found = true
		}
		tmp = tmp.next
	}

	if !found {
		return res, errors.New("not found")
	}

	return res, nil
}

func (l *LinkedList2) FindAll(n int) []Node {
	nodes := make([]Node, 0)
	tmp := l.head

	for tmp != nil {
		if tmp.value == n {
			nodes = append(nodes, *tmp)
		}
		tmp = tmp.next
	}

	return nodes
}

func (l *LinkedList2) Delete(n int, all bool) {

}

func (l *LinkedList2) Insert(after *Node, add Node) {

}

func (l *LinkedList2) InsertFirst(first Node) {

}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
}

func iter(pred func() bool, act func()) {

}
