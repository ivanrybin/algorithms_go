package list

import "fmt"

type DoublyLinkedList[T any] struct {
	size int
	head *dlnode[T]
	tail *dlnode[T]
}

type dlnode[T any] struct {
	v    T
	next *dlnode[T]
	prev *dlnode[T]
}

// Get - O(N)
func (l *DoublyLinkedList[T]) Get(i int) T {
	if i < 0 || i >= l.size {
		panic(fmt.Sprintf("out of bound: %d >= %d", i, l.size))
	}
	curr := l.head
	for j := 0; j < i; j++ {
		curr = curr.next
	}
	return curr.v
}

// GetAll - O(N)
func (l *DoublyLinkedList[T]) GetAll() []T {
	vs := make([]T, 0, l.size)
	for n := l.head; n != nil; n = n.next {
		vs = append(vs, n.v)
	}
	return vs
}

// PushBack - O(1)
func (l *DoublyLinkedList[T]) PushBack(v T) {
	if l.IsEmpty() {
		l.head = &dlnode[T]{v: v, next: nil, prev: nil}
		l.tail = l.head
	} else {
		l.tail.next = &dlnode[T]{v: v, next: nil, prev: l.tail}
		l.tail = l.tail.next
	}
	l.size++
}

// PopBack - O(1)
func (l *DoublyLinkedList[T]) PopBack() T {
	if l.IsEmpty() {
		panic("list is empty")
	}
	tail := l.tail
	l.tail = l.tail.prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	l.size--
	return tail.v
}

// PushFront - O(1)
func (l *DoublyLinkedList[T]) PushFront(v T) {
	if l.IsEmpty() {
		l.head = &dlnode[T]{v: v, next: nil, prev: nil}
		l.tail = l.head
	} else {
		node := &dlnode[T]{v: v, next: l.head, prev: nil}
		l.head, l.head.prev = node, node
	}
	l.size++
}

// PopFront - O(1)
func (l *DoublyLinkedList[T]) PopFront() T {
	if l.IsEmpty() {
		panic("list is empty")
	}
	head := l.head
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	l.size--
	return head.v
}

// Invert - O(N)
func (l *DoublyLinkedList[T]) Invert() {
	if l.head == l.tail {
		return
	}
	for curr := l.head; curr != nil; curr, curr.prev, curr.next = curr.next, curr.next, curr.prev {
	}
	l.tail, l.head = l.head, l.tail
}

// Size - O(1)
func (l *DoublyLinkedList[T]) Size() int {
	return l.size
}

// IsEmpty - O(1)
func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.head == nil
}
