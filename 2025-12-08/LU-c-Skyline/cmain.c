#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
#include "skyline.h"

#include "forkjoin.h"

#define PRINT

int size0 = 300;

// typedef long long GoInt64;
// typedef GoInt64 GoInt;
// void forkjoin(GoInt i, GoInt N);
int main(void) {
	size = size0;
	allocArrays(size);
	//fmt.Println("【スカイライン法】")
    char title[] = "skyline";
#ifdef PRINT
    printf("[%s] size=%d\n", title, size);
    // printf("omp_get_num_threads()=%d\n", omp_get_num_threads());
    // printf("omp_get_num_procs()=%d\n", omp_get_num_procs());
#endif // PRINT
	struct timespec ts_start, ts_stop;
	double t_diff;
#ifdef PRINT
	printf("single thread execution\n");
#endif // PRINT
	srand(0);
	init();
	int N = getN();
	int isk;
	
	reset();
	clock_gettime(CLOCK_REALTIME, &ts_start);
	for (int a = 1; a < N; a++) {
		c = 1
		isk = getIsk(c)
		for (int b = 1; b < N; b++){
			i = c - (isk - b)
			j = c
			if (i == a) {
				//fmt.Println("(i, j)=", i, j)
				call(b, i, j)
			}
			if (b == isk) {
				c += 1
				isk = getIsk(c)
			}
		}
	}
	
	clock_gettime(CLOCK_REALTIME, &ts_stop);
	t_diff = (ts_stop.tv_sec - ts_start.tv_sec) + (ts_stop.tv_nsec - ts_start.tv_nsec) / 1000000000.0;
	printf("[%s] single N=%d time=%f\n", title, N, t_diff);

#ifdef PRINT
    // printMatrix();
	printMatrix3();
#endif

#ifdef PRINT
	printf("multithreaded execution\n");
#endif // PRINT
	srand(0);
	reset();

	clock_gettime(CLOCK_REALTIME, &ts_start);
	for (int a = 1; a < N; a++) {
		c = 1
		isk = getIsk(c)
		forkjoin(a, N,c,isk);
	}
    clock_gettime(CLOCK_REALTIME, &ts_stop);
	t_diff = (ts_stop.tv_sec - ts_start.tv_sec) + (ts_stop.tv_nsec - ts_start.tv_nsec) / 1000000000.0;
    printf("[%s] multi N=%d time=%f\n", title, N, t_diff);

#ifdef PRINT
    // printMatrix();
	printMatrix3();
#endif // PRINT

	return 0;
}
