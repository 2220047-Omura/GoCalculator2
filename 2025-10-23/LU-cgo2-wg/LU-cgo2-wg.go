package main

// #cgo CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
//
// #include <stdio.h>
// #include <stdlib.h>
// //#include <time.h>
// #include <mpfi.h>
// #include <mpfi_io.h>
//
// void printInterval(__mpfi_struct *b);
// void comp(void);
//
// #define N 300
// int acc = 1024;
// char buf[256];
//
// mpfi_t hilbert[N][N];
// mpfi_t b[N];
// mpfi_t calc[N][N];
// mpfi_t tmp1;
// mpfi_t tmp2;
// //mpfi_t tmp;
//
// int def(void){
//     return N;
// }
//
// int init(void) {
//	   mpfr_t a;
// 	   mpfr_init2(a, acc);
//
//     // allocate
//     for (int i = 0; i < N; i++) {
//         mpfi_init2(b[i], acc);
//         for (int j = 0; j < N; j++) {
//             mpfi_init2(hilbert[i][j], acc);
//	           mpfi_init2(calc[i][j], acc);
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
//
//	           double r = ((double)rand())/RAND_MAX;
//	           mpfr_set_d(a, r, MPFR_RNDN);
//	           mpfi_interv_fr(hilbert[i][j], a, a);
//
// /*
//             mpfi_set_str(tmp1, "1", 10);
//             sprintf(buf, "%d", (i+1)+(j+1)-1);
//             mpfi_set_str(tmp2, buf, 10);
//             mpfi_div(hilbert[i][j], tmp1, tmp2);
// */
//
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
// void LUfact1(int k, int i){
//     // lu factorization
//     mpfi_div(hilbert[i][k], hilbert[i][k], hilbert[k][k]);
// }
//
// void LUfact2(int k, int i, int j){
// 	   //mpfi_t tmp;
//     //mpfi_init2(tmp, acc);
//     // lu factorization
//     mpfi_mul(calc[i][j], hilbert[i][k], hilbert[k][j]);
//     mpfi_sub(hilbert[i][j], hilbert[i][j], calc[i][j]);
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
import "C"

import (
	"fmt"
	"sync"
	"time"
)

func call1(k int, i int, N int) {

	C.LUfact1(C.int(k), C.int(i))

	for j := k + 1; j < N; j++ {
		call2(k, i, j)
	}
}

func call2(k int, i int, j int) {

	C.LUfact2(C.int(k), C.int(i), C.int(j))
}

func call1WG(k int, i int, N int, wg *sync.WaitGroup) {
	defer wg.Done()
	var wg2 sync.WaitGroup

	C.LUfact1(C.int(k), C.int(i))

	for j := k + 1; j < N; j++ {
		wg2.Add(1)
		go call2WG(k, i, j, &wg2)
	}
	wg2.Wait()
}

func call2WG(k int, i int, j int, wg2 *sync.WaitGroup) {
	defer wg2.Done()

	C.LUfact2(C.int(k), C.int(i), C.int(j))
}

func main() {
	var t time.Time
	var wg sync.WaitGroup

	N := int(C.def())
	C.init()

	//fmt.Println("-----逐次-----")

	t = time.Now()
	for k := 0; k < N; k++ {
		for i := k + 1; i < N; i++ {
			//fmt.Println(k, i)
			call1(k, i, N)
		}
	}
	t2 := time.Now().Sub(t)
	//C.comp()
	fmt.Println("逐次：", t2, "\n")

	//fmt.Println("-----並列-----")
	C.init()

	t = time.Now()
	for k := 0; k < N; k++ {
		for i := k + 1; i < N; i++ {
			//fmt.Println(k, i)
			wg.Add(1)
			go call1WG(k, i, N, &wg)
		}
		wg.Wait()
	}
	t2 = time.Now().Sub(t)
	//C.comp()
	fmt.Println("並列：", t2, "\n")
}
