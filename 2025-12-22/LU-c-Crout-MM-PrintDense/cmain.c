#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
#include "crout.h"

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

    char title[] = "skyline";

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
    printf("N=%d\n",N);

    reset();

    printMatrix3();
	
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
	//printMatrix3();
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
		forkjoin(i, N);
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
