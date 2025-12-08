package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"sync"
	"time"
)

var Ask [size2]big.Float
var isk [size]int
var Lsk [size2]big.Float
var MULsk [size2]big.Float
var SUMsk [size2]big.Float

const size = 5
const size2 = 10

func initialize() {

	//setSkyline()
	setSkylineTest()

	for i := 0; i < size2; i++ {
		Lsk[i].SetString("0")
		MULsk[i].SetString("0")
		SUMsk[i].SetString("0")
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
			//A[i][j].SetPrec(1024).Mul(&a, &b)
		}
	}
}

func setSkylineTest() {
	var AskTest = [size2]int{2, 1, 3, 0, 4, 7, 8, 2, 3, 5}
	for i := 0; i < size2; i++ {
		Ask[i].SetInt64(int64(AskTest[i]))
		SUMsk[i].SetString("0")
		MULsk[i].SetString("0")
	}
	fmt.Println(AskTest)
	var iskTest = [size]int{0, 1, 4, 5, 9}
	for i := 0; i < size; i++ {
		isk[i] = iskTest[i]
	}
}

func Usetsk(a int, i int, j int) {
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
		Lsk[a].Quo(&Ask[isk[i]-(s-k)], &Ask[isk[i-(s-k)]])
		MULsk[a].Mul(&Lsk[a], &Ask[isk[j]-(j-i)-(s-k)])
		SUMsk[a].Add(&SUMsk[a], &MULsk[a])
	}
	Ask[a].Sub(&Ask[a], &SUMsk[a])
}

func UsetskWG(a int, i int, j int, wg *sync.WaitGroup) {
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
		Lsk[a].Quo(&Ask[isk[i]-(s-k)], &Ask[isk[i-(s-k)]])
		MULsk[a].Mul(&Lsk[a], &Ask[isk[j]-(j-i)-(s-k)])
		SUMsk[a].Add(&SUMsk[a], &MULsk[a])
	}
	Ask[a].Sub(&Ask[a], &SUMsk[a])
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
	var i, j, c int
	var wg sync.WaitGroup

	//fmt.Println("-----逐次-----")
	initialize()
	ts = time.Now()
	for a := 1; a < size2; a++ {
		c = 1
		for b := 1; b < size2; b++ {
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
	PrintArr(&Ask)

	//fmt.Println("-----並列-----")
	initialize()
	ts = time.Now()
	for a := 1; a < size2; a++ {
		c = 1
		for b := 1; b < size2; b++ {
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
	PrintArr(&Ask)
}
