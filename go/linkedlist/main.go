package main

import (
	"errors"
)

type Node struct {
	next  *Node
	value int
}

var deleteType map[bool]func(l *LinkedList, n int) = map[bool]func(l *LinkedList, n int){
	true:  deleteAll,
	false: deleteOne,
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
}

func (l *LinkedList) Count() int {
	tmp := l.head
	count := 0

	pred := func() bool { return tmp != nil }
	act := func() {
		count++
		tmp = tmp.next
	}

	iter(pred, act)
	return count
}

// error не nil, если узел не найден
func (l *LinkedList) Find(n int) (Node, error) {
	res := &Node{value: -1, next: nil}

	tmp := l.head
	found := false

	pred := func() bool { return tmp != nil && !found }
	act := func() {
		if tmp.value == n {
			res = tmp
			found = true
		}
		tmp = tmp.next
	}

	iter(pred, act)

	if !found {
		return *res, errors.New("not found")
	}
	return *res, nil
}

func (l *LinkedList) FindAll(n int) []Node {
	nodes := make([]Node, 0)
	tmp := l.head

	pred := func() bool { return tmp != nil }
	act := func() {
		if tmp.value == n {
			nodes = append(nodes, *tmp)
		}
		tmp = tmp.next
	}

	iter(pred, act)
	return nodes
}

func (l *LinkedList) Delete(n int, all bool) {
	if l.head == nil {
		return
	}

	deleteF := deleteType[all]
	deleteF(l, n)
}

func (l *LinkedList) Insert(after *Node, add Node) {
	tmp := l.head
	pred := func() bool { return tmp.value != add.value }
	act := func() { tmp = tmp.next }

	iter(pred, act)

	add.next = tmp.next
	tmp.next = &add
}

func (l *LinkedList) InsertFirst(first Node) {
	first.next = l.head
	l.head = &first
}

func (l *LinkedList) Clean() {
	l.head, l.tail = nil, nil
}

func deleteOne(l *LinkedList, n int) {

	if l.head.value == n {
		l.head = l.head.next
	}

	if l.head == nil {
		l.tail = nil
		return
	}

	tmp := l.head
	deleted := false

	pred := func() bool { return !deleted && tmp.next != nil }
	act := func() {

		if tmp.next.value == n {
			if tmp.next == l.tail {
				l.tail = tmp
				l.tail.next = nil
			} else {
				tmp.next = tmp.next.next
			}

			deleted = true
		}

		tmp = tmp.next
	}

	iter(pred, act)
}

func deleteAll(l *LinkedList, n int) {
	deleteAllHead(l, n)
	deleteAllBody(l, n)

}

func deleteAllHead(l *LinkedList, n int) {
	pred := func() bool { return l.head != nil && l.head.value == n }
	act := func() { l.head = l.head.next }

	iter(pred, act)

	if l.head == l.tail && l.head.value == n || l.head == nil {
		l.head = nil
		l.tail = nil
	}

}

func deleteAllBody(l *LinkedList, n int) {
	tmp := l.head
	pred := func() bool { return tmp != nil && tmp.next != nil && tmp.next != l.tail }

	act := func() {
		for tmp.next.value == n && tmp.next != l.tail {
			tmp.next = tmp.next.next
		}

		if tmp.next != l.tail {
			tmp = tmp.next
		}
	}

	iter(pred, act)

	if tmp != nil && tmp.next == l.tail && l.tail.value == n {
		l.tail = tmp
		l.tail.next = nil
	}

}

func iter(cond func() bool, action func()) {
	for cond() {
		action()
	}
}
