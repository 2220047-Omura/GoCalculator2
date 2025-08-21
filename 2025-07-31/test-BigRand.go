package main

import (
	//"crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

func main() {

	/*
		max := new(big.Int)
		max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
		for i := 0; i < 5; i++ {
			n, _ := rand.Int(rand.Reader, max)

			print(n, "\n")
		}
	*/
	//var x [2][2]big.Int
	//var max *big.Int = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(130), nil)
	/*
		for i := 0; i < 5; i++ {
			n, _ := rand.Int(rand.Reader, max)
			fmt.Println(n)
		}
	*/
	//n, _ := rand.Int(rand.Reader, max)
	//fmt.Println(n)
	//x[0][0] = *n

	/*
		fmt.Println(&x[0][0])
		fmt.Println(&x[0][1])
		a := big.NewInt(1)
		b := big.NewInt(2)
		c := big.NewInt(3)
		a.Mul(b, c)
		fmt.Println(a)
		var d big.Int
		fmt.Println(d.Cmp(a))
		e := big.NewInt(4)
		f := big.NewInt(5)
		a.Mul(e, f)
		fmt.Println(a)
		if a.Cmp(big.NewInt(20)) == 0 {
			print("あ\n")
		}
		if big.NewInt(20).Cmp(a) == 0 {
			print("い\n")
		}
		a.Div(b, c)
		fmt.Println(a)
	*/

	var A, B1, C1 big.Float
	for i := 0; i < 10; i++ {
		B2 := rand.Float64()
		C2 := rand.Float64()
		B1.SetFloat64(B2)
		C1.SetFloat64(C2)
		A.SetPrec(1024).Mul(&B1, &C1)
		fmt.Println(&A)
	}
}
