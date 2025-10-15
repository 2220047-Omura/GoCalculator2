package main

//#cgo CFLAGS: -I/opt/homebrew/include
//#cgo LDFLAGS: -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
//
//#include <stdio.h>
//#include <mpfi.h>
//
//void printInterval(__mpfi_struct *b);
//
//int add(int x, int y) {
//    return x + y;
//}
//int mul(int x, int y) {
//    return x * y;
//}
//void comp() {
//    int a = 5;
//    int b = 6;
//    printf("Hello, World! : %d\n", mul(add(a, b), add(a, b)));
//    mpfi_t ma, mb, mc, md;
//    mpfi_init2(ma, 2000);
//    mpfi_init2(mb, 2000);
//    mpfi_init2(mc, 2000);
//    mpfi_init2(md, 2000);
//    mpfi_set_str(ma, "5", 10);
//    mpfi_set_str(mb, "6", 10);
//    mpfi_add(mc, ma, mb);
//    mpfi_mul(md, mc, mc);
//    printInterval((__mpfi_struct *)&(md));
//}
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
import "C"

func main() {
	C.comp()
}
