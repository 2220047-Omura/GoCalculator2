package main

import (
	"gonum.org/v1/gonum/mat"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	A := mat.NewDense(3, 4, x)
	// ⎡1  2  3  4⎤
	// ⎢5  6  7  8⎥
	// ⎣9 10 11 12⎦
}
