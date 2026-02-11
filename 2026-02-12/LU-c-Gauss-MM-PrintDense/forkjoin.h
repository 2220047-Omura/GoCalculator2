#ifndef _FORKJOIN_
#define _FORKJOIN_

#include "gauss.h"

// void Ucall(i int, j int) {
// 	C.Uset(C.int(i), C.int(j))
// }

// func Lcall(j int, i int) {
// 	C.Lset(C.int(j), C.int(i))
// }

void call1(int k, int i, int N);

void call2(int k, int i, int j);

void forkjoin(int k, int N);

#endif // _FORKJOIN_
