package vector

import "math"

func lc(a, b, c int) int8 {
	var x = [3]int{a, b, c}
	switch x {
	case [3]int{0, 1, 2}, [3]int{2, 0, 1}, [3]int{1, 2, 0}:
		return 1
	case [3]int{2, 1, 0}, [3]int{0, 2, 1}, [3]int{1, 0, 3}:
		return -1
	default:
		return 0
	}
}

// 0: x, 1: y, 2: z
type Vector [3]float64

var vector Vector

func (co *Vector) Normalized() *Vector {
	var ds = co.DotSquare()
	if ds != 0 {
		return Multiply(co, 1/math.Sqrt(ds))
	} else {
		return &Vector{0, 0, 0}
	}

}

func (co *Vector) Len() float64 {
	return math.Sqrt(co.DotSquare())
}

func (co *Vector) Add(c *Vector) {
	for i := range vector {
		co[i] += c[i]
	}
}

func (co *Vector) Subtract(c *Vector) {
	for i := range vector {
		co[i] -= c[i]
	}
}

func (co *Vector) Multiply(factor float64) {
	for i := range vector {
		co[i] *= factor
	}
}

func (co *Vector) DotSquare() float64 {
	return DotProduct(co, co)
}

func (co Vector) Copy() *Vector {
	return &co
}

func Add(v *Vector, w *Vector) (y *Vector) {
	y = new(Vector)
	for i := range vector {
		y[i] = v[i] + w[i]
	}
	return
}

func Subtract(v *Vector, w *Vector) (y *Vector) {
	y = new(Vector)
	for i := range vector {
		y[i] = v[i] - w[i]
	}
	return
}

func Multiply(v *Vector, factor float64) (y *Vector) {
	y = new(Vector)
	for i := range vector {
		y[i] = factor * v[i]
	}
	return
}

func DotProduct(v *Vector, w *Vector) (re float64) {
	for i := range vector {
		re += v[i] * w[i]
	}
	return
}

func CrossProduct(v *Vector, w *Vector) (y *Vector) {
	y = new(Vector)
	for i := range vector {
		for j := range vector {
			for k := range vector {
				y[i] += float64(lc(i, j, k)) * v[j] * w[k]
			}
		}
	}
	return
}
