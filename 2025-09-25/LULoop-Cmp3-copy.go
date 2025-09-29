package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup

const size = 500

func SimpleA(A *[size][size]big.Float) {
	//各要素が左上から1, 2, 3, ... と決められる行列を生成

	var n big.Float
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			n.Add(&n, big.NewFloat(1))
			A[i][j].SetPrec(1024).Set(&n)
		}
	}

	/*
		np := new(big.Float).SetPrec(1024)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				np.Add(np, big.NewFloat(1))
				A[i][j].SetPrec(1024).Set(np)
			}
		}
	*/
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
	one := big.NewFloat(1)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			i2.SetInt64(int64(i))
			j2.SetInt64(int64(j))
			n.Add(&i2, &j2)
			// n.Add(&n, big.NewFloat(1))
			n.Add(&n, one)
			// a.SetPrec(1024).Quo(big.NewFloat(1), &n)
			a.SetPrec(1024).Quo(one, &n)
			A[i][j].SetPrec(1024).Set(&a)
		}

	}
}

func LU_NF(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float) {
	var Aij, Uij, Aji, Lji, c big.Float
	var J int
	zero := *big.NewFloat(0)
	for i := 0; i < size; i++ {
		for j := J; j < size; j++ {
			for k := 0; k < size; k++ {
				c.Mul(&L[i][k], &U[k][j])
				Aij.Add(&Aij, &c)
			}
			Uij.Sub(&A[i][j], &Aij)
			U[i][j].Set(&Uij)
			// Aij = *big.NewFloat(0)
			Aij = zero

			if i != j {
				// if big.NewFloat(0).Cmp(&U[i][i]) == 0 {
				if zero.Cmp(&U[i][i]) == 0 {
					// Lji = *big.NewFloat(0)
					Lji = zero
				} else {
					for k := 0; k < size; k++ {
						c.Mul(&L[j][k], &U[k][i])
						Aji.Add(&Aji, &c)
					}
					c.Sub(&A[j][i], &Aji)
					Lji.Quo(&c, &U[i][i])
				}
				L[j][i].Set(&Lji)
				// Aji = *big.NewFloat(0)
				Aji = zero
			}
		}
		J += 1
	}
}

func LU_Str(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float) {
	var Aij, Uij, Aji, Lji, c big.Float
	var J int
	var zero big.Float
	zero.SetString("0")
	for i := 0; i < size; i++ {
		for j := J; j < size; j++ {
			for k := 0; k < size; k++ {
				c.Mul(&L[i][k], &U[k][j])
				Aij.Add(&Aij, &c)
			}
			Uij.Sub(&A[i][j], &Aij)
			U[i][j].Set(&Uij)
			// Aij.SetString("0")
			Aij = zero

			if i != j {
				// if big.NewFloat(0).Cmp(&U[i][i]) == 0 {
				if zero.Cmp(&U[i][i]) == 0 {
					// Lji.SetString("0")
					Lji = zero
				} else {
					for k := 0; k < size; k++ {
						c.Mul(&L[j][k], &U[k][i])
						Aji.Add(&Aji, &c)
					}
					c.Sub(&A[j][i], &Aji)
					Lji.Quo(&c, &U[i][i])
				}
				L[j][i].Set(&Lji)
				// Aji.SetString("0")
				Aji = zero
			}
		}
		J += 1
	}
}

func LUgo_NF(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,
	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float) {
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			wg.Add(2)
			go Uset_NF(A, L, U, Lch, Uch, i, j)
			go Lset_NF(A, L, U, Lch, Uch, i, j)
		}
	}
	wg.Wait()
}

