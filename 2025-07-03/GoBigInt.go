package main

import (
	"fmt"
	"math/big"
)

func main() {
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

}
