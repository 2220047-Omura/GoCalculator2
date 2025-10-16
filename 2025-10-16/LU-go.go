package main

// #cgo CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
//
// #include <stdio.h>
// #include <mpfi.h>
// #include <mpfi_io.h>
//
// void printInterval(__mpfi_struct *b);
// void comp(void);
//
// #define N 8
// int acc = 2000;
// char buf[256];
//
// mpfi_t hilbert[N][N];
// mpfi_t b[N];
// mpfi_t tmp1;
// mpfi_t tmp2;
// //mpfi_t tmp;
//
// int def(void){
//     return N;
// }
//
// int init(void) {
//
//     // allocate
//     for (int i = 0; i < N; i++) {
//         mpfi_init2(b[i], acc);
//         for (int j = 0; j < N; j++) {
//             mpfi_init2(hilbert[i][j], acc);
//         }
//     }
//     //mpfi_init2(tmp, acc);
//     mpfi_init2(tmp1, acc);
//     mpfi_init2(tmp2, acc);
//
//     // initialize
//     mpfi_set_str(b[0], "1", 10);
//     for (int i = 1; i < N; i++) {
//         mpfi_set_str(b[i], "0", 10);
//     }
//     for (int i = 0; i < N; i++) {
//         for (int j = 0; j < N; j++) {
//             mpfi_set_str(tmp1, "1", 10);
//             sprintf(buf, "%d", (i+1)+(j+1)-1);
//             mpfi_set_str(tmp2, buf, 10);
//             mpfi_div(hilbert[i][j], tmp1, tmp2);
//         }
//     }
//
// /*
//     printf("----- Hilbert Matrix -----\n\n");
//     for (int i = 0; i < N; i++) {
//         for (int j = 0; j < N; j++) {
//             printInterval((__mpfi_struct *)&(hilbert[i][j]));
//         }
//         printf("\n");
//     }
//     printf("----- b -----\n\n");
//     for (int i = 0; i < N; i++) {
//         printInterval((__mpfi_struct *)&(b[i]));
//     }
//     printf("\n");
// */
//
//	   return 0;
// }
//
//
// int LUfact1(int k, int i){
//     // lu factorization
//     mpfi_div(hilbert[i][k], hilbert[i][k], hilbert[k][k]);
//	   return 1;
// }
//
// int LUfact2(int k, int i, int j){
// 	   mpfi_t tmp;
//     mpfi_init2(tmp, acc);
//     // lu factorization
//     mpfi_mul(tmp, hilbert[i][k], hilbert[k][j]);
//     mpfi_sub(hilbert[i][j], hilbert[i][j], tmp);
//	   return 1;
//}
//
// void comp(void) {
// 	   mpfi_t tmp;
//     mpfi_init2(tmp, acc);
//     // forward substitution
//     for (int i = 1; i < N; i++) {
//         for (int j = 0; j <= i - 1; j++) {
//             mpfi_mul(tmp, b[j], hilbert[i][j]);
//             mpfi_sub(b[i], b[i], tmp);
//         }
//     }
//
//     // backward substitution
//     for (int i = N-1; i >= 0; i--) {
//         for (int j = N-1; j > i; j--) {
//             mpfi_mul(tmp, b[j], hilbert[i][j]);
//             mpfi_sub(b[i], b[i], tmp);
//         }
//         mpfi_div(b[i], b[i], hilbert[i][i]);
//     }
//
//     // print results
//	   printf("\n");
//     mpfr_exp_t exp;
//     for (int i = 0; i < N; i++) {
//         mpfr_get_str(buf, &exp, 10, 15,
//             &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
//         printf("[%sx(%d), ", buf, (int)exp);
//         mpfr_get_str(buf, &exp, 10, 15,
//             &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
//         printf("%sx(%d)]\n", buf, (int)exp);
//     }
//	   printf("\n");
//
//     // deallocate
//     for (int i = 0; i < N; i++) {
//         mpfi_clear(b[i]);
//         for (int j = 0; j < N; j++) {
//             mpfi_clear(hilbert[i][j]);
//         }
//     }
//     mpfi_clear(tmp1);
//     mpfi_clear(tmp2);
// }
//
// void printInterval(__mpfi_struct *b) {
//     char buf[256];
//     mpfr_exp_t exp;
//     mpfr_get_str(buf, &exp, 10, 15,
//         // &((__mpfi_struct *)&(b))->left, MPFR_RNDD);
//         &(b->left), MPFR_RNDD);
//     printf("[%sx(%d), ", buf, (int)exp);
//     mpfr_get_str(buf, &exp, 10, 15,
//         &(b->right), MPFR_RNDU);
//     printf("%sx(%d)]\n", buf, (int)exp);
// }
//
// void printtest(void) {
//     mpfi_t b[N];
//     mpfi_t one;
//     mpfi_init2(one, 150);
//     mpfi_set_str(one, "1", 10);
//     char buf[2560];
//     for (int i = 0; i < N; i++) {
//         mpfi_init2(b[i], 150);
//         sprintf(buf, "-%d", i+1);
//         mpfi_set_str(b[i], buf, 10);
//         mpfi_div(b[i], one, b[i]);
//     }
// /*
//     mpfr_exp_t exp;
//     for (int i = 0; i < N; i++) {
//         mpfr_get_str(buf, &exp, 10, 15,
//             &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
//         printf("[%sx(%d), ", buf, (int)exp);
//         mpfr_get_str(buf, &exp, 10, 15,
//             &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
//         printf("%sx(%d)]\n", buf, (int)exp);
//     }
// */
//
//     for (int i = 0; i < N; i++) {
//         printInterval((__mpfi_struct *)&(b[i]));
//     }
//     for (int i = 0; i < N; i++) {
//         mpfi_clear(b[i]);
//     }
//     mpfi_clear(one);
// }
//
//
//
//
//import "C"
import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