func Uset_NF(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,
	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float, i int, j int) {
	defer wg.Done()
	var aij, uij, lij, c big.Float
	zero := *big.NewFloat(0)
	aij.SetPrec(1024)
	uij.SetPrec(1024)
	lij.SetPrec(1024)
	c.SetPrec(1024)

	for k := 0; k < size; k++ {
		if k == i {
			// lij = *big.NewFloat(0)
			lij = zero
		} else {
			lij = <-Lch[i][k]
			Lch[i][k] <- lij
		}
		// if lij.Cmp(big.NewFloat(0)) == 0 {
		if lij.Cmp(&zero) == 0 {
			// uij = *big.NewFloat(0)
			uij = zero
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

func Lset_NF(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,

	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float, i int, j int) {
	defer wg.Done()
	if i != j {
		var uii, aji, uji, lji, lji2, c big.Float
		zero := *big.NewFloat(0)
		aji.SetPrec(1024)
		uji.SetPrec(1024)
		lji.SetPrec(1024)
		lji2.SetPrec(1024)
		c.SetPrec(1024)

		uii = <-Uch[i][i]
		Uch[i][i] <- uii
		// if big.NewFloat(0).Cmp(&uii) == 0 {
		if zero.Cmp(&uii) == 0 {
			// lji = *big.NewFloat(0)
			lji = zero
		} else {
			for k := 0; k < size; k++ {
				if k == i {
					// uji = *big.NewFloat(0)
					uji = zero
				} else {
					uji = <-Uch[k][i]
					Uch[k][i] <- uji
				}
				// if k == i || big.NewFloat(0).Cmp(&uji) == 0 {
				if k == i || zero.Cmp(&uji) == 0 {
					// lji = *big.NewFloat(0)
					lji = zero
				} else {
					lji = <-Lch[j][k]
					Lch[j][k] <- lji
				}
				//fmt.Println(&lji, lji) //ljiの構造を確認できる
				c.Mul(&lji, &uji)
				aji.Add(&aji, &c)
			}
			c.Sub(&A[j][i], &aji)
			lji2.SetPrec(1024).Quo(&c, &uii)
		}
		L[j][i].Set(&lji2)
		Lch[j][i] <- lji2
	}
}

func LUgo_Str(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,
	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float) {
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			wg.Add(2)
			go Uset_Str(A, L, U, Lch, Uch, i, j)
			go Lset_Str(A, L, U, Lch, Uch, i, j)
		}
	}
	wg.Wait()
}

func Uset_Str(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,
	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float, i int, j int) {
	defer wg.Done()
	var aij, uij, lij, c big.Float
	var zero big.Float
	zero.SetString("0")
	aij.SetPrec(1024)
	uij.SetPrec(1024)
	lij.SetPrec(1024)
	c.SetPrec(1024)

	for k := 0; k < size; k++ {
		if k == i {
			// lij.SetString("0")
			lij = zero
		} else {
			lij = <-Lch[i][k]
			Lch[i][k] <- lij
		}
		if lij.Cmp(big.NewFloat(0)) == 0 {
			// uij.SetString("0")
			uij = zero
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

func Lset_Str(A *[size][size]big.Float, L *[size][size]big.Float, U *[size][size]big.Float,

	Lch *[size][size]chan big.Float, Uch *[size][size]chan big.Float, i int, j int) {
	defer wg.Done()
	if i != j {
		var uii, aji, uji, lji, lji2, c big.Float
		var zero big.Float
		zero.SetString("0")
		aji.SetPrec(1024)
		uji.SetPrec(1024)
		lji.SetPrec(1024)
		lji2.SetPrec(1024)
		c.SetPrec(1024)

		uii = <-Uch[i][i]
		Uch[i][i] <- uii
		if big.NewFloat(0).Cmp(&uii) == 0 {
			// lji.SetString("0")
			lji = zero
		} else {
			for k := 0; k < size; k++ {
				if k == i {
					// uji.SetString("0")
					uji = zero
				} else {
					uji = <-Uch[k][i]
					Uch[k][i] <- uji
				}
				if k == i || big.NewFloat(0).Cmp(&uji) == 0 {
					// lji.SetString("0")
					lji = zero
				} else {
					lji = <-Lch[j][k]
					Lch[j][k] <- lji
				}
				//fmt.Println(&lji, lji) //ljiの構造を確認できる
				c.Mul(&lji, &uji)
				aji.Add(&aji, &c)
			}
			c.Sub(&A[j][i], &aji)
			lji2.SetPrec(1024).Quo(&c, &uii)
		}
		L[j][i].Set(&lji2)
		Lch[j][i] <- lji2
	}
}

func CalcX(L *[size][size]big.Float, U *[size][size]big.Float) {
	//Ax=b (b={(1),(0),(0)...}) におけるxを、LU分解後の行列L, Uから計算

	fmt.Println("x")

	var B [size]big.Float
	for i := 0; i < size; i++ {
		if i == 0 {
			B[i].SetString("1")
		} else {
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
			c.SetString("0")
		}
	}
	norm.SetPrec(1024).Sqrt(&norm)
	fmt.Println(&norm)
	norm.SetString("0")
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
	var A, L_NF, U_NF, L_Str, U_Str, Lgo_NF, Ugo_NF, Lgo_Str, Ugo_Str [size][size]big.Float

	//行列Aの作り方を指定
	//SimpleA(&A)
	// Random(&A)
	Hilbert(&A)

	//PrintM(&A)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			L_NF[i][j].SetPrec(1024)
			L_Str[i][j].SetPrec(1024)
			Lgo_NF[i][j].SetPrec(1024)
			Lgo_Str[i][j].SetPrec(1024)
			U_NF[i][j].SetPrec(1024)
			U_Str[i][j].SetPrec(1024)
			Ugo_NF[i][j].SetPrec(1024)
			Ugo_Str[i][j].SetPrec(1024)
			if i == j {
				L_NF[i][j].SetString("1")
				L_Str[i][j].SetString("1")
				Lgo_NF[i][j].SetString("1")
				Lgo_Str[i][j].SetString("1")
			}
		}
	}

	var t time.Time

	t = time.Now()
	LU_NF(&A, &L_NF, &U_NF)
	fmt.Println("LU_NF:", time.Now().Sub(t))
	Norm(&A, &L_NF, &U_NF)

	t = time.Now()
	LU_Str(&A, &L_Str, &U_Str)
	fmt.Println("LU_Str:", time.Now().Sub(t))
	Norm(&A, &L_Str, &U_Str)

	//L1, U1の結果の表示方法を指定
	// CalcX(&L1, &U1)
	// Norm(&A, &L1, &U1)

	//PrintM(&L1)
	//PrintM(&U1)

	t = time.Now()
	var Lch_NF [size][size]chan big.Float
	var Uch_NF [size][size]chan big.Float

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			Lch_NF[i][j] = make(chan big.Float, 1)
			Uch_NF[i][j] = make(chan big.Float, 1)
		}
	}

	var n_NF big.Float
	n_NF.SetPrec(1024)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				n_NF.SetString("1")
				Lch_NF[i][j] <- n_NF
			} else if i < j {
				n_NF.SetString("0")
				Lch_NF[i][j] <- n_NF
				Uch_NF[j][i] <- n_NF
			}
		}
	}

	LUgo_NF(&A, &Lgo_NF, &Ugo_NF, &Lch_NF, &Uch_NF)
	fmt.Println("LUgo_NF:", time.Now().Sub(t))
	Norm(&A, &Lgo_NF, &Ugo_NF)

	t = time.Now()
	var Lch_Str [size][size]chan big.Float
	var Uch_Str [size][size]chan big.Float

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			Lch_Str[i][j] = make(chan big.Float, 1)
			Uch_Str[i][j] = make(chan big.Float, 1)
		}
	}

	var n_Str big.Float
	n_Str.SetPrec(1024)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				n_Str.SetString("1")
				Lch_Str[i][j] <- n_Str
			} else if i < j {
				n_Str.SetString("0")
				Lch_Str[i][j] <- n_Str
				Uch_Str[j][i] <- n_Str
			}
		}
	}

	LUgo_Str(&A, &Lgo_Str, &Ugo_Str, &Lch_Str, &Uch_Str)
	fmt.Println("LUgo_Str:", time.Now().Sub(t))
	Norm(&A, &Lgo_Str, &Ugo_Str)

	//L2, U2の結果の表示方法を指定
	//CalcX(&L2, &U2)
	//Norm(&A, &L2, &U2)

	//PrintM(&L2)
	//PrintM(&U2)
}
