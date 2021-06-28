package actor

import (
	"github.com/imakiri/ps/barrier"
	"github.com/imakiri/ps/particle"
	"github.com/imakiri/ps/vector"
)

func NewGravityPointActor(g, c float64, point *vector.Vector) Actor {
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

func (i *gp) Act(p *particle.Particle, bs []barrier.Barrier) {
	var r = vector.Subtract(i.point, p.Position)
	var v = p.Velocity.Copy()
	p.Velocity.Add(vector.Multiply(r, i.g*i.c/(r.DotSquare()*vector.Add(r, vector.Multiply(v, i.c)).Len())))
}
