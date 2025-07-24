package main

import (
	"fmt"
	"sync"
	"time"

	"gonum.org/v1/gonum/mat"
)

var wg sync.WaitGroup

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func Uset(A *mat.Dense, L *mat.Dense, U *mat.Dense, Lch *[3][3]chan float64, Uch *[3][3]chan float64, i int, j int) {
	defer wg.Done()
	r, _ := A.Dims()
	var aij, uij, lij float64
	for k := 0; k < r; k++ {
		lij = <-Lch[i][k]
		Lch[i][k] <- lij
		if k == i || lij == 0 {
			uij = 0
		} else {
			uij = <-Uch[k][j]
			Uch[k][j] <- uij
		}
		/*
			if k != i {
				uij = <-Uch[k][j]
				Uch[k][j] <- uij
			} else {
				uij = 0
			}
		*/
		aij += lij * uij
	}
	uij = A.At(i, j) - aij
	U.Set(i, j, uij)
	Uch[i][j] <- uij
}

func Lset(A *mat.Dense, L *mat.Dense, U *mat.Dense, Lch *[3][3]chan float64, Uch *[3][3]chan float64, i int, j int) {
	defer wg.Done()
	if i != j {
		r, _ := A.Dims()
		var aji, uji, lji float64
		for k := 0; k < r; k++ {
			uji = <-Uch[k][i]
			Uch[k][i] <- uji
			if k == i || uji == 0 {
				lji = 0
			} else {
				lji = <-Lch[j][k]
				Lch[j][k] <- lji
			}
			/*
				if k != i {
					lji = <-Lch[j][k]
					Lch[j][k] <- lji
				} else {
					lji = 0
				}
			*/
			aji += lji * uji
		}
		uii := <-Uch[i][i]
		Uch[i][i] <- uii
		if uii == 0 {
			lji = 0
		} else {
			lji = (A.At(j, i) - aji) / uii
		}
		L.Set(j, i, lji)
		Lch[j][i] <- lji
	}
}

func LUgo(A *mat.Dense, L *mat.Dense, U *mat.Dense, Lch *[3][3]chan float64, Uch *[3][3]chan float64) {
	r, _ := A.Dims()
	var J int
	for i := 0; i < r; i++ {
		for j := J; j < r; j++ {
			//fmt.Println(i, j)
			wg.Add(1)
			go Uset(A, L, U, Lch, Uch, i, j)
			wg.Add(1)
			go Lset(A, L, U, Lch, Uch, i, j)
		}
		J += 1
	}
	wg.Wait()
}

func LU(A *mat.Dense, L *mat.Dense, U *mat.Dense) {
	r, _ := A.Dims()
	var Aij, Aji, Uij, Lji float64
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			for k := 0; k < r; k++ {
				Aij += L.At(i, k) * U.At(k, j)
			}
			Uij = A.At(i, j) - Aij
			U.Set(i, j, Uij)
			Aij = 0

			if i != j {
				for k := 0; k < r; k++ {
					Aji += L.At(j, k) * U.At(k, i)
				}
				if U.At(i, i) == 0 {
					Lji = 0
				} else {
					Lji = (A.At(j, i) - Aji) / U.At(i, i)
				}
				L.Set(j, i, Lji)
				Aji = 0
			}
		}
	}
}

func main() {
	/*
		n := 3
		var x []float64
		for i := 0; i < n*n; i++ {
			r := rand.Intn(100)
			x = append(x, float64(r))
		}
		A := mat.NewDense(n, n, x)
	*/

	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	A := mat.NewDense(3, 3, x)
	//x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	//A := mat.NewDense(4, 4, x)
	//matPrint(A)

	r, c := A.Dims()
	L1 := mat.NewDense(r, c, nil)
	U1 := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		L1.Set(i, i, 1)
	}

	t1 := time.Now()
	var Lch [3][3]chan float64
	var Uch [3][3]chan float64
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			Lch[i][j] = make(chan float64, 1)
			Uch[i][j] = make(chan float64, 1)
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			if i == j {
				Lch[i][j] <- 1
			}
			if i < j {
				Lch[i][j] <- 0
				Uch[j][i] <- 0
			}
		}
	}
	LUgo(A, L1, U1, &Lch, &Uch)
	fmt.Println("LUgo:", time.Now().Sub(t1))

	matPrint(L1)
	matPrint(U1)

	L2 := mat.NewDense(r, c, nil)
	U2 := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		L2.Set(i, i, 1)
	}

	t2 := time.Now()
	LU(A, L2, U2)
	fmt.Println("LU:", time.Now().Sub(t2))
	//matPrint(L2)
	//matPrint(U2)
	fmt.Println("")

}
