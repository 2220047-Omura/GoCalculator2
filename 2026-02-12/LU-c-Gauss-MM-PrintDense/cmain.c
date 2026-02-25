#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
#include "gauss.h"

#include "forkjoin.h"

#define PRINT


// typedef long long GoInt64;
// typedef GoInt64 GoInt;
// void forkjoin(GoInt i, GoInt N);
int main(int argc, char **argv) {
	/* ===== 追加：引数チェック ===== */
    if (argc < 2) {
        fprintf(stderr, "usage: %s matrix.mtx\n", argv[0]);
        return 1;
    }

    /* ===== 追加：Matrix Market ファイル名を skyline に渡す ===== */
    setMMFilename(argv[1]);

    char title[] = "gauss";

    init();

    int N = getN();

#ifdef PRINT
    printf("[%s] size=%d\n", title, N);
#endif

    struct timespec ts_start, ts_stop;
    double t_diff;

#ifdef PRINT
    printf("single thread execution\n");
#endif

    srand(0);

    /* skyline 初期化 */
	reset();
	
    //printf("N=%d\n",N);

    //printMatrix3();
	
	clock_gettime(CLOCK_REALTIME, &ts_start);
	for (int k = 0; k < N; k++) {
		for (int i = k + 1; i < N; i++) {
			//fmt.Println(k, i)
			call1(k, i, N);
		}
	}
	clock_gettime(CLOCK_REALTIME, &ts_stop);
	t_diff = (ts_stop.tv_sec - ts_start.tv_sec) + (ts_stop.tv_nsec - ts_start.tv_nsec) / 1000000000.0;
	printf("[%s] single N=%d time=%f\n", title, N, t_diff);

#ifdef PRINT
    // printMatrix();
	//printMatrix3();
	//InfoSub();
	//InfoMul();
	Norm2();
#endif

	printf("\n");

#ifdef PRINT
	printf("multithreaded execution\n");
#endif // PRINT
	srand(0);
	reset();

	clock_gettime(CLOCK_REALTIME, &ts_start);
	for (int i = 0; i < N; i++) {
		forkjoin(i, N);
	}
    clock_gettime(CLOCK_REALTIME, &ts_stop);
	t_diff = (ts_stop.tv_sec - ts_start.tv_sec) + (ts_stop.tv_nsec - ts_start.tv_nsec) / 1000000000.0;
    printf("[%s] multi N=%d time=%f\n", title, N, t_diff);

#ifdef PRINT
    // printMatrix();
	//printMatrix3();
	//InfoSub();
	//InfoMul();
	Norm2();
#endif // PRINT

	return 0;
}
