#ifndef _FORKJOIN_
#define _FORKJOIN_

#include "crout.h"

// void Ucall(i int, j int) {
// 	C.Uset(C.int(i), C.int(j))
// }

// func Lcall(j int, i int) {
// 	C.Lset(C.int(j), C.int(i))
// }

void UcallWG(int i, int j);
void LcallWG(int j, int i);

void forkjoin(int i, int N);

#endif // _FORKJOIN_
