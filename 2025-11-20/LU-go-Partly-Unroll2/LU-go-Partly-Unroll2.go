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

		}
	}

	B[0].SetPrec(1024).SetString("1")
	for i := 1; i < size; i++ {
		B[i].SetPrec(1024).SetString("0")
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

func call3(k int, i int, N int, wg *sync.WaitGroup) {
	defer wg.Done()

	LUfact1(k, i)

	for j := k + 1; j < N; j++ {
		call2(k, i, j)
	}
}

func call4(k int, i int, N int, wg *sync.WaitGroup) {
	defer wg.Done()

	LUfact1(k, i)

	for j := k + 1; j < N; j++ {
		LUfact2(k, i, j)
	}
}

func callUnroll(k int, i int, N int, M int) {
	for i2 := i; i2 < i+M; i2++ {
		//fmt.Println("(k, i2) = ", k, i2)
		LUfact1(k, i2)

		for j := k + 1; j < N; j++ {
			LUfact2(k, i2, j)
		}
	}
}

func callUnrollWG(k int, i int, N int, wg *sync.WaitGroup, M int) {
	defer wg.Done()
	for i2 := i; i2 < i+M; i2++ {
		//fmt.Println("(k, i2) = ", k, i2)
		LUfact1(k, i2)

		for j := k + 1; j < N; j++ {
			LUfact2(k, i2, j)
		}
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

func main() {
	var t time.Time
	//var A [size][size]big.Float
	//var B [size]big.Float

	//fmt.Println("-----逐次-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i++ {
			//fmt.Println(k, i)
			call1(k, i, size)
		}
	}
	t2 := time.Now().Sub(t)
	//comp()
	fmt.Println("逐次：", t2, "\n")

	//fmt.Println("-----逐次(直接呼び出し)-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i++ {
			//fmt.Println(k, i)
			LUfact1(k, i)

			for j := k + 1; j < size; j++ {
				LUfact2(k, i, j)
			}
		}
	}
	t2 = time.Now().Sub(t)
	//comp()
	fmt.Println("逐次(直接呼び出し)：", t2, "\n")

	//fmt.Println("-----アンローリング(逐次)-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i += 2 {
			//fmt.Println(k, i)
			call1(k, i, size)
			if i+1 != size {
				call1(k, i+1, size)
			}
		}
	}
	t2 = time.Now().Sub(t)
	//C.comp()
	fmt.Println("アンローリング(逐次)：", t2, "\n")

	//fmt.Println("-----アンローリング改(逐次)-----")
	initialize()

	p := 8
	t = time.Now()
	for k := 0; k < size; k++ {
		M := (size - k - 1) / p
		//fmt.Println("M:", M)
		for i := k + 1; i < k+1+M*p; i += M {
			//fmt.Println("(k, i) =", k, i)
			callUnroll(k, i, size, M)
		}
		for i := k + 1 + M*p; i < size; i++ {
			//fmt.Println("あまり")
			//fmt.Println("(k, i) =", k, i)
			call1(k, i, size)
		}
	}
	t2 = time.Now().Sub(t)
	//comp()
	fmt.Println("アンローリング改(逐次)：", t2, "\n")

	//fmt.Println("-----並列-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i++ {
			//fmt.Println(k, i)
			wg.Add(1)
			go call1WG(k, i, size)
		}
		wg.Wait()
	}
	t2 = time.Now().Sub(t)
	//comp()
	fmt.Println("並列：", t2, "\n")

	//fmt.Println("-----アンローリング(並列)-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i += 2 {
			//fmt.Println(k, i)
			wg.Add(1)
			go call1WG(k, i, size)
			if i+1 != size {
				wg.Add(1)
				go call1WG(k, i+1, size)
			}
		}
		wg.Wait()
	}
	t2 = time.Now().Sub(t)
	//C.comp()
	fmt.Println("アンローリング(並列)：", t2, "\n")

	//fmt.Println("-----一部並列-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i++ {
			//fmt.Println(k, i)
			wg.Add(1)
			go call3(k, i, size, &wg)
		}
		wg.Wait()
	}
	t2 = time.Now().Sub(t)
	//comp()
	fmt.Println("一部並列：", t2, "\n")

	//fmt.Println("-----一部並列(直接呼び出し)-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i++ {
			//fmt.Println(k, i)
			wg.Add(1)
			go call4(k, i, size, &wg)
		}
		wg.Wait()
	}
	t2 = time.Now().Sub(t)
	//comp()
	fmt.Println("一部並列(直接呼び出し)：", t2, "\n")

	//fmt.Println("-----アンローリング(一部並列)-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i += 2 {
			//fmt.Println(k, i)
			wg.Add(1)
			go call3(k, i, size, &wg)
			if i+1 != size {
				wg.Add(1)
				go call3(k, i+1, size, &wg)
			}
		}
		wg.Wait()
	}
	t2 = time.Now().Sub(t)
	//C.comp()
	fmt.Println("アンローリング(一部並列)：", t2, "\n")

	//fmt.Println("-----アンローリング(4)-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i += 4 {
			//fmt.Println(k, i)
			wg.Add(1)
			go call3(k, i, size, &wg)
			if i+1 < size {
				wg.Add(1)
				go call3(k, i+1, size, &wg)
				if i+2 < size {
					wg.Add(1)
					go call3(k, i+2, size, &wg)
					if i+3 < size {
						wg.Add(1)
						go call3(k, i+3, size, &wg)
					}
				}
			}
		}
		wg.Wait()
	}
	t2 = time.Now().Sub(t)
	//C.comp()
	fmt.Println("アンローリング(4)：", t2, "\n")

	//fmt.Println("-----アンローリング(一部並列)-----")
	initialize()

	p = 4
	t = time.Now()
	for k := 0; k < size; k++ {
		M := (size - k - 1) / p
		//fmt.Println("M:", M)
		for i := k + 1; i < k+1+M*p; i += M {
			//fmt.Println("(k, i) =", k, i)
			wg.Add(1)
			go callUnrollWG(k, i, size, &wg, M)
		}
		for i := k + 1 + M*p; i < size; i++ {
			//fmt.Println("あまり")
			//fmt.Println("(k, i) =", k, i)
			wg.Add(1)
			go call3(k, i, size, &wg)
		}
		wg.Wait()
	}
	t2 = time.Now().Sub(t)
	comp()
	fmt.Println("アンローリング改(一部並列)：", t2, "\n")
}
