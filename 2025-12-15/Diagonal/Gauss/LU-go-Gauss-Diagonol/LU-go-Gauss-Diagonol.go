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
var calc [size][size]big.Float

const size = 300

func initialize() {

	//setHilbert()
	setRand()
	//setSkyline()
	//setSimple()
	mulDiagonal()
	//setSperse()

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

func setSkyline(){
	var c int
    var a big.Float
    for i := 0; i < size; i++ {
        r := rand.Float64()
		a.SetFloat64(r)
		A[i][i].SetPrec(1024).Set(&a)
        if (i-c<0){
            c = 0;
        }
        for j := i-1; j >= c; j-- {
	        r := rand.Float64()
		    a.SetFloat64(r)
		    A[i][j].SetPrec(1024).Set(&a)
			A[j][i].SetPrec(1024).Set(&a)
        }
        c += 2;
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

func call3(k int, i int, N int, wg *sync.WaitGroup) {
	defer wg.Done()

	LUfact1(k, i)

	for j := k + 1; j < N; j++ {
		call2(k, i, j)
	}
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
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i++ {
			//fmt.Println(k, i)
			call1(k, i, size)
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

	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i++ {
			//fmt.Println(k, i)
			wg.Add(1)
			go call3(k, i, size, &wg)
		}
		wg.Wait()
	}

	te = time.Now()
	fmt.Println("並列：", te.Sub(ts), "\n")
	//PrintM(&L)
	//PrintM(&A)
	//comp()
}
