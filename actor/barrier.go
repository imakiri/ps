package actor

import (
	"github.com/imakiri/ps/barrier"
	"github.com/imakiri/ps/particle"
)

type b struct {
	collisionFactors [][2]float64
}

func (i *b) Act(p *particle.Particle, bs []barrier.Barrier) {
	var cf [2]float64
	for j := range bs {
		cf[0], cf[1] = bs[j].Collision(p.Position)
		if cf[0]*i.collisionFactors[j][0] <= 0 && cf[1] > 0 {
			p.Reflect(bs[j].Side())
			continue
		}
		i.collisionFactors[j] = cf
	}
}

func NewBarrierActor(particle *particle.Particle, barriers []barrier.Barrier) Actor {
	var i = new(b)
	i.collisionFactors = make([][2]float64, len(barriers))

	var cf [2]float64
	for j := range barriers {
		cf[0], cf[1] = barriers[j].Collision(particle.Position)
		i.collisionFactors[j] = cf
	}

	return i
}
