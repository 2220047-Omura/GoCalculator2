package main

import (
	"bufio"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strconv"

	//"sync"
	"time"
)

// var wg sync.WaitGroup
var Pi1, Pi2, Pi1Sum int

func MCgo(c chan int) {
	local_Pi := 0
	X := new(big.Float).SetPrec(1024)
	Y := new(big.Float).SetPrec(1024)
	Z := new(big.Float).SetPrec(1024)
	for i := 0; i < 10000; i++ {
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

func MC() {
	X := new(big.Float).SetPrec(1024)
	Y := new(big.Float).SetPrec(1024)
	Z := new(big.Float).SetPrec(1024)
	for i := 0; i < 10000; i++ {
		x := rand.Float64()
		xBig := new(big.Float).SetFloat64(x)
		X.Mul(xBig, xBig)
		y := rand.Float64()
		yBig := new(big.Float).SetFloat64(y)
		Y.Mul(yBig, yBig)
		Z.Add(X, Y)
		cmp := Z.Cmp(big.NewFloat(1))
		if cmp != 1 {
			Pi2 += 1
		}
	}
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("数値を入力してください")
		scanner.Scan()
		if scanner.Text() == "" {
			break
		} else {
			n, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("数値を入力してください")
			} else {
				fmt.Println("MCgo")
				//wg.Add(n)
				c := make(chan int, n)
				t1 := time.Now()
				for i := 0; i < n; i++ {
					go MCgo(c)
					Pi1 = <-c
					Pi1Sum += Pi1Sum
				}
				//wg.Wait()
				t2 := time.Now()
				var Ans1 float64 = (float64(Pi1Sum) / float64(n)) * 4 / 10000
				fmt.Println(Ans1)
				fmt.Println(t2.Sub(t1))
				fmt.Println("MC")
				t3 := time.Now()
				for i := 0; i < n; i++ {
					MC()
				}
				t4 := time.Now()
				var Ans2 float64 = (float64(Pi2) / float64(n)) * 4 / 10000
				fmt.Println(Ans2)
				fmt.Println(t4.Sub(t3), "\n")
			}
		}
		Pi1, Pi2, Pi1Sum = 0, 0, 0
	}
}
