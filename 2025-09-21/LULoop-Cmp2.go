package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup

const size = 8

func SimpleA(A *[size][size]big.Float) {
	//各要素が左上から1, 2, 3, ... と決められる行列を生成

	var n big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			n.Add(&n, big.NewFloat(1))
			A[i][j].SetPrec(1024).Set(&n)
		}
	}
}

func Random(A *[size][size]big.Float) {
	//各要素が乱数で決められる行列を生成

	var a, x1, y1 big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			x2 := rand.Float64()
			x1.SetFloat64(x2)
			y2 := rand.Float64()
			y1.SetFloat64(y2)
			a.SetPrec(1024).Mul(&x1, &y1)
			A[i][j].SetPrec(1024).Set(&a)
		}
	}
}

func Hilbert(A *[size][size]big.Float) {
	//ヒルベルト行列を生成

	var a, n, i2, j2 big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			i2.SetInt64(int64(i))
			j2.SetInt64(int64(j))
			n.Add(&i2, &j2)
			n.Add(&n, big.NewFloat(1))
			a.SetPrec(1024).Quo(big.NewFloat(1), &n)
			A[i][j].SetPrec(1024).Set(&a)
		}

	}
}

func LU(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float) {
	var Aij, Uij, Aji, Lji, c big.Float
	var J int
	for i := 0; i < size; i++ {
		for j := J; j < size; j++ {
			for k := 0; k < size; k++ {
				c.Mul(&L[i][k], &U[k][j])
				Aij.Add(&Aij, &c)
			}
			Uij.Sub(&A[i][j], &Aij)
			U[i][j].Set(&Uij)
			//Aij = *big.NewFloat(0)
			//Aij.Set(big.NewFloat(0))
			Aij.SetString("0")

			if i != j {
				if big.NewFloat(0).Cmp(&U[i][i]) == 0 {
					//Lji = *big.NewFloat(0)
					//Lji.Set(big.NewFloat(0))
					Lji.SetString("0")
				} else {
					for k := 0; k < size; k++ {
						c.Mul(&L[j][k], &U[k][i])
						Aji.Add(&Aji, &c)
					}
					c.Sub(&A[j][i], &Aji)
					Lji.Quo(&c, &U[i][i])
				}
				L[j][i].Set(&Lji)
				//Aji = *big.NewFloat(0)
				//Aji.Set(big.NewFloat(0))
				Aji.SetString("0")
			}
		}
		J += 1
	}
}

func LUgo(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,
	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float) {
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			wg.Add(2)
			go Uset(A, L, U, Lch, Uch, i, j)
			go Lset(A, L, U, Lch, Uch, i, j)
		}
	}
	wg.Wait()
}

func Uset(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,
	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float, i int, j int) {
	defer wg.Done()
	var aij, uij, lij, c big.Float
	aij.SetPrec(1024)
	uij.SetPrec(1024)
	lij.SetPrec(1024)
	c.SetPrec(1024)

	for k := 0; k < size; k++ {
		if k == i {
			lij = *big.NewFloat(0)
			//lij.SetString("0")
		} else {
			lij = <-Lch[i][k]
			Lch[i][k] <- lij
		}
		if lij.Cmp(big.NewFloat(0)) == 0 {
			uij = *big.NewFloat(0)
			//uij.SetString("0")
		} else {
			uij = <-Uch[k][j]
			Uch[k][j] <- uij
		}
		c.Mul(&lij, &uij)
		aij.Add(&aij, &c)
	}
	c.Sub(&A[i][j], &aij)
	U[i][j].Set(&c)
	Uch[i][j] <- c
}

func Lset(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,

	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float, i int, j int) {
	defer wg.Done()
	if i != j {
		var uii, aji, uji, lji, c big.Float
		aji.SetPrec(1024)
		uji.SetPrec(1024)
		lji.SetPrec(1024)
		c.SetPrec(1024)

		uii = <-Uch[i][i]
		Uch[i][i] <- uii
		if big.NewFloat(0).Cmp(&uii) == 0 {
			lji = *big.NewFloat(0)
			//lji.SetString("0")
			//fmt.Println("(i,j)=", i, j)
		} else {
			for k := 0; k < size; k++ {
				if k == i {
					uji = *big.NewFloat(0)
					//uji.SetString("0")
				} else {
					uji = <-Uch[k][i]
					Uch[k][i] <- uji
				}
				if k == i || big.NewFloat(0).Cmp(&uji) == 0 {
					lji = *big.NewFloat(0)
					//lji.SetString("0")
				} else {
					lji = <-Lch[j][k]
					Lch[j][k] <- lji
				}
				c.Mul(&lji, &uji)
				aji.Add(&aji, &c)
			}
			c.Sub(&A[j][i], &aji)
			lji.SetPrec(1024).Quo(&c, &uii)
		}
		L[j][i].Set(&lji)
		Lch[j][i] <- lji
	}
}

