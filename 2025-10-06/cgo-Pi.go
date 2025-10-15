package main

//#cgo CFLAGS: -I/opt/homebrew/include
//#cgo LDFLAGS: -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
//
//#include <stdio.h>
//#include <stdlib.h>
//#include <sys/time.h>
//#include <unistd.h>
//#include <mpfi.h>
//#include <mpfi_io.h>
//
//
//void printInterval(__mpfi_struct *b) {
//    char buf[256];
//    mpfr_exp_t exp;
//    mpfr_get_str(buf, &exp, 10, 15,
//        // &((__mpfi_struct *)&(b))->left, MPFR_RNDD);
//        &(b->left), MPFR_RNDD);
//    printf("[%sx(%d), ", buf, (int)exp);
//    mpfr_get_str(buf, &exp, 10, 15,
//        &(b->right), MPFR_RNDU);
//    printf("%sx(%d)]\n", buf, (int)exp);
//}
//
//int MC(void) {
//	  //unsigned int now = (unsigned int)time(NULL);
//	  //srand((unsigned int)time(NULL)*getpid());
//	  struct timeval t1;
//	  gettimeofday(&t1,NULL);
//	  srand(t1.tv_usec * t1.tv_sec);
//	  int randx = rand();
//	  //srand((unsigned int)time(NULL)*getpid());
//	  int randy = rand();
//
//    int acc = 2000;
//    char buf[256];
//
//    mpfi_t x;
//    mpfi_t y;
//    mpfi_t M;
//    mpfi_init2(x, acc);
//    mpfi_init2(y, acc);
//    mpfi_init2(M, acc);
//
//	  sprintf(buf, "%d", RAND_MAX);
//	  mpfi_set_str(M, buf, 10);
//	  sprintf(buf, "%d", randx);
//	  mpfi_set_str(x, buf, 10);
//	  sprintf(buf, "%d", randy);
//	  mpfi_set_str(y, buf, 10);
//
//
//	  mpfi_div(x, x, M);
//	  mpfi_mul(x, x, x);
//
//	  mpfi_div(y, y, M);
//	  mpfi_mul(y, y, y);
//
//	  mpfi_add(x, x, y);
//
//    //printInterval((__mpfi_struct *)&(x));
//
//    mpfr_exp_t exp;
//    mpfr_get_str(buf, &exp, 10, 15,
//        // &((__mpfi_struct *)&(x))->left, MPFR_RNDD);
//        &(x->right), MPFR_RNDU);
//
//	  if ((int)exp == 0){
//	      //printf("x<1\n");
//	      return 1;
//	  }else{
//	      return 0;
//	  }
//
//}
//
//
//
import "C"
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func echo1(c *chan int) {
	result := C.MC()
	*c <- int(result)
}

func echo2(c *chan int) {
	defer wg.Done()
	result := C.MC()
	*c <- int(result)
}

func main() {
	var plot int = 10000
	var sum1, sum2 int
	var c chan int
	c = make(chan int, 1)

	for i := 0; i < plot; i++ {
		echo1(&c)
		sum1 += <-c
	}
	var ans float64 = float64(sum1) / float64(plot) * 4
	fmt.Println("Ans1: ", ans)

	for i := 0; i < plot; i++ {
		wg.Add(1)
		go echo2(&c)
		sum2 += <-c
	}
	wg.Wait()
	ans = float64(sum2) / float64(plot) * 4
	fmt.Println("Ans2: ", ans)
}
