package helpers

type PriorityQueue[T any] struct {
	Comp func(l, r T) bool
	xs   []*Elem[T]
}

type Elem[T any] struct {
	Value T
	id    int
}

func (e *Elem[T]) Index() int {
	return e.id
}

func (p PriorityQueue[T]) Len() int {
	return len(p.xs)
}

func (p PriorityQueue[T]) Less(i, j int) bool {
	return p.Comp(p.xs[i].Value, p.xs[j].Value)
}

func (p PriorityQueue[T]) Swap(i, j int) {
	p.xs[i], p.xs[j] = p.xs[j], p.xs[i]
	p.xs[i].id, p.xs[j].id = j, i
}

func (p PriorityQueue[T]) Push(x any) {
	elem := x.(*Elem[T])
	elem.id = len(p.xs)
	p.xs = append(p.xs, elem)
}

func (p PriorityQueue[T]) Pop() any {
	elem := p.xs[len(p.xs)-1]
	p.xs = p.xs[:len(p.xs)-1]
	return elem
}
