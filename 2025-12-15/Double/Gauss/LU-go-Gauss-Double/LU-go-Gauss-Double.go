package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup
var A [size][size]float64
var B [size]float64

var calc [size][size]float64

const size = 300

func initialize() {
    var pcg = rand.NewPCG(0, 0)
	var rng = rand.New(pcg)

	//setHilbert()
	setRand(rng)
	//setSkyline(rng)
	mulDiagonal()
	//setSperse()

	for i := 0; i < size; i ++{
		B[i] = 0
		for j := 0; j < size; j ++{
			calc[i][j]=0
		}
	}
	B[0] = 1
}

func setSperse() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i-j > 1 || j-i > 1 {
				A[i][j] = 0
			}
		}
	}
}

func setHilbert() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			A[i][j] = float64(1 / (i + j + 1))
		}
	}
}

func setRand(rng *rand.Rand) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			A[i][j] = rng.Float64()
		}
	}
}

func setSkyline(rng *rand.Rand){
	var c int
    for i := 0; i < size; i++ {
        A[i][i] = rng.Float64()
        if (i-c<0){
            c = 0;
        }
        for j := i-1; j >= c; j-- {
	        A[i][j] = rng.Float64()
			A[j][i] = rng.Float64()
        }
        c += 2;
    }
}

func mulDiagonal() {
	for i := 0; i < size; i++ {
		A[i][i] *= 100
		for j := i + 1; j < size; j++ {
			A[i][j] *= float64(100/(j-i+1))
		}
	}
}

func LUfact1(k int, i int) {
	A[i][k] = A[i][k]/A[k][k]
}

func LUfact2(k int, i int, j int) {
	calc[i][j] = A[i][k] * A[k][j]
	A[i][j] -= calc[i][j]
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
	var tmp float64

	// forward substitution
	for i := 1; i < size; i++ {
		for j := 0; j <= i-1; j++ {
			tmp = B[j]*A[i][j]
			B[i] -= tmp
		}
	}

	// backward substitution
	for i := size - 1; i >= 0; i-- {
		for j := size - 1; j > i; j-- {
			tmp = B[j]*A[i][j]
			B[i] -= tmp
		}
		B[i] = B[i]/A[i][i]
	}

	for i := 0; i < size; i++ {
		fmt.Println(B[i])
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
