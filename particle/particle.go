package particle

import (
	"github.com/imakiri/ps/vector"
)

type Particle struct {
	Position *vector.Vector
	Velocity *vector.Vector
}

func (p *Particle) Act(c float64) {
	p.Position.Add(vector.Multiply(p.Velocity, c))
}

func (p *Particle) Reflect(side *vector.Vector) {
	p.Velocity = vector.Subtract(vector.Multiply(side, 2*vector.DotProduct(side, p.Velocity)/side.DotSquare()), p.Velocity)
}

type GravityParticle struct {
	Particle
	Attractor *vector.Vector
}

func (gp *GravityParticle) React(c, g float64) {
	var r = vector.Subtract(gp.Position, gp.Attractor)
	r.Multiply(-1 / c / g)
	gp.Velocity.Add(r)
}
