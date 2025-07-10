package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func main() {
	//x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//A := mat.NewDense(3, 4, x)
	//matPrint(A)
	// ⎡1  2  3  4⎤
	// ⎢5  6  7  8⎥
	// ⎣9 10 11 12⎦
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	A := mat.NewDense(3, 3, x)
	matPrint(A)
	//fmt.Printf("Matrix:\n%v\n", mat.Formatted(A))
	var lu mat.LU
	lu.Factorize(A)
	fmt.Println(lu)
}
