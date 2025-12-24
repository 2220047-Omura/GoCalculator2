#include "forkjoin.h"
#include "crout.h"

// void Ucall(i int, j int) {
// 	C.Uset(C.int(i), C.int(j))
// }

// func Lcall(j int, i int) {
// 	C.Lset(C.int(j), C.int(i))
// }

void UcallWG(int i, int j) {
	Uset(i, j);
}

void LcallWG(int j, int i) {
	Lset(j, i);
}

void forkjoin(int i, int N) {
#pragma omp parallel
	{
#pragma omp for
	for (int j = i; j < N; j++) {
		UcallWG(i, j);
	}
#pragma omp for
	for (int j = i + 1; j < N; j++) {
		LcallWG(j, i);
	}
	}
}