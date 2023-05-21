package helpers

type PriorityQueue[T any] struct {
	comp  func(l, r T) bool
	elems *Elems[T]
}

func NewPriorityQueue[T any](comp func(l, r T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		comp:  comp,
		elems: &Elems[T]{},
	}
}

type Elems[T any] struct {
	xs []*Elem[T]
}

type Elem[T any] struct {
	Value T
	id    int
}

func (e *Elem[T]) Index() int {
	return e.id
}

func (p PriorityQueue[T]) Len() int {
	return len(p.elems.xs)
}

func (p PriorityQueue[T]) Less(i, j int) bool {
	return p.comp(p.elems.xs[i].Value, p.elems.xs[j].Value)
}

func (p PriorityQueue[T]) Swap(i, j int) {
	p.elems.xs[i], p.elems.xs[j] = p.elems.xs[j], p.elems.xs[i]
	p.elems.xs[i].id, p.elems.xs[j].id = i, j
}

func (p PriorityQueue[T]) Push(x any) {
	elem := x.(*Elem[T])
	elem.id = len(p.elems.xs)
	p.elems.xs = append(p.elems.xs, elem)
}

func (p PriorityQueue[T]) Pop() any {
	elem := p.elems.xs[len(p.elems.xs)-1]
	elem.id = -1
	p.elems.xs = p.elems.xs[:len(p.elems.xs)-1]
	return elem
}
