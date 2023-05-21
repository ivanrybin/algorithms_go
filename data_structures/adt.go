package data_structures

type Queue[T any] interface {
	Push(v T)
	Pop() T
	Size() int
}

type Stack[T any] interface {
	Push(v T)
	Pop() T
	Size() int
}

type List[T any] interface {
	Get(i int) T
	GetAll() []T
	PushBack(v T)
	PopBack() T
	PushFront(v T)
	PopFront() T
	Invert()
	Size() int
	IsEmpty() bool
}

type DisjointSet[T any] interface {
	MakeSet(x T)
	Find(x T) T
	Union(x, y T)
}

type DisjointSetInt = DisjointSet[int]
