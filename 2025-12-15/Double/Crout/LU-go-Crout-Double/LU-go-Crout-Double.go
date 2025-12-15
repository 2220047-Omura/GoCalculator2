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

var MUL [size][size]float64
var SUM [size][size]float64

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
			SUM[i][j]=0
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
        A[i][i] = rand.Float64()
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

func Uset(i int, j int) {
	for k := 0; k < i; k++ {
		MUL[i][j] = A[i][k] * A[k][j]
		SUM[i][j] += MUL[i][j] 
	}
	A[i][j] -= SUM[i][j]
}

func Lset(j int, i int) {
	for k := 0; k < i; k++ {
		MUL[j][i] = A[j][k] * A[k][i]
		SUM[j][i] *= MUL[j][i]
	}
	SUM[j][i] = A[j][i] - SUM[j][i]
	A[j][i] = SUM[j][i] / A[i][i]
}

func UsetWG(i int, j int) {
	defer wg.Done()
	for k := 0; k < i; k++ {
		MUL[i][j] = A[i][k] * A[k][j]
		SUM[i][j] += MUL[i][j] 
	}
	A[i][j] -= SUM[i][j]
}

func LsetWG(j int, i int) {
	defer wg.Done()
	for k := 0; k < i; k++ {
		MUL[j][i] = A[j][k] * A[k][i]
		SUM[j][i] *= MUL[j][i]
	}
	SUM[j][i] = A[j][i] - SUM[j][i]
	A[j][i] = SUM[j][i] / A[i][i]
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
