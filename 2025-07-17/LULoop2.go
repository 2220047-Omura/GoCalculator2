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

type LchSt struct {
	i, j int
	c    chan float64
	flag bool
}

type UchSt struct {
	i, j int
	c    chan float64
	flag bool
}

func LchBool(i int, j int) struct {
	c    chan float64
	flag bool
}

func UchBool(i int, j int) struct {
	c    chan float64
	flag bool
}

func getLij(L *mat.Dense, i int, j int) {
	for {
		if LchBool(i, j).flag == true {
			LchBool(i, j).c <- L.At(i, j)
			break
		}
	}
}

func getUij(U *mat.Dense, i int, j int) {
	for {
		if UchBool(i, j).flag == true {
			UchBool(i, j).c <- U.At(i, j)
			break
		}
	}
}

func Uset(A mat.Matrix, L *mat.Dense, U *mat.Dense, i int, j int) {
	r, _ := A.Dims()
	var aij, uij, lij float64
	for k := 0; k < r; k++ {
		getLij(L, i, k)
		if k != i {
			getUij(U, k, j)
		}
		lij = <-LchBool(i, k).c
		uij = <-UchBool(k, j).c
		aij += lij * uij
	}
	uij = A.At(i, j) - aij
	//c1 <- Uij
	a := UchBool(i, j)
	U.Set(i, j, uij)
	a.flag = true
	aij = 0
}

func Lset(A mat.Matrix, L *mat.Dense, U *mat.Dense, i int, j int) {
	if i != j {
		r, _ := A.Dims()
		var aij, uij, lij float64
		for k := 0; k < r; k++ {
			getUij(U, j, k)
			if k != i {
				getLij(L, k, i)
			}
			lij = <-LchBool(k, i).c
			uij = <-UchBool(j, k).c
			aij += lij * uij
		}
		if U.At(i, i) == 0 {
			uij = 0
		} else {
			uij = (A.At(j, i) - aij) / U.At(i, i)
		}
		//c1 <- Uij
		L.Set(j, i, uij)
		aij = 0
	}
}

func LUgo(A mat.Matrix, L *mat.Dense, U *mat.Dense) {
	r, _ := A.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			go Uset(A, L, U, i, j)
			go Lset(A, L, U, i, j)
		}
	}
}

func LU(A mat.Matrix, L *mat.Dense, U *mat.Dense) {
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

	var N int
	L1 := mat.NewDense(r, c, nil)
	U1 := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		L1.Set(i, i, 1)
	}
	t1 := time.Now()
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			Lch := LchBool(i, j)
			Uch := UchBool(j, i)
			if j >= N {
				Lch.flag = true
			}
			if j >= N+1 {
				Uch.flag = true
			}
		}
		N += 1
	}
	fmt.Println(LchBool(1, 1).flag)
	fmt.Println(UchBool(0, 1).flag)
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

}
