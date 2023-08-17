package main

import (
	"errors"
)

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

var deleteF map[bool]func(int, *LinkedList2) = map[bool]func(int, *LinkedList2){
	true:  deleteAll,
	false: deleteOne,
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
	pred := func() bool { return tmp != nil }
	act := func() {
		count++
		tmp = tmp.next
	}

	iter(pred, act)
	return count
}

// error не nil, если узел не найден
func (l *LinkedList2) Find(n int) (Node, error) {
	res := Node{value: -1, next: nil}
	found := false
	tmp := l.head

	pred := func() bool { return !found && tmp != nil }
	act := func() {
		if tmp.value == n {
			res = *tmp
			found = true
		}
		tmp = tmp.next
	}

	iter(pred, act)
	if !found {
		return res, errors.New("not found")
	}

	return res, nil
}

func (l *LinkedList2) FindAll(n int) []Node {
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

func (l *LinkedList2) Delete(n int, all bool) {
	if l.head == nil {
		return
	}

	f := deleteF[all]
	f(n, l)
}

func (l *LinkedList2) Insert(after *Node, add Node) {
	if after == l.tail {
		add.prev = after
		after.next = &add
		l.tail = &add
		return
	}

	tmp := l.head
	for tmp.value != after.value {
		tmp = tmp.next
	}

	add.next = tmp.next
	tmp.next.prev = &add
	tmp.next = &add
	add.prev = tmp
}

func (l *LinkedList2) InsertFirst(first Node) {
	if l.head == nil {
		l.head = &first
		l.tail = &first
		return
	}

	l.head.prev = &first
	first.next = l.head
	l.head = &first
}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
}

// pretty bad code
func deleteOne(n int, l *LinkedList2) {
	if l.head == nil {
		return
	}

	if l.head.value == n && l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	if l.head.value == n {
		l.head = l.head.next
		l.head.prev = nil
		return
	}

	tmp := l.head

	// pretty bad, but i still can't find out
	// how to make it without nested if's
	for tmp != nil {
		if tmp.value == n {
			if tmp == l.tail {
				l.tail = l.tail.prev
				l.tail.next = nil
				return
			} else {
				tmp.next.prev = tmp.prev
				tmp.prev.next = tmp.next
				return
			}
		}
		tmp = tmp.next
	}

}

func deleteAll(n int, l *LinkedList2) {
	deleteAllHead(n, l)
	deleteAllTail(n, l)
	deleteAllBody(n, l)
}

func deleteAllBody(n int, l *LinkedList2) {

	tmp := l.head
	pred := func() bool { return tmp != nil }
	act := func() {
		if tmp.value == n {
			tmp.next.prev = tmp.prev
			tmp.prev.next = tmp.next
		}
		tmp = tmp.next
	}

	iter(pred, act)
}

func deleteAllHead(n int, l *LinkedList2) {
	pred := func() bool {
		return l.head != l.tail && l.head.value == n
	}

	act := func() {
		l.head = l.head.next
		if l.head.prev != nil {
			l.head.prev = nil
		}
	}

	iter(pred, act)

	if l.head == l.tail && l.head.value == n {
		l.head, l.tail = nil, nil
		return
	}
}

func deleteAllTail(n int, l *LinkedList2) {
	pred := func() bool { return l.tail != nil && l.tail.value == n }
	act := func() {
		l.tail = l.tail.prev
		if l.tail.next != nil {
			l.tail.next = nil
		}
	}

	iter(pred, act)
}

func iter(pred func() bool, act func()) {
	for pred() {
		act()
	}
}
