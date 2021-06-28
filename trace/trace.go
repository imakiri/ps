package trace

import "github.com/imakiri/ps/vector"

func NewTrace(size int) *Trace {
	var p = new(Trace)
	p.path = make([]*vector.Vector, size)
	p.size = size
	return p
}

type Trace struct {
	path []*vector.Vector
	head int
	size int
}

func (p *Trace) Push(v *vector.Vector) {
	p.path[p.head] = v.Copy()
	if p.head+1 == p.size {
		p.head = 0
	} else {
		p.head++
	}
}

func (p Trace) Get() []*vector.Vector {
	var pa = make([]*vector.Vector, p.size)
	var h = p.head
	for i := 0; i < p.size; i++ {
		pa[i] = p.path[h]
		if h == 0 {
			h = p.size - 1
		} else {
			h--
		}
	}
	return pa
}
