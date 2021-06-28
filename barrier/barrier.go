package barrier

import "github.com/imakiri/ps/vector"

type Barrier [2]*vector.Vector

func (b Barrier) Side() *vector.Vector {
	return vector.Subtract(b[1], b[0])
}

func (b Barrier) Collision(v *vector.Vector) (float64, float64) {
	var c = b.Side()
	var y = vector.Add(vector.CrossProduct(v, b.Side()), vector.CrossProduct(b[1], b[0]))
	return vector.DotProduct(y, &vector.Vector{0, 0, 1}), vector.DotProduct(vector.Subtract(v, b[0]), c) * vector.DotProduct(vector.Subtract(b[1], v), c)
}