var wg sync.WaitGroup
var A [size][size]big.Float
var B [size]big.Float

const size = 8

func initialize() {
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

	for i := 0; i < size; i++ {
		if i == 0 {
			B[i].SetPrec(1024).SetString("1")
		} else {
			B[i].SetPrec(1024).SetString("0")
		}
	}
}

func LUfact1(k int, i int) int {
	A[i][k].SetPrec(1024).Mul(&A[i][k], &A[k][k])
	return 1
}

func LUfact2(k int, i int, j int) int {
	var tmp big.Float
	tmp.SetPrec(1024).Mul(&A[i][k], &A[k][j])
	A[i][j].SetPrec(1024).Sub(&A[i][j], &tmp)
	return 1
}

func call1(k int, i int) {
	c := make(chan int, 1)
	c <- int(LUfact1(k, i))
}

func call2(k int, i int, j int) {
	c := make(chan int, 1)
	c <- int(LUfact2(k, i, j))
}

func call1WG(k int, i int) {
	defer wg.Done()
	c := make(chan int, 1)
	c <- int(LUfact1(k, i))
}

func call2WG(k int, i int, j int) {
	defer wg.Done()
	c := make(chan int, 1)
	c <- int(LUfact2(k, i, j))
}

func comp() {
	var tmp big.Float
	tmp.SetPrec(1024)

	// forward substitution
	for i := 1; i < size; i++ {
		for j := 0; j <= i-1; j++ {
			tmp.Mul(&B[j], &A[i][j])
			B[i].Sub(&B[i], &tmp)
		}
	}

	// backward substitution
	for i := size - 1; i >= 0; i-- {
		for j := size - 1; j > i; j-- {
			tmp.Mul(&B[j], &A[i][j])
			B[i].Sub(&B[i], &tmp)
		}
		B[i].Quo(&B[i], &A[i][i])
	}

	for i := 0; i < size; i++ {
		fmt.Println(&B[i])
	}
}

func main() {
	var t time.Time
	//var A [size][size]big.Float
	//var B [size]big.Float
	initialize()

	fmt.Println("-----逐次-----")

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i++ {
			//fmt.Println(k, i)
			call1(k, i)
		}
		for i := k + 1; i < size; i++ {
			for j := k + 1; j < size; j++ {
				//fmt.Println(k, i, j)
				call2(k, i, j)
			}
		}
	}
	t2 := time.Now().Sub(t)
	comp()
	fmt.Println("逐次：", t2, "\n")

	fmt.Println("-----並列-----")
	initialize()

	t = time.Now()
	for k := 0; k < size; k++ {
		for i := k + 1; i < size; i++ {
			//fmt.Println(k, i)
			wg.Add(1)
			go call1WG(k, i)
		}
		wg.Wait()
		for i := k + 1; i < size; i++ {
			for j := k + 1; j < size; j++ {
				//fmt.Println(k, i, j)
				wg.Add(1)
				go call2WG(k, i, j)
			}
		}
		wg.Wait()
	}
	t2 = time.Now().Sub(t)
	comp()
	fmt.Println("並列：", t2, "\n")
}
