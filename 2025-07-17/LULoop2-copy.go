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

type Lst struct {
	c    chan float64
	flag bool
}

type Ust struct {
	c    chan float64
	flag bool
}

func getLij(L *mat.Dense, Larr *[3][3]Lst, i int, j int) {
	for {
		if Larr[i][j].flag == true {
			Larr[i][j].c <- L.At(i, j)
			break
		}
	}
}

func getUij(U *mat.Dense, Uarr *[3][3]Ust, i int, j int) {
	for {
		if Uarr[i][j].flag == true {
			Uarr[i][j].c <- U.At(i, j)
			break
		}
	}
}

func Uset(A mat.Matrix, L *mat.Dense, U *mat.Dense, Larr *[3][3]Lst, Uarr *[3][3]Ust, i int, j int) {
	r, _ := A.Dims()
	var aij, uij, lij float64
	for k := 0; k < r; k++ {
		getLij(L, Larr, i, k)
		lij = <-Larr[i][k].c
		if k != i {
			getUij(U, Uarr, k, j)
			uij = <-Uarr[k][j].c
		} else {
			uij = 0
		}
		aij += lij * uij
	}
	uij = A.At(i, j) - aij
	//c1 <- Uij
	U.Set(i, j, uij)
	Uarr[i][j].flag = true
	aij = 0
}

func Lset(A mat.Matrix, L *mat.Dense, U *mat.Dense, Larr *[3][3]Lst, Uarr *[3][3]Ust, i int, j int) {
	if i != j {
		r, _ := A.Dims()
		var aij, uij, lij float64
		for k := 0; k < r; k++ {
			getUij(U, Uarr, j, k)
			uij = <-Uarr[j][k].c
			if k != i {
				getLij(L, Larr, k, i)
				lij = <-Larr[k][i].c
			} else {
				lij = 0
			}
			aij += lij * uij
		}
		if U.At(i, i) == 0 {
			lij = 0
		} else {
			lij = (A.At(j, i) - aij) / U.At(i, i)
		}
		//c1 <- Uij
		L.Set(j, i, lij)
		Larr[j][i].flag = true
		aij = 0
	}
}

func LUgo(A mat.Matrix, L *mat.Dense, U *mat.Dense, Larr *[3][3]Lst, Uarr *[3][3]Ust) {
	r, _ := A.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			go Uset(A, L, U, Larr, Uarr, i, j)
			go Lset(A, L, U, Larr, Uarr, i, j)
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

	L1 := mat.NewDense(r, c, nil)
	U1 := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		L1.Set(i, i, 1)
	}

	t1 := time.Now()
	var Larr [3][3]Lst
	var Uarr [3][3]Ust
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			Larr[i][j].c = make(chan float64, 1)
			Uarr[i][j].c = make(chan float64, 1)
			if i <= j {
				Larr[i][j].flag = true
			}
			if i < j {
				Uarr[j][i].flag = true
			}
		}
	}
	fmt.Println(Larr)
	fmt.Println(Uarr)
	LUgo(A, L1, U1, &Larr, &Uarr)
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
