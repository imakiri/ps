package reactor

import (
	"github.com/imakiri/ps/barrier"
	"github.com/imakiri/ps/particle"
	"github.com/imakiri/ps/vector"
)

func NewGravityPointReactor(g, c float64, point *vector.Vector) Reactor {
	var gp = new(gp)
	gp.point = point
	gp.g = g
	gp.c = c

	return gp
}

type gp struct {
	g, c  float64
	point *vector.Vector
}

func (gp *gp) Reactor(p *particle.Particle, bs []barrier.Barrier) {
	var r = vector.Subtract(p.Position, gp.point)
	r.Multiply(-1 / gp.c / gp.g)
	p.Velocity.Add(r)
}
