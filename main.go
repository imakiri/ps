package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/imakiri/ps/vector"
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
	"log"
	"math"
)

const c float64 = 0.005

var g float64 = math.Sqrt(2*c) * 100000000
var l float64 = -math.Sqrt(2 * c)

type particle struct {
	position *vector.Vector
	velocity *vector.Vector
}

func (p *particle) fly() {
	p.position.Add(vector.Multiply(p.velocity, c))
}

func (p *particle) reflect(side *vector.Vector) {
	p.velocity = vector.Subtract(vector.Multiply(side, 2*vector.DotProduct(side, p.velocity)/side.DotSquare()), p.velocity)
}

type barrier [2]*vector.Vector

func (b barrier) side() *vector.Vector {
	return vector.Subtract(b[1], b[0])
}

func (b barrier) collision(v *vector.Vector) (float64, float64) {
	var c = b.side()
	var y = vector.Add(vector.CrossProduct(v, b.side()), vector.CrossProduct(b[1], b[0]))
	return vector.DotProduct(y, &vector.Vector{0, 0, 1}), vector.DotProduct(vector.Subtract(v, b[0]), c) * vector.DotProduct(vector.Subtract(b[1], v), c)
}

func NewActor(particle *particle, barriers []barrier, g *vector.Vector) Actor {
	var i = new(interaction)
	i.collisionFactors = make([][2]float64, len(barriers))
	i.g = g
	i.f = &vector.Vector{640, 360}

	var cf [2]float64
	for j := range barriers {
		cf[0], cf[1] = barriers[j].collision(particle.position)
		i.collisionFactors[j] = cf
	}

	i.g.Multiply(c)
	return i
}

type interaction struct {
	collisionFactors [][2]float64
	g                *vector.Vector
	f                *vector.Vector
}

func (i *interaction) Act(p *particle, bs []barrier) {
	var cf [2]float64
	var gr = true
	for j := range bs {
		cf[0], cf[1] = bs[j].collision(p.position)
		if cf[0]*i.collisionFactors[j][0] <= 0 && cf[1] > 0 {
			p.reflect(bs[j].side())
			gr = false
			continue
		}
		i.collisionFactors[j] = cf
	}

	if gr {
		p.velocity.Add(i.g)

		var r = vector.Subtract(i.f, p.position)
		var v = p.velocity.Copy()
		p.velocity.Add(vector.Multiply(r, g*c/(r.DotSquare()*vector.Add(r, vector.Multiply(v, c)).Len())))
	}
}

type Actor interface {
	Act(p *particle, bs []barrier)
}

func NewPath(size int) *path {
	var p = new(path)
	p.path = make([]*vector.Vector, size)
	p.size = size
	return p
}

type path struct {
	path []*vector.Vector
	head int
	size int
}

func (p *path) push(v *vector.Vector) {
	p.path[p.head] = v.Copy()
	if p.head+1 == p.size {
		p.head = 0
	} else {
		p.head++
	}
}

func (p path) get() []*vector.Vector {
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

type sandbox struct {
	particle *particle
	barriers []barrier
	path     *path
	Actor
}

func (s *sandbox) tick() {
	s.path.push(s.particle.position)
	s.particle.fly()
	s.Act(s.particle, s.barriers)
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
	var path = s.path.get()
	var l = len(path)
	for i := l - 1; i >= 0; i-- {
		if path[i] != nil {
			ebitenutil.DrawRect(screen, path[i][0]-1., path[i][1]-1., 2., 2., colorful.Hsv(300*float64(i)/float64(l), 1, 1))
		}
	}

	ebitenutil.DrawRect(screen, s.particle.position[0]-1., s.particle.position[1]-1., 2., 2., color.White)
	s.tick()

	fmt.Printf("x:%.0f, y:%.0f\n", s.particle.position[0], s.particle.position[1])
}

func (s *sandbox) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	var s = &sandbox{
		particle: &particle{
			position: &vector.Vector{640, 180},
			velocity: &vector.Vector{120, -80},
		},
		barriers: []barrier{
			{&vector.Vector{0, 0}, &vector.Vector{screenWidth, 0}},
			{&vector.Vector{screenWidth, 0}, &vector.Vector{screenWidth, screenHeight}},
			{&vector.Vector{screenWidth, screenHeight}, &vector.Vector{0, screenHeight}},
			{&vector.Vector{0, screenHeight}, &vector.Vector{0, 0}},
			{&vector.Vector{50, 50}, &vector.Vector{200, 200}},
			{&vector.Vector{200, 200}, &vector.Vector{400, 50}},
			{&vector.Vector{400, 50}, &vector.Vector{50, 50}},
			//{&vector.Vector{1000, 50}, &vector.Vector{800, 600}},
		},
		path: NewPath(500),
	}

	s.Actor = NewActor(s.particle, s.barriers, &vector.Vector{0, 0})

	if err := ebiten.RunGame(s); err != nil {
		log.Fatal(err)
	}

	//	var stdin = os.Stdin
	//	var stdout = os.Stdout
	//	var input = make([]byte, 10)
	//	var output = make([]byte, 100)
	//
	//	for range [3]int{} {
	//		var n, err = stdin.Read(input)
	//		fmt.Printf("Read %v, %v, % x\n", n, err, input[:n])
	//	}
	//
	//	for range [1]int{} {
	//		var n, err = stdout.Read(output)
	//		fmt.Printf("Read %v, %v, % x\n", n, err, output[:n])
	//	}
	//
	//	time.Sleep(time.Second)
}
