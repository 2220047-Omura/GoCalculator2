#include "forkjoin.h"
#include "gauss.h"

// void Ucall(i int, j int) {
// 	C.Uset(C.int(i), C.int(j))
// }

// func Lcall(j int, i int) {
// 	C.Lset(C.int(j), C.int(i))
// }

void call1(int k, int i, int N) {
	LUfact1(k, i);
	for (int j = k + 1; j < N; j++) {
		call2(k, i, j);
	}
}

void call2(int k, int i, int j) {
	LUfact2(k, i, j);
}

void forkjoin(int k, int N) {
#pragma omp parallel for
	for (int i = k + 1; i < N; i++) {
		call1(k, i, N);
	}
}
/*
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
*/