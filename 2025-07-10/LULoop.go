package main

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func AijFunc(L mat.Matrix, U mat.Matrix, i int, j int, k int, c chan<- float64) {
	c <- L.At(i, k) * U.At(k, j)
}

func AjiFunc(L mat.Matrix, U mat.Matrix, i int, j int, k int, c chan<- float64) {
	c <- L.At(j, k) * U.At(k, i)
}

// func Uset(A mat.Matrix, L *mat.Dense, U *mat.Dense, i int, j int, c1 chan float64) {
func Uset(A mat.Matrix, L *mat.Dense, U *mat.Dense, i int, j int) {
	r, _ := A.Dims()
	var Aij, AijSum, Uij float64
	for k := 0; k < r; k++ {
		c3 := make(chan float64)
		go AijFunc(L, U, i, j, k, c3)
		Aij, _ = <-c3
		AijSum += Aij
		close(c3)
	}
	Uij = A.At(i, j) - AijSum
	//c1 <- Uij
	U.Set(i, j, Uij)
}

// func Lset(A mat.Matrix, L *mat.Dense, U *mat.Dense, i int, j int,c2 chan float64) {
func Lset(A mat.Matrix, L *mat.Dense, U *mat.Dense, i int, j int) {
	r, _ := A.Dims()
	var Aji, AjiSum, Lji float64
	if i != j {
		for k := 0; k < r; k++ {
			c3 := make(chan float64)
			go AjiFunc(L, U, i, j, k, c3)
			Aji, _ = <-c3
			AjiSum += Aji
		}
		if U.At(i, i) == 0 {
			Lji = 0
		} else {
			Lji = (A.At(j, i) - AjiSum) / U.At(i, i)
		}
		//c2 <- Lji
		L.Set(j, i, Lji)
	}
}

func LUgo(A mat.Matrix, L *mat.Dense, U *mat.Dense) {
	r, _ := A.Dims()
	//J := 0
	//var Aji, Uij, Lji float64
	for i := 0; i < r; i++ {
		//for j := J; j < r; j++ {
		for j := 0; j < r; j++ {
			Uset(A, L, U, i, j)
			Lset(A, L, U, i, j)

			/*
				c1 := make(chan float64, j)
				go Uset(A, L, U, i, j, c1)
				Uij := <-c1
				U.Set(i, j, Uij)
				close(c1)

				c2 := make(chan float64, j)
				go Lset(A, L, U, i, j, c2)
				Lji := <-c2
				L.Set(j, i, Lji)
				close(c2)
			*/
		}
		/*
			for j := J; j < c; j++ {
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
		*/
		//J += 1
	}
}

func LU(A mat.Matrix, L *mat.Dense, U *mat.Dense) {
	r, _ := A.Dims()
	//J := 0
	var Aij, Aji, Uij, Lji float64
	for i := 0; i < r; i++ {
		//for j := J; j < r; j++ {
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
		/*
			for j := J; j < c; j++ {
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
		*/
		//J += 1
	}
}

func main() {
	for {
		/*
			scanner := bufio.NewScanner(os.Stdin)
			print("正方行列Aのサイズを入力してください：")
			scanner.Scan()
			if scanner.Text() == "" {
				break
			}
			n, _ := strconv.Atoi(scanner.Text())
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
		//fmt.Println(r,c)

		L1 := mat.NewDense(r, c, nil)
		U1 := mat.NewDense(r, c, nil)
		for i := 0; i < r; i++ {
			L1.Set(i, i, 1)
		}
		t1 := time.Now()
		LUgo(A, L1, U1)
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

		break
	}
}
