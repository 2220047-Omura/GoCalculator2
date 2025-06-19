package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var Pi1, Pi2 int

func MCgo() {
	x := rand.Float64()
	X := math.Pow(x, 2)
	y := rand.Float64()
	Y := math.Pow(y, 2)
	if X+Y <= 1 {
		Pi1 += 1
	}
	wg.Done()
}

func MC() {
	x := rand.Float64()
	X := math.Pow(x, 2)
	y := rand.Float64()
	Y := math.Pow(y, 2)
	if X+Y <= 1 {
		Pi2 += 1
	}
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "" {
			break
		} else {
			n, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("数値を入力してください")
			} else {
				fmt.Println("MCgo")
				wg.Add(n)
				t1 := time.Now()
				for i := 0; i < n; i++ {
					go MCgo()
				}
				wg.Wait()
				t2 := time.Now()
				var Ans1 float64 = (float64(Pi1) / float64(n)) * 4
				fmt.Println(Ans1)
				fmt.Println(t2.Sub(t1))
				fmt.Println("MC")
				t3 := time.Now()
				for i := 0; i < n; i++ {
					MC()
				}
				t4 := time.Now()
				var Ans2 float64 = (float64(Pi2) / float64(n)) * 4
				fmt.Println(Ans2)
				fmt.Println(t4.Sub(t3), "\n")
			}
		}
		Pi1, Pi2 = 0, 0
	}
}
