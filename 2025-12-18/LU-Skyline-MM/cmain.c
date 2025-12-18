#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
#include "skyline.h"
#include "forkjoin.h"

#define PRINT

int main(int argc, char **argv) {

    /* ===== 追加：引数チェック ===== */
    if (argc < 2) {
        fprintf(stderr, "usage: %s matrix.mtx\n", argv[0]);
        return 1;
    }

    /* ===== 追加：Matrix Market ファイル名を skyline に渡す ===== */
    setMMFilename(argv[1]);

    char title[] = "skyline";

#ifdef PRINT
    printf("[%s] size=%d\n", title, size);
#endif

    struct timespec ts_start, ts_stop;
    double t_diff;

#ifdef PRINT
    printf("single thread execution\n");
#endif

    srand(0);

    /* skyline 初期化 */
    init();

    int E = getN();
    int isk;
    int i, j, c;

    reset();
	printMatrix3();

    clock_gettime(CLOCK_REALTIME, &ts_start);
    for (int a = 1; a < E; a++) {
        c = 1;
        isk = getIsk(c);
        for (int b = 1; b < E; b++){
            i = c - (isk - b);
            j = c;
            if (i == a) {
                Usetsk(b, i, j);
            }
            if (b == isk) {
                c += 1;
                isk = getIsk(c);
            }
        }
    }

    clock_gettime(CLOCK_REALTIME, &ts_stop);
    t_diff = (ts_stop.tv_sec - ts_start.tv_sec)
           + (ts_stop.tv_nsec - ts_start.tv_nsec) / 1e9;

    printf("[%s] single size=%d time=%f\n", title, size, t_diff);

    printMatrix3();

#ifdef PRINT
    printf("multithreaded execution\n");
#endif

    srand(0);
    reset();

    clock_gettime(CLOCK_REALTIME, &ts_start);
    for (int a = 1; a < E; a++) {
        c = 1;
        isk = getIsk(c);
        forkjoin(a, c, isk);
    }

    clock_gettime(CLOCK_REALTIME, &ts_stop);
    t_diff = (ts_stop.tv_sec - ts_start.tv_sec)
           + (ts_stop.tv_nsec - ts_start.tv_nsec) / 1e9;

    printf("[%s] multi size=%d time=%f\n", title, size, t_diff);

    printMatrix3();

    return 0;
}
