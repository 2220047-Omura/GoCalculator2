package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var Ask []float64
var isk [size]int
var Lsk []float64
var MULsk []float64
var SUMsk []float64

const size = 500

var E int

//var N = 10

func initialize() {
	var n int
	for i := 1; i < size; i++ {
		n -= 1
		if n < 0 {
			n = i
		}
		E = E + n + 1
	}
	var zero float64 = 0
	for i := 0; i < E; i++ {
		Ask = append(Ask, zero)
		Lsk = append(Lsk, zero)
		MULsk = append(MULsk, zero)
		SUMsk = append(SUMsk, zero)
	}
}

func reset() {
	var pcg = rand.NewPCG(0, 0)
	var rng = rand.New(pcg)
	//fmt.Println(N)
	setSkyline(rng)
	mulDiagonal()
	//setSkylineTest()
	for i := 0; i < E; i++ {
		Lsk[i] = 0
		MULsk[i] = 0
		SUMsk[i] = 0
	}
}

func setSkyline(rng *rand.Rand) {
	var n int
	isk[0] = 0
	for i := 1; i < size; i++ {
		n -= 1
		if n < 0 {
			n = i
		}
		isk[i] = isk[i-1] + n + 1
	}

	for i := 0; i < E; i++ {
		Ask[i] = rng.Float64()
	}

}

func mulDiagonal() {
	c := 0
	for i := 0; i < E; i++ {
		Ask[i] = Ask[i] * 100 / float64(isk[c]-i+1)
		if i == isk[c] {
			c += 1
		}
	}
}

func setSkylineTest() {
	var AskTest = []int{2, 1, 3, 0, 4, 7, 8, 2, 3, 5}
	//N = 10
	for i := 0; i < E; i++ {
		Ask[i] = float64(AskTest[i])
	}
	fmt.Println(AskTest)
	var iskTest = [size]int{0, 1, 4, 5, 9}
	for i := 0; i < size; i++ {
		isk[i] = iskTest[i]
	}
	fmt.Println(iskTest)
}

func Usetsk(b int, i int, j int) {
	var s int
	if isk[j]-isk[j-1]-(j-i)-1 < isk[i]-isk[i-1]-1 {
		s = isk[j] - isk[j-1] - (j - i) - 1
	} else {
		s = isk[i] - isk[i-1] - 1
	}
	//fmt.Println("s:", s)

	for k := 0; k < s; k++ {
		Lsk[b] = Ask[isk[i]-(s-k)] / Ask[isk[i-(s-k)]]
		MULsk[b] = Lsk[b] * Ask[isk[j]-(j-i)-(s-k)]
		SUMsk[b] += MULsk[b]
	}
	Ask[b] -= SUMsk[b]
}

func UsetskWG(b int, i int, j int, wg *sync.WaitGroup) {
	defer wg.Done()
	var s int
	if isk[j]-isk[j-1]-(j-i)-1 < isk[i]-isk[i-1]-1 {
		s = isk[j] - isk[j-1] - (j - i) - 1
	} else {
		s = isk[i] - isk[i-1] - 1
	}
	//fmt.Println("s:", s)

	for k := 0; k < s; k++ {
		Lsk[b] = Ask[isk[i]-(s-k)] / Ask[isk[i-(s-k)]]
		MULsk[b] = Lsk[b] * Ask[isk[j]-(j-i)-(s-k)]
		SUMsk[b] += MULsk[b]
	}
	Ask[b] -= SUMsk[b]
}

func PrintArr(M []float64) {
	//行列をプリント
	//for i := 0; i < E; i++ {
	for i := E - 3; i < E; i++ {
		fmt.Print(M[i], " ")
	}
	print("\n")
}

func main() {
	//fmt.Println("【スカイライン法】")
	var ts, te time.Time
	var i, j, c int
	var wg sync.WaitGroup
	initialize()

	//fmt.Println("-----逐次-----")
	reset()
	ts = time.Now()
	for a := 1; a < E; a++ {
		c = 1
		for b := 1; b < E; b++ {
			i = c - (isk[c] - b)
			j = c
			if i == a {
				//fmt.Println("(i, j)=", i, j)
				Usetsk(b, i, j)
			}
			if b == isk[c] {
				c += 1
			}
		}
	}
	te = time.Now()
	fmt.Println("逐次：", te.Sub(ts), "\n")
	PrintArr(Ask)

	//fmt.Println("-----並列-----")
	reset()
	ts = time.Now()
	for a := 1; a < E; a++ {
		c = 1
		for b := 1; b < E; b++ {
			i = c - (isk[c] - b)
			j = c
			if i == a {
				//fmt.Println("(i, j)=", i, j)
				wg.Add(1)
				go UsetskWG(b, i, j, &wg)
			}
			if b == isk[c] {
				c += 1
			}
		}
		wg.Wait()
	}
	te = time.Now()
	fmt.Println("並列：", te.Sub(ts), "\n")
	PrintArr(Ask)
}
