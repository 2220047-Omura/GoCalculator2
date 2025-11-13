package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup
var A [size][size]big.Float
var L [size][size]big.Float
var U [size][size]big.Float
var B [size]big.Float
var P [size][size]big.Float
var z [size]big.Float
var x [size]big.Float
var r big.Float

var MUL [size][size]big.Float
var SUM [size][size]big.Float

const size = 8

func initialize() {

	//setHilbert()
	//setRand()
	setSimple()

	one := big.NewFloat(1)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			L[i][j].SetPrec(1024)
			U[i][j].SetPrec(1024)
			MUL[i][j].SetPrec(1024)
			SUM[i][j].SetPrec(1024)
			if i == j {
				L[i][j].Set(one)
			}
		}
	}

	B[0].SetPrec(1024).SetString("1")
	P[0][0].SetPrec(1024).SetString("1")
	for i := 1; i < size; i++ {
		B[i].SetPrec(1024).SetString("0")
		P[i][i].SetPrec(1024).SetString("1")
	}
}

func setHilbert() {
	var a, n, i2, j2 big.Float
	one := big.NewFloat(1)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			i2.SetInt64(int64(i))
			j2.SetInt64(int64(j))
			n.Add(&i2, &j2)
			n.Add(&n, one)
			a.SetPrec(1024).Quo(one, &n)
			A[i][j].SetPrec(1024).Set(&a)
		}
	}
}

func setRand() {
	var a, b big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			r := rand.Float64()
			a.SetFloat64(r)
			r = rand.Float64()
			b.SetFloat64(r)
			A[i][j].SetPrec(1024).Mul(&a, &b)
		}
	}
}

func setSimple() {
	var n big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			n.Add(&n, big.NewFloat(1))
			A[i][j].SetPrec(1024).Set(&n)
		}
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

	L[k], L[pivot_n] = L[pivot_n], L[k]
	PrintM(&L)
}

func LU() {
	for i := 0; i < size; i++ {
		Pivot(i)
		for j := i; j < size; j++ {
			U[i][j].Set(&A[i][j])
			L[j][i].Quo(&A[j][i], &U[i][i])
		}
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				MUL[j][k].Mul(&L[j][i], &U[i][k])
				A[j][k].Sub(&A[j][k], &MUL[j][k])
			}
		}
	}
}

func solveZ() {
	for i := 0; i < size; i++ {
		r.SetString("0")
		for j := 0; j < i; j++ {
			MUL[i][j].Mul(&L[i][j], &z[j])
			r.Add(&r, &MUL[i][j])
		}
		z[i].Sub(&B[i], &r)
	}
}

func solveX() {
	for i := size - 1; i >= 0; i-- {
		r.SetString("0")
		for j := i + 1; j < size; j++ {
			MUL[i][j].Mul(&U[i][j], &x[j])
			r.Add(&r, &MUL[i][j])
		}
		x[i].Sub(&z[i], &r)
		x[i].Quo(&x[i], &U[i][i])
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
	var ts, te time.Time

	//fmt.Println("-----逐次-----")
	initialize()
	ts = time.Now()
	LU()
	te = time.Now()
	fmt.Println("逐次：", te.Sub(ts), "\n")
	//PrintM(&L)
	//PrintM(&U)
	//comp()
	solveZ()
	solveX()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(&x[i], " \n")
		}
	}
	print("\n")
}
