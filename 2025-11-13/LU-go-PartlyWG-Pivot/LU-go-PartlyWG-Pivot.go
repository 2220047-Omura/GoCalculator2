package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

var wg sync.WaitGroup
var A [size][size]big.Float
var B [size]big.Float
var calc [size][size]big.Float
var P [size][size]big.Float

const size = 8

func initialize() {
	//var a, b big.Float
	//var a, n, i2, j2 big.Float
	var n big.Float
	one := big.NewFloat(1)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			/*
				r := rand.Float64()
				a.SetFloat64(r)
				r = rand.Float64()
				b.SetFloat64(r)
				A[i][j].SetPrec(1024).Mul(&a, &b)
			*/
			/*
				i2.SetInt64(int64(i))
				j2.SetInt64(int64(j))
				n.Add(&i2, &j2)
				n.Add(&n, one)
				a.SetPrec(1024).Quo(one, &n)
				A[i][j].SetPrec(1024).Set(&a)
			*/
			n.Add(&n, one)
			A[i][j].SetPrec(1024).Set(&n)
		}
	}

	B[0].SetPrec(1024).SetString("1")
	P[0][0].SetPrec(1024).SetString("1")
	for i := 1; i < size; i++ {
		B[i].SetPrec(1024).SetString("0")
		P[i][i].SetPrec(1024).SetString("1")
	}
}

func LUfact1(k int, i int) {
	A[i][k].SetPrec(1024).Quo(&A[i][k], &A[k][k])
}

func LUfact2(k int, i int, j int) {
	calc[i][j].SetPrec(1024).Mul(&A[i][k], &A[k][j])
	A[i][j].SetPrec(1024).Sub(&A[i][j], &calc[i][j])
}

func call1(k int, i int, N int) {

	LUfact1(k, i)
	for j := k + 1; j < N; j++ {
		call2(k, i, j)
	}
}

func call2(k int, i int, j int) {

	LUfact2(k, i, j)
}

func call1WG(k int, i int, N int) {
	defer wg.Done()
	var wg2 sync.WaitGroup

	LUfact1(k, i)
	for j := k + 1; j < N; j++ {
		wg2.Add(1)
		go call2WG(k, i, j, &wg2)
	}
	wg2.Wait()
}

func call2WG(k int, i int, j int, wg2 *sync.WaitGroup) {
	defer wg2.Done()

	LUfact2(k, i, j)
}

func call(k int, i int, N int, wg *sync.WaitGroup) {
	defer wg.Done()

	LUfact1(k, i)

	for j := k + 1; j < N; j++ {
		call2(k, i, j)
	}
}

func Pivot(k int) {
	var n_max big.Float
	n_max.Abs(&A[k][k])
	pivot_n := k

	for i := k + 1; i < size; i++ {
		if n_max.Cmp(&A[i][k]) == -1 {
			n_max.Set(&A[i][k])
			pivot_n = i
		}
	}

	P[k], P[pivot_n] = P[pivot_n], P[k]
	PrintM(&P)

	A[k], A[pivot_n] = A[pivot_n], A[k]
	PrintM(&A)
}

func comp() {
	var tmp, p big.Float
	tmp.SetPrec(1024)
	one := big.NewFloat(1)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if P[i][j].Cmp(one) == 0 {
				B[i].Mul(&B[i], &P[i][j])
				break
			}
		}
	}

	// forward substitution
	for i := 1; i < size; i++ {
		for j := 0; j <= i-1; j++ {
			tmp.Mul(&B[j], &A[i][j])
			B[i].Sub(&B[i], &tmp)
		}
	}

	// backward substitution
	for i := size - 1; i >= 0; i-- {
		for j := size - 1; j > i; j-- {
			tmp.Mul(&B[j], &A[i][j])
			B[i].Sub(&B[i], &tmp)
		}
		B[i].Quo(&B[i], &A[i][i])
	}

	for i := 0; i < size; i++ {
		p.SetPrec(100).Set(&B[i])
		fmt.Println(&p)
	}
}

func PrintM(M *[size][size]big.Float) {
	//行列をプリント

	for i := 0; i < size; i++ {
		print("\n")
		for j := 0; j < size; j++ {
			fmt.Print(&M[i][j], " ")
		}
	}
	print("\n")
}

func main() {
	var t time.Time

	//fmt.Println("-----逐次-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		if A[k][k].Cmp(big.NewFloat(0)) == 0 {
			fmt.Println("Pivot")
			Pivot(k)
		}
		for i := k + 1; i < size; i++ {
			call1(k, i, size)
		}
	}
	t2 := time.Now().Sub(t)
	comp()
	fmt.Println("逐次：", t2, "\n")
	/*
	   //fmt.Println("-----並列-----")
	   initialize()

	   t = time.Now()

	   	for k := 0; k < size; k++ {
	   		for i := k + 1; i < size; i++ {
	   			wg.Add(1)
	   			go call1WG(k, i, size)
	   		}
	   		wg.Wait()
	   	}

	   t2 = time.Now().Sub(t)
	   //comp()
	   fmt.Println("並列：", t2, "\n")

	   //fmt.Println("-----一部並列-----")
	   initialize()

	   t = time.Now()

	   	for k := 0; k < size; k++ {
	   		for i := k + 1; i < size; i++ {
	   			wg.Add(1)
	   			go call(k, i, size, &wg)
	   		}
	   		wg.Wait()
	   	}

	   t2 = time.Now().Sub(t)
	   //comp()
	   fmt.Println("一部並列：", t2, "\n")
	*/
}
