package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"

	//"sync"
	"time"
)

// var wg sync.WaitGroup

func MCFloat(n int, c chan int) {
	local_Pi := 0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		X := math.Pow(x, 2)
		y := rand.Float64()
		Y := math.Pow(y, 2)
		if X+Y <= 1 {
			local_Pi += 1
		}
	}
	c <- local_Pi
}

func MCBig(n int, c chan int) {
	local_Pi := 0
	X := new(big.Float).SetPrec(1024)
	Y := new(big.Float).SetPrec(1024)
	Z := new(big.Float).SetPrec(1024)
	for i := 0; i < n; i++ {
		x := rand.Float64()
		xBig := new(big.Float).SetFloat64(x)
		X.Mul(xBig, xBig)
		y := rand.Float64()
		yBig := new(big.Float).SetFloat64(y)
		Y.Mul(yBig, yBig)
		Z.Add(X, Y)
		cmp := Z.Cmp(big.NewFloat(1))
		if cmp != 1 {
			local_Pi += 1
		}
	}
	c <- local_Pi
}

func MCgoFloat(n int, c chan int) {
	local_Pi := 0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		X := math.Pow(x, 2)
		y := rand.Float64()
		Y := math.Pow(y, 2)
		if X+Y <= 1 {
			local_Pi += 1
		}
	}
	c <- local_Pi
	//wg.Done()
}

func MCgoBig(n int, c chan int) {
	local_Pi := 0
	X := new(big.Float).SetPrec(1024)
	Y := new(big.Float).SetPrec(1024)
	Z := new(big.Float).SetPrec(1024)
	for i := 0; i < n; i++ {
		x := rand.Float64()
		xBig := new(big.Float).SetFloat64(x)
		X.Mul(xBig, xBig)
		y := rand.Float64()
		yBig := new(big.Float).SetFloat64(y)
		Y.Mul(yBig, yBig)
		Z.Add(X, Y)
		cmp := Z.Cmp(big.NewFloat(1))
		if cmp != 1 {
			local_Pi += 1
		}
	}
	c <- local_Pi
	//wg.Done()
}

func main() {
	var Pi1, Pi1Sum, Pi2, Pi2Sum, Pi3, Pi3Sum, Pi4, Pi4Sum int
	var plot = 10000000 //打つ点の数
	var t = 1           //スレッド数
	var n = plot / t    //スレッド内の必要ループ数

	c1 := make(chan int, t)
	c2 := make(chan int, t)
	c3 := make(chan int, t)
	c4 := make(chan int, t)

	fmt.Println("MCFloat")
	t1 := time.Now()
	for i := 0; i < t; i++ {
		MCFloat(n, c3)
	}
	for i := 0; i < t; i++ {
		Pi1 = <-c3
		Pi1Sum += Pi1
	}
	t2 := time.Now()
	var Ans1 float64 = (float64(Pi1Sum) / float64(plot)) * 4
	fmt.Println(Ans1)
	fmt.Println(t2.Sub(t1), "\n")

	fmt.Println("MCBig")
	t3 := time.Now()
	for i := 0; i < t; i++ {
		MCBig(n, c4)
	}
	for i := 0; i < t; i++ {
		Pi2 = <-c4
		Pi2Sum += Pi2
	}
	t4 := time.Now()
	var Ans2 float64 = (float64(Pi2Sum) / float64(plot)) * 4
	fmt.Println(Ans2)
	fmt.Println(t4.Sub(t3), "\n")

	fmt.Println("MCgoFloat")
	t5 := time.Now()
	for i := 0; i < t; i++ {
		go MCgoFloat(n, c1)
	}
	for i := 0; i < t; i++ {
		Pi3 = <-c1
		Pi3Sum += Pi3
	}
	t6 := time.Now()
	var Ans3 float64 = (float64(Pi3Sum) / float64(plot)) * 4
	fmt.Println(Ans3)
	fmt.Println(t6.Sub(t5), "\n")

	fmt.Println("MCgoBig")
	t7 := time.Now()
	for i := 0; i < t; i++ {
		go MCgoBig(n, c2)
	}
	for i := 0; i < t; i++ {
		Pi4 = <-c2
		Pi4Sum += Pi4
	}
	t8 := time.Now()
	var Ans4 float64 = (float64(Pi4Sum) / float64(plot)) * 4
	fmt.Println(Ans4)
	fmt.Println(t8.Sub(t7), "\n")
}
