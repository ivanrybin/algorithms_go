package stack_list

import (
	"github.com/ivanrybin/algorithms_go/data_structures"
	"github.com/ivanrybin/algorithms_go/data_structures/list"
)

type Push[T any] func(T)

type Pop[T any] func() T

type Size func() int

type Stack[T any, L data_structures.List[T]] struct {
	l    L
	push Push[T]
	pop  Pop[T]
	size Size
}

func NewStack[T any](l data_structures.List[T], push Push[T], pop Pop[T]) *Stack[T, data_structures.List[T]] {
	return &Stack[T, data_structures.List[T]]{
		l:    l,
		push: push,
		pop:  pop,
		size: l.Size,
	}
}

func NewArrayListStack[T any]() *Stack[T, data_structures.List[T]] {
	l := &list.ArrayList[T]{}
	return NewStack[T](l, l.PushBack, l.PopBack)
}

func NewSinglyLinkedListStack[T any]() *Stack[T, data_structures.List[T]] {
	l := &list.SinglyLinkedList[T]{}
	return NewStack[T](l, l.PushFront, l.PopFront)
}

func NewDoublyLinkedListStack[T any]() *Stack[T, data_structures.List[T]] {
	l := &list.DoublyLinkedList[T]{}
	return NewStack[T](l, l.PushFront, l.PopFront)
}

// Push
//
//	O(1) singly/doubly linked list
//	amortized O(1) array list.
func (s *Stack[T, L]) Push(v T) {
	s.push(v)
}

// Pop
//
//	O(1) singly/doubly linked list
//	O(1) array list.
func (s *Stack[T, L]) Pop() T {
	return s.pop()
}

// Size O(1).
func (s *Stack[T, L]) Size() int {
	return s.size()
}
