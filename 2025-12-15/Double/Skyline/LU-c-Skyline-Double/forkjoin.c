#include "forkjoin.h"
#include "skyline.h"

// void Ucall(i int, j int) {
// 	C.Uset(C.int(i), C.int(j))
// }

// func Lcall(j int, i int) {
// 	C.Lset(C.int(j), C.int(i))
// }

/*
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
*/

void forkjoin(int a, int c, int isk) {
	int i;
	int j;
#pragma omp parallel
	{
#pragma omp single
    {
	for (int b = 1; b < E; b++){
		i = c - (isk - b);
		j = c;
		if (i == a) {
#pragma omp task firstprivate(b, i, j)
            {
                Usetsk(b, i, j);
            }
		}
		if (b == isk) {
			c += 1;
			isk = getIsk(c);
		}
	}
#pragma omp taskwait
	}
    }
}