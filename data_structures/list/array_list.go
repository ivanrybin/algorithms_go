package list

type ArrayList[T any] struct {
	a []T
}

// Get - O(1)
func (l *ArrayList[T]) Get(i int) T {
	return l.a[i]
}

// GetAll - O(N)
func (l *ArrayList[T]) GetAll() []T {
	return append([]T{}, l.a...)
}

// PushBack - amortized O(1)
func (l *ArrayList[T]) PushBack(v T) {
	l.a = append(l.a, v)
}

// PopBack - O(1)
func (l *ArrayList[T]) PopBack() T {
	if len(l.a) == 0 {
		panic("list is empty")
	}
	var v T
	l.a, v = l.a[:len(l.a)-1], l.a[len(l.a)-1]
	return v
}

// PushFront - O(N)
func (l *ArrayList[T]) PushFront(v T) {
	l.a = append([]T{v}, l.a...)
}

// PopFront - O(1)
func (l *ArrayList[T]) PopFront() T {
	if len(l.a) == 0 {
		panic("list is empty")
	}
	var v T
	v, l.a = l.a[0], l.a[1:]
	return v
}

// Invert - O(N)
func (l *ArrayList[T]) Invert() {
	for i := 0; i < len(l.a)/2; i++ {
		l.a[i], l.a[len(l.a)-i] = l.a[len(l.a)-i], l.a[i]
	}
}

// Size - O(1)
func (l *ArrayList[T]) Size() int {
	return len(l.a)
}

// IsEmpty - O(1)
func (l *ArrayList[T]) IsEmpty() bool {
	return len(l.a) == 0
}
