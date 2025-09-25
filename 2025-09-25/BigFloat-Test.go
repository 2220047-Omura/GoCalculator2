package main

import (
	"fmt"
	"math/big"
)

func main() {
	//ビット数の確認
	var a1, a2, a3 big.Float

	a1.SetPrec(1024)
	a1 = *big.NewFloat(1)
	a2.SetString("1")
	a3.SetPrec(1024).SetString("1")

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	var Ach1, Ach2 chan big.Float
	var An1 big.Float
	Ach1 = make(chan big.Float, 1)
	Ach2 = make(chan big.Float, 1)
	An1.SetPrec(1024).SetString("0")

	a1 = *big.NewFloat(0)
	Ach1 <- An1
	a1 = <-Ach1
	fmt.Println(a1)

	a2.SetString("0")
	Ach2 <- An1
	a2 = <-Ach2
	fmt.Println(a2)

	//ビット数の変化に伴うif文の挙動の変化の確認
	var b1, b2, b3 big.Float

	b1 = *big.NewFloat(0)
	if b1.Cmp(big.NewFloat(0)) == 0 {
		fmt.Println("b1.Cmp(big.NewFloat(0)) == 0")
	}
	if big.NewFloat(0).Cmp(&b1) == 0 {
		fmt.Println("big.NewFloat(0).Cmp(&b1) == 0")
	}

	b2.SetString("0")
	if b2.Cmp(big.NewFloat(0)) == 0 {
		fmt.Println("b2.Cmp(big.NewFloat(0)) == 0")
	}
	if big.NewFloat(0).Cmp(&b2) == 0 {
		fmt.Println("big.NewFloat(0).Cmp(&b2) == 0")
	}

	b3.SetPrec(1024).SetString("0")
	if b3.Cmp(big.NewFloat(0)) == 0 {
		fmt.Println("b3.Cmp(big.NewFloat(0)) == 0")
	}
	if big.NewFloat(0).Cmp(&b3) == 0 {
		fmt.Println("big.NewFloat(0).Cmp(&b3) == 0")
	}

	//----LULoop-Cmp2を調べると、170行目だけ結果が間違ってしまうことがわかった----

	//big.Float型の数cの与え方の変化に伴う乗算の変化の確認
	var c1, c2, c3, c4, c5 big.Float

	c1.SetString("0")
	c2 = *big.NewFloat(0)
	c3.SetString("3")
	c4 = *big.NewFloat(3)
	c5.Mul(&c2, &c4)
	fmt.Println("c5 = ", c5)

	c2.SetPrec(64)
	c4.SetPrec(64)

	//ビット数の変更について確認
	var d big.Float

	d.SetPrec(1024)
	fmt.Println(d)
	//d = *big.NewFloat(1)
	d.SetString("0")
	fmt.Println(d)
	d.SetPrec(1024).SetFloat64(3.14)
	fmt.Println(d)
}
