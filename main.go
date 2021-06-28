package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/imakiri/ps/barrier"
	"github.com/imakiri/ps/particle"
	"github.com/imakiri/ps/reactor"
	"github.com/imakiri/ps/trace"
	"github.com/imakiri/ps/vector"
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
	"log"
	"math"
)

const c float64 = 0.01

var g float64 = math.Sqrt(2*c) * 50000
var l float64 = -math.Sqrt(2 * c)

type sandbox struct {
	particle *particle.GravityParticle
	barriers []barrier.Barrier
	path     *trace.Trace
	reactor.Reactor
}

func (s *sandbox) tick() {
	s.path.Push(s.particle.Position)
	s.particle.React(c, g)
	s.particle.Act(c)
	//s.React(s.particle, s.barriers)
}

const screenWidth = 1280
const screenHeight = 720

func (s *sandbox) Update() error {
	return nil
}

func (s *sandbox) Draw(screen *ebiten.Image) {
	for i := range s.barriers {
		ebitenutil.DrawLine(screen, s.barriers[i][0][0], s.barriers[i][0][1], s.barriers[i][1][0], s.barriers[i][1][1], color.White)
	}
	var path = s.path.Get()
	var l = len(path)
	for i := l - 1; i >= 0; i-- {
		if path[i] != nil {
			ebitenutil.DrawRect(screen, path[i][0]-1., path[i][1]-1., 2., 2., colorful.Hsv(300*float64(i)/float64(l), 1, 1))
		}
	}

	ebitenutil.DrawRect(screen, s.particle.Position[0]-1., s.particle.Position[1]-1., 2., 2., color.White)
	s.tick()

	fmt.Printf("x:%.0f, y:%.0f\n", s.particle.Position[0], s.particle.Position[1])
}

func (s *sandbox) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	var s = &sandbox{
		particle: &particle.GravityParticle{
			Particle: particle.Particle{
				Position: &vector.Vector{640, 180},
				Velocity: &vector.Vector{120, -80},
			},
			Attractor: &vector.Vector{640, 360},
		},
		barriers: []barrier.Barrier{
			{&vector.Vector{0, 0}, &vector.Vector{screenWidth, 0}},
			{&vector.Vector{screenWidth, 0}, &vector.Vector{screenWidth, screenHeight}},
			{&vector.Vector{screenWidth, screenHeight}, &vector.Vector{0, screenHeight}},
			{&vector.Vector{0, screenHeight}, &vector.Vector{0, 0}},
			{&vector.Vector{50, 50}, &vector.Vector{200, 200}},
			{&vector.Vector{200, 200}, &vector.Vector{400, 50}},
			{&vector.Vector{400, 50}, &vector.Vector{50, 50}},
			//{&vector.Vector{1000, 50}, &vector.Vector{800, 600}},
		},
		path: trace.NewTrace(500),
	}

	s.Reactor = reactor.NewGravityPointReactor(g, c, &vector.Vector{640, 360})

	if err := ebiten.RunGame(s); err != nil {
		log.Fatal(err)
	}
}
