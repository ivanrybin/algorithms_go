package stack_list

import (
	"ivanrybin/algorithms_go/data_structures"
	"ivanrybin/algorithms_go/data_structures/list"
)

type Stack[T any, L data_structures.List[T]] struct {
	l    L
	push func(T)
	pop  func() T
	size func() int
}

func NewArrayListStack[T any]() *Stack[T, data_structures.List[T]] {
	l := &list.ArrayList[T]{}
	return &Stack[T, data_structures.List[T]]{
		l:    l,
		push: l.PushBack,
		pop:  l.PopBack,
		size: l.Size,
	}
}

func NewSinglyLinkedListStack[T any]() *Stack[T, data_structures.List[T]] {
	l := &list.SinglyLinkedList[T]{}
	return &Stack[T, data_structures.List[T]]{
		l:    l,
		push: l.PushFront,
		pop:  l.PopFront,
		size: l.Size,
	}
}

func NewDoublyLinkedListStack[T any]() *Stack[T, data_structures.List[T]] {
	l := &list.DoublyLinkedList[T]{}
	return &Stack[T, data_structures.List[T]]{
		l:    l,
		push: l.PushFront,
		pop:  l.PopFront,
		size: l.Size,
	}
}

// Push
// - O(1) singly/doubly linked list
// - amortized O(1) array list
func (s *Stack[T, L]) Push(v T) {
	s.push(v)
}

// Pop
// - O(1) singly/doubly linked list
// - O(1) array list
func (s *Stack[T, L]) Pop() T {
	return s.pop()
}

// Size - O(1)
func (s *Stack[T, L]) Size() int {
	return s.size()
}
