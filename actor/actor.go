package actor

import (
	"github.com/imakiri/ps/barrier"
	"github.com/imakiri/ps/particle"
)

type Actor interface {
	Act(p *particle.Particle, bs []barrier.Barrier)
}
