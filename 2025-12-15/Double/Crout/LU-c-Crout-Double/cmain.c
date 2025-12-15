#include <stdio.h>
#include <stdlib.h>
#include <float.h>
#include <time.h>
#include "crout.h"

#include "forkjoin.h"

#define PRINT

//int N0 = 500;

// typedef long long GoInt64;
// typedef GoInt64 GoInt;
// void forkjoin(GoInt i, GoInt N);
int main(void) {
	//N = N0;
	//allocArrays(N);
	//fmt.Println("【クラウト法】")
    char title[] = "crout-v2";
#ifdef PRINT
    printf("[%s] N=%d\n", title, N);
    // printf("omp_get_num_threads()=%d\n", omp_get_num_threads());
    // printf("omp_get_num_procs()=%d\n", omp_get_num_procs());
#endif // PRINT
	struct timespec ts_start, ts_stop;
	double t_diff;
#ifdef PRINT
	printf("single thread execution\n");
#endif // PRINT
	srand(0);
	//init();
	reset();

	clock_gettime(CLOCK_REALTIME, &ts_start);
	for (int i = 0; i < N; i++) {
		for (int j = i; j < N; j++) {
			// Ucall(i, j);
			Uset(i, j);
		}
		for (int j = i + 1; j < N; j++) {
			// Lcall(j, i);
			Lset(j, i);
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
	for (int i = 0; i < N; i++) {
/*
		wg.Add(N - i)
		for j := i; j < N; j++ {
			go UcallWG(i, j)
		}
		wg.Wait()
		wg.Add(N - i - 1)
		for j := i + 1; j < N; j++ {
			go LcallWG(j, i)
		}
		wg.Wait()
*/
		forkjoin(i);
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
