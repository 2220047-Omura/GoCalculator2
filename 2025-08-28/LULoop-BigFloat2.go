package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup

const size = 700

func Uset(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,
	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float, i int, j int) {
	defer wg.Done()
	var aij, uij, lij, c big.Float
	for k := 0; k < size; k++ {
		if k == i {
			lij = *big.NewFloat(0)
		} else {
			lij = <-Lch[i][k]
			Lch[i][k] <- lij
		}
		if lij.Cmp(big.NewFloat(0)) == 0 {
			uij = *big.NewFloat(0)
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
		var aji, uji, lji, c big.Float
		uii := <-Uch[i][i]
		Uch[i][i] <- uii
		if big.NewFloat(0).Cmp(&uii) == 0 {
			lji = *big.NewFloat(0)
			//fmt.Println("(i,j)=", i, j)
		} else {
			for k := 0; k < size; k++ {
				if k == i {
					uji = *big.NewFloat(0)
				} else {
					uji = <-Uch[k][i]
					Uch[k][i] <- uji
				}
				if k == i || big.NewFloat(0).Cmp(&uji) == 0 {
					lji = *big.NewFloat(0)
				} else {
					lji = <-Lch[j][k]
					Lch[j][k] <- lji
				}
				c.Mul(&lji, &uji)
				aji.Add(&aji, &c)
			}
			c.Sub(&A[j][i], &aji)
			lji.Quo(&c, &uii)
		}
		L[j][i].Set(&lji)
		Lch[j][i] <- lji
	}
}

func LUgo(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,
	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float) {
	var J int
	for i := 0; i < size; i++ {
		for j := J; j < size; j++ {
			wg.Add(2)
			go Uset(A, L, U, Lch, Uch, i, j)
			go Lset(A, L, U, Lch, Uch, i, j)
		}
		J += 1
	}
	wg.Wait()
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
			Aij = *big.NewFloat(0)

			if i != j {
				if big.NewFloat(0).Cmp(&U[i][i]) == 0 {
					Lji = *big.NewFloat(0)
				} else {
					for k := 0; k < size; k++ {
						c.Mul(&L[j][k], &U[k][i])
						Aji.Add(&Aji, &c)
					}
					c.Sub(&A[j][i], &Aji)
					Lji.Quo(&c, &U[i][i])
				}
				L[j][i].Set(&Lji)
				Aji = *big.NewFloat(0)
			}
		}
		J += 1
	}
}

func main() {
	var A, L1, U1, L2, U2 [size][size]big.Float
	var a, x1, y1 big.Float
	//var n big.Float

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			x2 := rand.Float64()
			x1.SetFloat64(x2)
			y2 := rand.Float64()
			y1.SetFloat64(y2)
			a.SetPrec(1024).Mul(&x1, &y1)
			A[i][j] = a
			if i == j {
				L1[i][j].Set(big.NewFloat(1))
				L2[i][j].Set(big.NewFloat(1))
			}
		}
	}

	/*
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				n.Add(&n, big.NewFloat(1))
				A[i][j].Set(&n)
				if i == j {
					L1[i][j].Set(big.NewFloat(1))
					L2[i][j].Set(big.NewFloat(1))
				}
			}

		}
	*/

	//fmt.Println(&A[0][0], &A[0][1], &A[0][2], "\n", &A[1][0], &A[1][1], &A[1][2], "\n", &A[2][0], &A[2][1], &A[2][2])
	//fmt.Println(&B[0], &B[1], &B[4])

	t1 := time.Now()
	LU(&A, &L1, &U1)
	fmt.Println("LU:", time.Now().Sub(t1))
	//fmt.Println(&L1[0][0], &L1[0][1], &L1[0][2], "\n", &L1[1][0], &L1[1][1], &L1[1][2], "\n", &L1[2][0], &L1[2][1], &L1[2][2])
	//fmt.Println(&U1[0][0], &U1[0][1], &U1[0][2], "\n", &U1[1][0], &U1[1][1], &U1[1][2], "\n", &U1[2][0], &U1[2][1], &U1[2][2])

	var b, c, sub, norm big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				b.Mul(&L1[i][k], &U1[k][j])
				c.Add(&c, &b)
			}
			//fmt.Println(&A[i][j], &c)
			sub.SetPrec(1024).Sub(&A[i][j], &c)
			sub.SetPrec(1024).Mul(&sub, &sub)
			norm.Add(&norm, &sub)
			c.Set(big.NewFloat(0))
		}
	}
	norm.SetPrec(1024).Sqrt(&norm)
	fmt.Println(&norm)
	norm.Set(big.NewFloat(0))

	t2 := time.Now()
	var Lch [size][size]chan big.Float
	var Uch [size][size]chan big.Float

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			Lch[i][j] = make(chan big.Float, 1)
			Uch[i][j] = make(chan big.Float, 1)
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				Lch[i][j] <- *big.NewFloat(1)
			}
			if i < j {
				Lch[i][j] <- *big.NewFloat(0)
				Uch[j][i] <- *big.NewFloat(0)
			}
		}
	}
	LUgo(&A, &L2, &U2, &Lch, &Uch)
	fmt.Println("LUgo:", time.Now().Sub(t2))
	//fmt.Println(&L2[0][0], &L2[0][1], &L2[0][2], "\n", &L2[1][0], &L2[1][1], &L2[1][2], "\n", &L2[2][0], &L2[2][1], &L2[2][2])
	//fmt.Println(&U2[0][0], &U2[0][1], &U2[0][2], "\n", &U2[1][0], &U2[1][1], &U2[1][2], "\n", &U2[2][0], &U2[2][1], &U2[2][2])

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				b.Mul(&L2[i][k], &U2[k][j])
				c.Add(&c, &b)
			}
			//fmt.Println(&A[i][j], &c)
			sub.Sub(&A[i][j], &c)
			sub.Mul(&sub, &sub)
			norm.Add(&norm, &sub)
			c.Set(big.NewFloat(0))
		}
	}
	norm.SetPrec(1024).Sqrt(&norm)
	fmt.Println(&norm)
	norm.Set(big.NewFloat(0))
}
