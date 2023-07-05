package list

import "fmt"

type SinglyLinkedList[T any] struct {
	size int
	head *slnode[T]
	tail *slnode[T]
}

type slnode[T any] struct {
	v    T
	next *slnode[T]
}

// Get - O(N)
func (l *SinglyLinkedList[T]) Get(i int) T {
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
func (l *SinglyLinkedList[T]) GetAll() []T {
	vs := make([]T, 0, l.size)
	for n := l.head; n != nil; n = n.next {
		vs = append(vs, n.v)
	}
	return vs
}

// PushBack - O(1)
func (l *SinglyLinkedList[T]) PushBack(v T) {
	if l.IsEmpty() {
		l.head = &slnode[T]{v: v, next: nil}
		l.tail = l.head
	} else {
		l.tail.next = &slnode[T]{v: v, next: nil}
		l.tail = l.tail.next
	}
	l.size++
}

// PopBack - O(N)
func (l *SinglyLinkedList[T]) PopBack() T {
	if l.IsEmpty() {
		panic("list is empty")
	}
	tail := l.tail
	if l.tail == l.head {
		l.head, l.tail = nil, nil
	} else {
		curr := l.head
		for ; curr.next != l.tail; curr = curr.next {
		}
		l.tail = curr
		l.tail.next = nil
	}
	l.size--
	return tail.v
}

// PushFront - O(1)
func (l *SinglyLinkedList[T]) PushFront(v T) {
	if l.IsEmpty() {
		l.head = &slnode[T]{v: v, next: nil}
		l.tail = l.head
	} else {
		l.head = &slnode[T]{v: v, next: l.head}
	}
	l.size++
}

// PopFront - O(1)
func (l *SinglyLinkedList[T]) PopFront() T {
	if l.IsEmpty() {
		panic("list is empty")
	}
	head := l.head
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return head.v
}

// Invert - O(N)
func (l *SinglyLinkedList[T]) Invert() {
	if l.tail == l.head {
		return
	}
	l.tail = l.head
	for prev, curr := (*slnode[T])(nil), l.head; curr != nil; prev, curr, curr.next = curr, curr.next, prev {
		l.head = curr
	}
}

// Size - O(1)
func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

// IsEmpty - O(1)
func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.head == nil
}
