package main

import (
	"fmt"
	"math/big"
)

func main() {
	//ビット数の確認
	var a1, a2, a3, a4, a0 big.Float // zero values have been set
	// fmt.Println(a1.Prec(), a2.Prec(), a3.Prec()) // 0 0 0, because a1, a2, a3 are zero-valued
	fmt.Printf("(1) a1=%v, a2=%v, a3=%v\n", a1, a2, a3)

	p := &a1
	fmt.Printf("(2) &a1=%p, p=%p, a1=%v *p=%v\n", &a1, p, a1, *p)
	a1.SetPrec(1024)
	fmt.Printf("(3) &a1=%p, p=%p, a1=%v *p=%v\n", &a1, p, a1, *p)
	a0 = a1
	fmt.Printf("(4) &a1=%p, p=%p, a1=%v &a0=%p a0=%v\n", &a1, p, a1, &a0, a0)
	x := *big.NewFloat(1.1) // modify the contents of a1
	a1 = x
	a1.SetPrec(20)
	fmt.Printf("(5) &a1=%p, p=%p, a1=%v &a0=%p a0=%v x=%v\n", &a1, p, a1, &a0, a0, x)
	p2 := &a2
	fmt.Println("&a2, p2, a2=", &a2, p2, a2)
	fmt.Printf("(6) &a2=%p, p2=%p, a2=%v\n", &a2, p2, a2)
	a2.SetString("1.1")
	fmt.Printf("(7) &a2=%p, p2=%p, a2=%v\n", &a2, p2, a2)
	a3.SetPrec(1024).SetString("1.7")
	// a4.SetString("1.1").SetPrec(1024)
	a4.SetString("1.7")
	// (*a4p).SetString("1.1")
	// a4p.SetString("1.1")
	a4.SetPrec(1024)
	fmt.Printf("(8) &a3=%p, &a4=%p, a3=%v a4=%v\n", &a3, &a4, a3, a4)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println()

	var Ach1, Ach2 chan big.Float
	var An1 big.Float
	Ach1 = make(chan big.Float, 1)
	Ach2 = make(chan big.Float, 1)
	An1.SetPrec(1024).SetString("1.1")

	a1 = *big.NewFloat(2.2)
	Ach1 <- An1
	a1 = <-Ach1
	fmt.Println(a1)

	a2.SetString("3.3")
	Ach2 <- An1
	a2 = <-Ach2
	fmt.Println(a2)

	fmt.Println()

	//ビット数の変化に伴うif文の挙動の変化の確認
	var b1, b2, b3 big.Float

	b1 = *big.NewFloat(1.1)
	if b1.Cmp(big.NewFloat(1.1)) == 0 {
		fmt.Println("b1.Cmp(big.NewFloat(1.1)) == 0")
	}
	if big.NewFloat(1.1).Cmp(&b1) == 0 {
		fmt.Println("big.NewFloat(1.1).Cmp(&b1) == 0")
	}

	b2.SetString("1.1") // 64bit
	// b2.SetPrec(53)
	if b2.Cmp(big.NewFloat(1.1)) == 0 {
		fmt.Println("b2.Cmp(big.NewFloat(1.1)) == 0")
	}
	if big.NewFloat(1.1).Cmp(&b2) == 0 {
		fmt.Println("big.NewFloat(1.1).Cmp(&b2) == 0")
	}

	b3.SetPrec(1024).SetString("1.1")
	// b3.SetPrec(53)
	if b3.Cmp(big.NewFloat(1.1)) == 0 {
		fmt.Println("b3.Cmp(big.NewFloat(1.1)) == 0")
	}
	if big.NewFloat(1.1).Cmp(&b3) == 0 {
		fmt.Println("big.NewFloat(1.1).Cmp(&b3) == 0")
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
	fmt.Printf("d=%v\n", d)
	var d2 big.Float
	d2.SetPrec(1024).SetString("3.14")
	fmt.Printf("d2=%v\n", d2)

	d2.SetString("8")
	fmt.Printf("d2=%v\n", d2)

	var y1, y2 big.Float
	y1.SetPrec(1024).SetString("3.14")
	y2.SetPrec(1024).SetString("0")
	fmt.Println("y1 = ", &y1)
	z := &y1
	w := 4.2
	fmt.Printf("y1 = %f\n", &y1)
	fmt.Printf("w = %f\n", w)
	fmt.Printf("z = %f\n", z)
	//fmt.Printf("w*z = %f\n", w*z)
	y1.Mul(&y1, &y2)
	fmt.Println("y1 = ", &y1)
	fmt.Println("y2 = ", &y2)
}
