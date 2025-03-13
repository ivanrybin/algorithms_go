package helpers

type PriorityQueue[T any] struct {
	less  func(l, r T) bool
	elems []*Elem[T]
}

func NewPriorityQueue[T any](less func(l, r T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		less:  less,
		elems: []*Elem[T]{},
	}
}

type Elem[T any] struct {
	Value T
	id    int
}

func (e *Elem[T]) Index() int {
	return e.id
}

func (p *PriorityQueue[T]) Len() int {
	return len(p.elems)
}

func (p *PriorityQueue[T]) Less(i, j int) bool {
	return p.less(p.elems[i].Value, p.elems[j].Value)
}

func (p *PriorityQueue[T]) Swap(i, j int) {
	p.elems[i], p.elems[j] = p.elems[j], p.elems[i]
	p.elems[i].id, p.elems[j].id = i, j
}

func (p *PriorityQueue[T]) Push(x any) {
	elem := x.(*Elem[T])
	elem.id = len(p.elems)
	p.elems = append(p.elems, elem)
}

func (p *PriorityQueue[T]) Pop() any {
	elem := p.elems[len(p.elems)-1]
	elem.id = -1
	p.elems = p.elems[:len(p.elems)-1]
	return elem
}
