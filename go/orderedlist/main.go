package main

import (
	"fmt"
	"constraints"
	"os"
)

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	_ascending bool
}

func (l *OrderedList[T]) Count() int {
	count := 0
	tmp := l.head
	for tmp != nil {
		tmp = tmp.next
		count++
	}
	return count
}

func (l *OrderedList[T]) Add(item T) {
	if l.head == nil {
		l.head = &Node[T]{nil, nil, item}
		return
	}

	if l._ascending {
		addAsc[T](l, item)
	} else {
		addDesc[T](l, item)
	}
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	if l.head == nil {
		return Node[T]{value: n, next: nil, prev: nil}, fmt.Errorf("not found")
	}

	if l._ascending {
		return l.findInAsc(n)
	}
	return l.findInDesc(n)

}

func (l *OrderedList[T]) findInAsc(n T) (Node[T], error) {
	if l.Compare(l.head.value, n) == 0 {
		return *l.head, nil
	}

	var result Node[T]
	tmp := l.head
	for tmp.next != nil && l.Compare(n, tmp.next.value) == 1 {
		tmp = tmp.next
	}

	if tmp.next == nil {
		return result, fmt.Errorf("not found")
	}

	if l.Compare(tmp.next.value, n) == 0 {
		result = *tmp.next
		return result, nil
	}

	return *tmp, fmt.Errorf("not found")

}

func (l *OrderedList[T]) findInDesc(n T) (Node[T], error) {
	if l.Compare(l.head.value, n) == 0 {
		return *l.head, nil
	}

	var result Node[T]
	tmp := l.head
	for tmp.next != nil && l.Compare(n, tmp.next.value) != 0 {
		tmp = tmp.next
	}

	if tmp.next == nil {
		return result, fmt.Errorf("not found")
	}

	if tmp.next.value == n {
		result = *tmp.next
		return result, nil
	}

	return *tmp, fmt.Errorf("not found")
}

func (l *OrderedList[T]) Delete(n T) {
	if l.head == nil {
		return
	}

	if l.Compare(l.head.value, n) == 0 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
		return
	}

	if l._ascending {
		l.deleteAsc(n)
		return
	}

	l.deleteDesc(n)
}

func (l *OrderedList[T]) deleteAsc(n T) {
	tmp := l.head
	pred := func() bool { return tmp != nil && l.Compare(tmp.value, n) != 0 }
	act := func() { tmp = tmp.next }
	l.iter(pred, act)
	if tmp.value == n && tmp.next == nil {
		tmp = tmp.prev
		tmp.next = nil
		return
	}

	if tmp.value == n {
		if tmp.next.prev != nil {
			tmp.next.prev = tmp.prev
		}
		tmp.prev.next = tmp.next
	}
}

func (l *OrderedList[T]) deleteDesc(n T) {
	tmp := l.head
	for tmp != nil && l.Compare(tmp.value, n) != 0 {
		tmp = tmp.next
	}

	if tmp == nil {
		return
	}

	if l.Compare(tmp.value, n) == 0 && tmp.next == nil {
		tmp = tmp.prev
		tmp.next = nil
		return
	}

	if l.Compare(tmp.value, n) == 0 {
		if tmp.next.prev != nil {
			tmp.next.prev = tmp.prev
		}
		tmp.prev.next = tmp.next
	}
}

func (l *OrderedList[T]) Clear(asc bool) {
	l._ascending = asc
	l.head = nil
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}

func addAsc[T constraints.Ordered](l *OrderedList[T], item T) {
	node := &Node[T]{nil, nil, item}
	tmp := l.head
	if l.tryHead(node, l.head, -1) {
		return
	}

	pred := func() bool { return tmp != nil && tmp.next != nil && l.Compare(node.value, tmp.next.value) > -1 }
	act := func() { tmp = tmp.next }
	l.iter(pred, act)

	if l.tryTail(tmp, node) {
		return
	}

	l.tryBody(node, tmp, -1)
}

func addDesc[T constraints.Ordered](l *OrderedList[T], item T) {
	node := &Node[T]{nil, nil, item}
	tmp := l.head
	if l.tryHead(node, l.head, 1) {
		return
	}

	tmp = l.move(tmp, node, 1)

	if l.tryTail(tmp, node) {
		return
	}

	l.tryBody(node, tmp, 1)
}

func (l *OrderedList[T]) insertBeforeHead(n1 *Node[T], n2 *Node[T]) {
	n1.prev = n2
	n2.next = n1
	l.head = n2
}

func (l *OrderedList[T]) insertInBody(n1 *Node[T], n2 *Node[T]) {
	n1.next.prev = n2
	n2.next = n1.next
	n1.next = n2
	n2.prev = n1
}

func (l *OrderedList[T]) insertInTail(n1 *Node[T], n2 *Node[T]) {
	n1.next = n2
	n2.prev = n1
}

func (l *OrderedList[T]) iter(pred func() bool, act func()) {
	for pred() {
		act()
	}
}

func (l *OrderedList[T]) move(curr *Node[T], node *Node[T], expected int) *Node[T] {
	pred := func() bool {
		return curr != nil && curr.next != nil && l.Compare(node.value, curr.next.value) < expected
	}
	act := func() { curr = curr.next }

	l.iter(pred, act)
	return curr
}

func (l *OrderedList[T]) tryHead(compare *Node[T], with *Node[T], expect int) bool {
	mustAdd := l.Compare(compare.value, with.value) == expect
	if mustAdd {
		l.insertBeforeHead(l.head, compare)
	}
	return mustAdd
}

func (l *OrderedList[T]) tryTail(node *Node[T], toAdd *Node[T]) bool {
	mustAdd := node.next == nil
	if mustAdd {
		l.insertInTail(node, toAdd)
	}
	return mustAdd
}

func (l *OrderedList[T]) tryBody(compare *Node[T], with *Node[T], expect int) bool {
	mustAdd := l.Compare(compare.value, with.next.value) == expect
	if mustAdd {
		l.insertInBody(with, compare)
	}
	return mustAdd
}
