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
    // printSquare();
#endif

    printf("single thread execution\n");
    srand(0);

    /* skyline 初期化 */
    init();

    reset();

    int E = getN();
    printf("E=%d\n",E);
    int l;
    int E2;

	//printMatrix3();

    
    clock_gettime(CLOCK_REALTIME, &ts_start);
    for (int a = 1; a < size; a++) {
        l = Dia[a];
        E2 = (size <= a + MAXp) ? E : Dia[a + MAXp];
        //forkjoin(a, l, E2);
        //printf("a, l = %d, %d\n",a,l);
        for (int m = l; m < E2; m++){
		    if (isk[m] == a){
                //printf("m, l = %d, %d\n",m,l);
			    Usetsk(m, l);
		    }
	    }
    }

    clock_gettime(CLOCK_REALTIME, &ts_stop);
    t_diff = (ts_stop.tv_sec - ts_start.tv_sec)
           + (ts_stop.tv_nsec - ts_start.tv_nsec) / 1e9;

    printf("[%s] single size=%d time=%f\n", title, size, t_diff);
    
    // printMatrix3();

    //Norm();

#ifdef PRINT
    //printSquare();
    //InfoAdd();
    //InfoMul();
    Norm2();
#endif

    printf("\n");

    printf("multithreaded execution\n");
    srand(0);
    reset();

    clock_gettime(CLOCK_REALTIME, &ts_start);
    for (int a = 1; a < size; a++) {
        l = Dia[a];
        E2 = (size <= a + MAXp) ? E : Dia[a + MAXp];
        forkjoin(a, l, E2);
    }

    clock_gettime(CLOCK_REALTIME, &ts_stop);
    t_diff = (ts_stop.tv_sec - ts_start.tv_sec)
           + (ts_stop.tv_nsec - ts_start.tv_nsec) / 1e9;

    printf("[%s] multi size=%d time=%f\n", title, size, t_diff);

    // printMatrix3();

    //Norm();

#ifdef PRINT
    //printSquare();
    //InfoAdd();
    //InfoMul();
    Norm2();
#endif

    freeArrays();

    return 0;
}