func CalcX(L *[size][size]big.Float, U *[size][size]big.Float) {
	//Ax=b (b={(1),(0),(0)...}) におけるxを、LU分解後の行列L, Uから計算

	fmt.Println("x")

	var B [size]big.Float
	for i := 0; i < size; i++ {
		if i == 0 {
			//B[i].Set(big.NewFloat(1))
			B[i].SetString("1")
		} else {
			//B[i].Set(big.NewFloat(0))
			B[i].SetString("0")
		}
	}

	var y1 [size]big.Float
	var mul, sum, y big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			mul.SetPrec(1024).Mul(&L[i][j], &y1[j])
			sum.SetPrec(1024).Add(&sum, &mul)
		}
		y.SetPrec(1024).Sub(&B[i], &sum)
		y1[i].SetPrec(1024).Set(&y)
		sum.SetString("0")
	}

	var x1 [size]big.Float
	var x big.Float
	for i := size - 1; i >= 0; i-- {
		for j := size - 1; j > i; j-- {
			mul.SetPrec(1024).Mul(&U[i][j], &x1[j])
			sum.SetPrec(1024).Add(&sum, &mul)
		}
		x.SetPrec(1024).Sub(&y1[i], &sum)
		x.Quo(&x, &U[i][i])
		x1[i].SetPrec(1024).Set(&x)
		sum.SetString("0")
	}

	for i := 0; i < size; i++ {
		fmt.Println(i, " : ", &x1[i])
	}
}

func Norm(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float) {
	//行列L, Uの積を行列Aから引き、その差のノルムを計算

	fmt.Println("Norm")
	var b, c, sub, norm big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				b.Mul(&L[i][k], &U[k][j])
				c.Add(&c, &b)
			}
			sub.SetPrec(1024).Sub(&A[i][j], &c)
			sub.Mul(&sub, &sub)
			norm.Add(&norm, &sub)
			c.Set(big.NewFloat(0))
		}
	}
	norm.SetPrec(1024).Sqrt(&norm)
	fmt.Println(&norm)
	norm.Set(big.NewFloat(0))
}

func PrintM(M *[size][size]big.Float) {
	//行列をプリント

	for i := 0; i < size; i++ {
		print("\n")
		for j := 0; j < size; j++ {
			fmt.Print(&M[i][j], " ")
		}
	}
	print("\n")
}

func main() {
	var A, L1, U1, L2, U2 [size][size]big.Float

	//行列Aの作り方を指定
	//SimpleA(&A)
	//Random(&A)
	Hilbert(&A)

	//PrintM(&A)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			L1[i][j].SetPrec(1024)
			L2[i][j].SetPrec(1024)
			U1[i][j].SetPrec(1024)
			U2[i][j].SetPrec(1024)
			if i == j {
				L1[i][j].SetString("1")
				L2[i][j].SetString("1")
			}
		}
	}

	t1 := time.Now()
	LU(&A, &L1, &U1)
	fmt.Println("LU:", time.Now().Sub(t1))

	//L1, U1の結果の表示方法を指定
	CalcX(&L1, &U1)
	//Norm(&A, &L1, &U1)

	//PrintM(&L1)
	//PrintM(&U1)

	t2 := time.Now()
	var Lch [size][size]chan big.Float
	var Uch [size][size]chan big.Float

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			Lch[i][j] = make(chan big.Float, 1)
			Uch[i][j] = make(chan big.Float, 1)
		}
	}

	var n big.Float
	n.SetPrec(1024)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				n.SetString("1")
				Lch[i][j] <- n
			} else if i < j {
				n.SetString("0")
				Lch[i][j] <- n
				Uch[j][i] <- n
			}
		}
	}

	LUgo(&A, &L2, &U2, &Lch, &Uch)
	fmt.Println("LUgo:", time.Now().Sub(t2))

	//L2, U2の結果の表示方法を指定
	CalcX(&L2, &U2)
	//Norm(&A, &L2, &U2)

	//PrintM(&L2)
	//PrintM(&U2)
}
