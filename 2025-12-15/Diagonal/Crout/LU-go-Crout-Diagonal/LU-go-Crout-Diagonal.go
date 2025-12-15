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
var B [size]big.Float

var MUL [size][size]big.Float
var SUM [size][size]big.Float

const size = 300

func initialize() {
	var pcg = rand.NewPCG(0, 0)
	var rng = rand.New(pcg)

	//setHilbert()
	setRand(rng)
	//setSkyline(rng)
	//setSimple()
	mulDiagonal()
	//setSperse()

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			MUL[i][j].SetPrec(1024).SetString("0")
			SUM[i][j].SetPrec(1024).SetString("0")
		}
	}

	B[0].SetPrec(1024).SetString("1")
	for i := 1; i < size; i++ {
		B[i].SetPrec(1024).SetString("0")
	}
}

func setSperse() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i-j > 1 || j-i > 1 {
				A[i][j].SetString("0")
			}
		}
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

func setRand(rng *rand.Rand) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			r := rng.Float64()
			A[i][j].SetPrec(1024).SetFloat64(r)
		}
	}
}

func setSkyline(rng *rand.Rand) {
	var c int
	for i := 0; i < size; i++ {
		r := rng.Float64()
		A[i][i].SetPrec(1024).SetFloat64(r)
		if i-c < 0 {
			c = 0
		}
		for j := i - 1; j >= c; j-- {
			r := rng.Float64()
			A[i][j].SetPrec(1024).SetFloat64(r)
			A[j][i].SetPrec(1024).SetFloat64(r)
		}
		c += 2
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

func mulDiagonal() {
	var hdr, one, toI, toJ, div big.Float
	hdr.SetString("100")
	one.SetString("1")
	for i := 0; i < size; i++ {
		A[i][i].Mul(&A[i][i], &hdr)
		toI.SetInt64(int64(i))
		for j := i + 1; j < size; j++ {
			toJ.SetInt64(int64(j))
			div.Sub(&toJ, &toI)
			div.Add(&div, &one)
			div.Quo(&hdr, &div)
			A[i][j].Mul(&A[i][j], &div)
			A[j][i].Mul(&A[j][i], &div)
		}
	}
}

func Uset(i int, j int) {
	//var MUL, SUM big.Float
	for k := 0; k < i; k++ {
		MUL[i][j].Mul(&A[i][k], &A[k][j])
		SUM[i][j].Add(&SUM[i][j], &MUL[i][j])
		//MUL.Mul(&L[i][k], &U[k][j])
		//SUM.Add(&SUM, &MUL)
	}
	A[i][j].Sub(&A[i][j], &SUM[i][j])
	//U[i][j].Sub(&A[i][j], &SUM)
}

func Lset(j int, i int) {
	//var MUL, SUM big.Float
	for k := 0; k < i; k++ {
		MUL[j][i].Mul(&A[j][k], &A[k][i])
		SUM[j][i].Add(&SUM[j][i], &MUL[j][i])
		//MUL.Mul(&L[j][k], &U[k][i])
		//SUM.Add(&SUM, &MUL)
	}
	SUM[j][i].Sub(&A[j][i], &SUM[j][i])
	A[j][i].Quo(&SUM[j][i], &A[i][i])
	//SUM.Sub(&A[j][i], &SUM)
	//L[j][i].Quo(&SUM, &U[i][i])
}

func UsetWG(i int, j int) {
	defer wg.Done()
	//var MUL, SUM big.Float
	for k := 0; k < i; k++ {
		MUL[i][j].Mul(&A[i][k], &A[k][j])
		SUM[i][j].Add(&SUM[i][j], &MUL[i][j])
		//MUL.Mul(&L[i][k], &U[k][j])
		//SUM.Add(&SUM, &MUL)
	}
	A[i][j].Sub(&A[i][j], &SUM[i][j])
	//U[i][j].Sub(&A[i][j], &SUM)
}

func LsetWG(j int, i int) {
	defer wg.Done()
	//var MUL, SUM big.Float
	for k := 0; k < i; k++ {
		MUL[j][i].Mul(&A[j][k], &A[k][i])
		SUM[j][i].Add(&SUM[j][i], &MUL[j][i])
		//MUL.Mul(&L[j][k], &U[k][i])
		//SUM.Add(&SUM, &MUL)
	}
	SUM[j][i].Sub(&A[j][i], &SUM[j][i])
	A[j][i].Quo(&SUM[j][i], &A[i][i])
	//SUM.Sub(&A[j][i], &SUM)
	//L[j][i].Quo(&SUM, &U[i][i])
}

func comp() {
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
	ts = time.Now()
	for j := 1; j < size; j++ {
		Lset(j, 0)
	}
	for i := 1; i < size; i++ {
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
	//comp()

	//fmt.Println("-----並列-----")
	initialize()
	ts = time.Now()

	for j := 1; j < size; j++ {
		wg.Add(1)
		go LsetWG(j, 0)
	}
	wg.Wait()
	for i := 1; i < size; i++ {
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
	//PrintM(&A)
	//comp()
}
