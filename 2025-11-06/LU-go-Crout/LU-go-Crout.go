package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

var wg sync.WaitGroup
var A [size][size]big.Float
var L [size][size]big.Float
var U [size][size]big.Float
var B [size]big.Float

const size = 8

func initialize() {
	//var a, b big.Float
	var a, n, i2, j2 big.Float
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

			i2.SetInt64(int64(i))
			j2.SetInt64(int64(j))
			n.Add(&i2, &j2)
			// n.Add(&n, big.NewFloat(1))
			n.Add(&n, one)
			// a.SetPrec(1024).Quo(big.NewFloat(1), &n)
			a.SetPrec(1024).Quo(one, &n)
			A[i][j].SetPrec(1024).Set(&a)

			L[i][j].SetPrec(1024)
			U[i][j].SetPrec(1024)
			if i == j {
				L[i][j].Set(one)
			}
		}
	}

	B[0].SetPrec(1024).SetString("1")
	for i := 1; i < size; i++ {
		B[i].SetPrec(1024).SetString("0")
	}
}

func Uset(i int, j int) {
	var MUL, SUM big.Float
	for k := 0; k < size; k++ {
		if k != i {
			MUL.Mul(&L[i][k], &U[k][j])
			SUM.Add(&SUM, &MUL)
		}
	}
	U[i][j].Sub(&A[i][j], &SUM)
}

func Lset(j int, i int) {
	var MUL, SUM big.Float
	for k := 0; k < size; k++ {
		if k != i {
			MUL.Mul(&L[j][k], &U[k][i])
			SUM.Add(&SUM, &MUL)
		}
	}
	SUM.Sub(&A[j][i], &SUM)
	L[j][i].Quo(&SUM, &U[i][i])
}

func UsetWG(i int, j int) {
	defer wg.Done()
	var MUL, SUM big.Float
	for k := 0; k < size; k++ {
		if k != i {
			MUL.Mul(&L[i][k], &U[k][j])
			SUM.Add(&SUM, &MUL)
		}
	}
	U[i][j].Sub(&A[i][j], &SUM)
}

func LsetWG(j int, i int) {
	defer wg.Done()
	var MUL, SUM big.Float
	for k := 0; k < size; k++ {
		if k != i {
			MUL.Mul(&L[j][k], &U[k][i])
			SUM.Add(&SUM, &MUL)
		}
	}
	SUM.Sub(&A[j][i], &SUM)
	L[j][i].Quo(&SUM, &U[i][i])
}

func comp() {
	for i := 0; i < size; i++ {
		A[i][i].Set(&U[i][i])
		for j := i + 1; j < size; j++ {
			A[i][j].Set(&U[i][j])
			A[j][i].Set(&L[j][i])
		}
	}
	//PrintM(&A)
	var tmp, p big.Float
	tmp.SetPrec(1024)

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

func SimpleA(A *[size][size]big.Float) {
	//各要素が左上から1, 2, 3, ... と決められる行列を生成

	var n big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			n.Add(&n, big.NewFloat(1))
			A[i][j].SetPrec(1024).Set(&n)
		}
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
	//fmt.Println("【クラウト法】")
	var ts, te time.Time

	//fmt.Println("-----逐次-----")
	initialize()
	//SimpleA(&A)
	ts = time.Now()
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			Uset(i, j)
		}
		for j := i + 1; j < size; j++ {
			Lset(j, i)
		}
	}
	te = time.Now()
	fmt.Println("逐次：", te.Sub(ts), "\n")
	//PrintM(&L)
	//PrintM(&U)
	comp()

	//fmt.Println("-----並列-----")
	initialize()
	//SimpleA(&A)
	ts = time.Now()
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			wg.Add(1)
			go UsetWG(i, j)
		}
		wg.Wait()
		for j := i + 1; j < size; j++ {
			wg.Add(1)
			go LsetWG(j, i)
		}
		wg.Wait()
	}
	te = time.Now()
	fmt.Println("並列：", te.Sub(ts), "\n")
	//PrintM(&L)
	//PrintM(&U)
	comp()
}
