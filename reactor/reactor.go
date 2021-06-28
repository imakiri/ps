package reactor

import (
	"github.com/imakiri/ps/barrier"
	"github.com/imakiri/ps/particle"
)

type Reactor interface {
	Reactor(p *particle.Particle, bs []barrier.Barrier)
}
