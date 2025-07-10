package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func mullerR(i int, n0 *big.Rat, n1 *big.Rat) {
	//n := new(big.Rat)
	t111 := new(big.Rat)
	t1130 := new(big.Rat)
	t3000 := new(big.Rat)
	t111.SetString("111")
	t1130.SetString("1130")
	t3000.SetString("3000")
	t := new(big.Rat)
	t1 := new(big.Rat)
	t2 := new(big.Rat)
	t2 = n0
	t1 = n1
	if i == 0 {
		t = n0
	} else if i == 1 {
		t = n1
	}
	for j := 0; j < i-1; j++ {
		//fmt.Println("t1,t2:", t1, t2)
		//fmt.Println("t3000,t2:", t3000, t2)
		t.Quo(t3000, t2)
		//fmt.Println("t:", t)
		//fmt.Println("t1130,t:", t1130, t)
		t.Sub(t1130, t)
		//fmt.Println("t:", t)
		//fmt.Println("t,t1:", t, t1)
		t.Quo(t, t1)
		//fmt.Println("t:", t)
		//fmt.Println("t111,t:", t111, t)
		t.Sub(t111, t)
		//fmt.Println("t:", t)
		t2.Set(t1)
		t1.Set(t)
		//fmt.Println("t:", t)
		//fmt.Println(t.FloatString(10))
		//fmt.Println("------")
	}
	fmt.Println(t)
	fmt.Println(t.FloatString(10), "\n")
}

func mullerF(i int, n0 *big.Float, n1 *big.Float) {
	//n := new(big.Rat)
	t111 := new(big.Float)
	t1130 := new(big.Float)
	t3000 := new(big.Float)
	t111.SetString("111")
	t1130.SetString("1130")
	t3000.SetString("3000")
	t := new(big.Float)
	t1 := new(big.Float)
	t2 := new(big.Float)
	t2 = n0
	t1 = n1
	if i == 0 {
		t = n0
	} else if i == 1 {
		t = n1
	}
	for j := 0; j < i-1; j++ {
		//fmt.Println("t1,t2:", t1, t2)
		//fmt.Println("t3000,t2:", t3000, t2)
		t.Quo(t3000, t2)
		//fmt.Println("t:", t)
		//fmt.Println("t1130,t:", t1130, t)
		t.Sub(t1130, t)
		//fmt.Println("t:", t)
		//fmt.Println("t,t1:", t, t1)
		t.Quo(t, t1)
		//fmt.Println("t:", t)
		//fmt.Println("t111,t:", t111, t)
		t.Sub(t111, t)
		//fmt.Println("t:", t)
		t2.Set(t1)
		t1.Set(t)
		//fmt.Println("t:", t)
		//fmt.Println(t)
		//fmt.Println("------")
	}
	fmt.Println(t, "\n")
}

func main() {
	/*
		// big.Int の例
		a := new(big.Int)
		b := new(big.Int)
		c := new(big.Int)

		a.SetString("123456789012345678901234567890", 10)
		b.SetString("987654321098765432109876543210", 10)

		c.Add(a, b)
		fmt.Printf("big.Int Sum: %s\n", c.String()) // 1111111110000000000000000000000

		c.Mul(a, b)
		fmt.Printf("big.Int Product: %s\n", c.String())

		// big.Float の例
		f1 := new(big.Float).SetPrec(200) // 精度を200ビットに設定
		f2 := new(big.Float).SetPrec(200)
		f3 := new(big.Float).SetPrec(200)

		f1.SetFloat64(3.1415926535)
		f2.SetString("2.71828182845904523536028747135266249775724709369995") // より高精度な値

		f3.Add(f1, f2)
		fmt.Printf("big.Float Sum: %.10f\n", f3) // 例として小数点以下10桁まで表示

		f3.Mul(f1, f2)
		fmt.Printf("big.Float Product: %.10f\n", f3)

		// big.Rat の例
		r1 := big.NewRat(1, 3) // 1/3
		r2 := big.NewRat(1, 2) // 1/2
		r3 := new(big.Rat)

		r3.Add(r1, r2)
		fmt.Printf("big.Rat Sum: %s\n", r3.String()) // 5/6
	*/
	for {
		scanner := bufio.NewScanner(os.Stdin)
		print("型を入力してください ( f or r ):")
		scanner.Scan()
		if scanner.Text() == "" {
			break
		} else {
			if scanner.Text() == "r" {
				scanner := bufio.NewScanner(os.Stdin)
				print("数値を入力してください:")
				scanner.Scan()
				n0 := big.NewRat(11, 2)
				n1 := big.NewRat(61, 11)
				i, _ := strconv.Atoi(scanner.Text())
				mullerR(i, n0, n1)
			} else if scanner.Text() == "f" {
				scanner := bufio.NewScanner(os.Stdin)
				print("数値を入力してください:")
				scanner.Scan()
				n0 := big.NewFloat(5.5)
				x := big.NewFloat(61)
				y := big.NewFloat(11)
				n1 := new(big.Float).SetPrec(1024).Quo(x, y)
				i, _ := strconv.Atoi(scanner.Text())
				mullerF(i, n0, n1)
			}
		}
	}
}
