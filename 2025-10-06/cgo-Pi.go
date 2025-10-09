package main

//#cgo CFLAGS: -I/opt/homebrew/include
//#cgo LDFLAGS: -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
//
//#include <stdio.h>
//#include <stdlib.h>
//#include <mpfi.h>
//#include <mpfi_io.h>
//
//
//int MC(void) {
//	  int randx = rand();
//	  int randy = rand();
//
//    int acc = 2000;
//    char buf[256];
//
//    mpfi_t x1;
//    mpfi_t y1;
//    mpfi_t x2;
//    mpfi_t y2;
//    mpfi_t M;
//    mpfi_init2(x1, acc);
//    mpfi_init2(y1, acc);
//    mpfi_init2(x2, acc);
//    mpfi_init2(y2, acc);
//    mpfi_init2(M, acc);
//
//	  //mpfi_set_str(x1, "%d", randx);
//
//
//	  //mpfi_div(x2, x1, RAND_MAX);
//	  //mpfi_mul(x2, x2, x2);
//
//	  //mpfi_div(y2, y1, RAND_MAX);
//	  //mpfi_mul(y2, y2, y2);
//
//	  //mpfi_add(x2, x2, y2);
//
//	  if(x2 >= 1){
//	      printf("x2 >= 1");
//	  }
//return 0;
//}
//
//
//
import "C"

func main() {
	C.MC()
}
