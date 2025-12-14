package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"sync"
	"time"
)

var Ask []big.Float
var isk [size]int
var Lsk []big.Float
var MULsk []big.Float
var SUMsk []big.Float

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
	var zero big.Float
	zero.SetString("0")
	for i := 0; i < E; i++ {
		Ask = append(Ask, zero)
		Lsk = append(Lsk, zero)
		MULsk = append(MULsk, zero)
		SUMsk = append(SUMsk, zero)
	}
}

func reset() {
	//fmt.Println(N)
	setSkyline()
	mulDiagonal()
	//setSkylineTest()
	for i := 0; i < E; i++ {
		Lsk[i].SetString("0")
		MULsk[i].SetString("0")
		SUMsk[i].SetString("0")
	}
}

func setSkyline() {
	var n int
	isk[0] = 0
	for i := 1; i < size; i++ {
		n -= 1
		if n < 0 {
			n = i
		}
		isk[i] = isk[i-1] + n + 1
	}

	var a, b big.Float
	for i := 0; i < E; i++ {
		r := rand.Float64()
		a.SetFloat64(r)
		r = rand.Float64()
		b.SetFloat64(r)
		Ask[i].SetPrec(1024).Quo(&a, &b)
	}

}

func mulDiagonal() {
	var hdr, one, toIsk, toI, div big.Float
	hdr.SetString("100")
	one.SetString("1")
	c := 0
	for i := 0; i < E; i++ {
		toIsk.SetInt64(int64(isk[c]))
		toI.SetInt64(int64(i))
		div.Sub(&toIsk, &toI)
		div.Add(&div, &one)
		div.Quo(&hdr, &div)
		Ask[i].Mul(&Ask[i], &div)
		if i == isk[c] {
			c += 1
		}
	}
}

func setSkylineTest() {
	var AskTest = []int{2, 1, 3, 0, 4, 7, 8, 2, 3, 5}
	//N = 10
	for i := 0; i < E; i++ {
		Ask[i].SetInt64(int64(AskTest[i]))
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
		//fmt.Println("Lki*Ukj: ", isk[i]-(s-k), isk[j]-(j-i)-(s-k))
		//fmt.Println("i,Lki,Uii: ", i-(s-k), &Ask[isk[i]-(s-k)], &Ask[isk[i-(s-k)]])
		Lsk[b].Quo(&Ask[isk[i]-(s-k)], &Ask[isk[i-(s-k)]])
		MULsk[b].Mul(&Lsk[b], &Ask[isk[j]-(j-i)-(s-k)])
		SUMsk[b].Add(&SUMsk[b], &MULsk[b])
	}
	Ask[b].Sub(&Ask[b], &SUMsk[b])
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
		//fmt.Println("Lki*Ukj: ", isk[i]-(s-k), isk[j]-(j-i)-(s-k))
		//fmt.Println("i,Lki,Uii: ", i-(s-k), &Ask[isk[i]-(s-k)], &Ask[isk[i-(s-k)]])
		Lsk[b].Quo(&Ask[isk[i]-(s-k)], &Ask[isk[i-(s-k)]])
		MULsk[b].Mul(&Lsk[b], &Ask[isk[j]-(j-i)-(s-k)])
		SUMsk[b].Add(&SUMsk[b], &MULsk[b])
	}
	Ask[b].Sub(&Ask[b], &SUMsk[b])
}

func PrintArr(M []big.Float) {
	//行列をプリント
	//for i := 0; i < E; i++ {
	for i := E - 3; i < E; i++ {
		fmt.Print(&M[i], " ")
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
