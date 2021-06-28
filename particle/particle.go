package particle

import (
	"github.com/imakiri/ps/vector"
)

type Particle struct {
	Position *vector.Vector
	Velocity *vector.Vector
}

func (p *Particle) Fly(c float64) {
	p.Position.Add(vector.Multiply(p.Velocity, c))
}

func (p *Particle) Reflect(side *vector.Vector) {
	p.Velocity = vector.Subtract(vector.Multiply(side, 2*vector.DotProduct(side, p.Velocity)/side.DotSquare()), p.Velocity)
}
