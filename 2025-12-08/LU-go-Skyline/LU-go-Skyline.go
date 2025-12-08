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

var Ask [size2]big.Float
var isk [size]int
var Lsk [size2]big.Float
var MULsk [size2]big.Float
var SUMsk [size2]big.Float

const size = 5
const size2 = 10

func initialize() {

	//setHilbert()
	//setRand()
	//setSimple()
	//setSkyline()
	setSkylineTest()
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

func setSkyline() {
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

func setSkylineTest() {
	var AskTest = [size2]int{2, 1, 3, 0, 4, 7, 8, 2, 3, 5}
	for i := 0; i < size2; i++ {
		Ask[i].SetInt64(int64(AskTest[i]))
	}
	fmt.Println(AskTest)
	var iskTest = [size]int{0, 1, 4, 5, 9}
	for i := 0; i < size; i++ {
		isk[i] = iskTest[i]
	}
	/*
		fmt.Println(iskTest)
		var c int
		var test = [15]int{2, 0, 3, 0, 0, 1, 0, 0, 8, 4, 0, 2, 7, 3, 5}
		for i := 0; i < size; i++ {
			for j := i; j < size; j++ {
				A[i][j].SetInt64(int64(test[c]))
				A[j][i].SetInt64(int64(test[c]))
				c += 1
			}
		}
	*/
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

func Uset(i int, j int) {
	for k := 0; k < i; k++ {
		MUL[i][j].Mul(&A[i][k], &A[k][j])
		SUM[i][j].Add(&SUM[i][j], &MUL[i][j])
	}
	A[i][j].Sub(&A[i][j], &SUM[i][j])
}

func Lset(j int, i int) {
	for k := 0; k < i; k++ {
		MUL[j][i].Mul(&A[j][k], &A[k][i])
		SUM[j][i].Add(&SUM[j][i], &MUL[j][i])
	}
	SUM[j][i].Sub(&A[j][i], &SUM[j][i])
	A[j][i].Quo(&SUM[j][i], &A[i][i])
}

func Usetsk(a int, i int, j int) {
	var s int
	if isk[j]-isk[j-1]-(j-i)-1 < isk[i]-isk[i-1]-1 {
		s = isk[j] - isk[j-1] - (j - i) - 1
	} else {
		s = isk[i] - isk[i-1] - 1
	}
	fmt.Println("s:", s)

	for k := 0; k < s; k++ {
		//fmt.Println("Lki*Ukj: ", isk[i]-(s-k), isk[j]-(j-i)-(s-k))
		//fmt.Println("i,Lki,Uii: ", i-(s-k), &Ask[isk[i]-(s-k)], &Ask[isk[i-(s-k)]])
		Lsk[a].Quo(&Ask[isk[i]-(s-k)], &Ask[isk[i-(s-k)]])
		MULsk[a].Mul(&Lsk[a], &Ask[isk[j]-(j-i)-(s-k)])
		SUMsk[a].Add(&SUMsk[a], &MULsk[a])
	}
	Ask[a].Sub(&Ask[a], &SUMsk[a])
}

func UsetWG(i int, j int) {
	defer wg.Done()
	for k := 0; k < i; k++ {
		MUL[i][j].Mul(&A[i][k], &A[k][j])
		SUM[i][j].Add(&SUM[i][j], &MUL[i][j])
	}
	A[i][j].Sub(&A[i][j], &SUM[i][j])
}

func LsetWG(j int, i int) {
	defer wg.Done()
	for k := 0; k < i; k++ {
		MUL[j][i].Mul(&A[j][k], &A[k][i])
		SUM[j][i].Add(&SUM[j][i], &MUL[j][i])
	}
	SUM[j][i].Sub(&A[j][i], &SUM[j][i])
	A[j][i].Quo(&SUM[j][i], &A[i][i])
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

func PrintArr(M *[size2]big.Float) {
	//行列をプリント

	for i := 0; i < size2; i++ {
		fmt.Print(&M[i], " ")
	}
	print("\n")
}

func main() {
	//fmt.Println("【スカイライン法】")
	var ts, te time.Time

	//fmt.Println("-----逐次-----")
	initialize()
	//fmt.Println(Ask)
	var i, j, c int
	ts = time.Now()
	for a := 1; a < size2; a++ {
		c = 1
		for b := 1; b < size2; b++ {
			i = c - (isk[c] - b)
			j = c
			if i == a {
				fmt.Println("(i, j)=", i, j)
				Usetsk(b, i, j)
			}
			if b == isk[c] {
				c += 1
			}
		}
	}
	te = time.Now()
	fmt.Println("逐次：", te.Sub(ts), "\n")
	//PrintM(&L)
	//fmt.Println(&Ask)
	PrintArr(&Ask)
	//comp()

	//fmt.Println("-----逐次-----")
	initialize()
	//PrintM(&A)
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
	PrintM(&A)
	//comp()
	/*
	   //fmt.Println("-----逐次-----")
	   initialize()
	   PrintM(&A)
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
	   PrintM(&A)
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
	   PrintM(&A)
	   //comp()
	*/
}
