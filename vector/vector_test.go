package vector

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var v = Vector{1, 1, 1}
	var w = v.Copy()
	v.Multiply(10)
	fmt.Println(w)
}
