#include <stdio.h>
#include <mpfi.h>
#include <mpfi_io.h>
#include <stdlib.h>
#include <time.h>

void printInterval(__mpfr_struct *b);

#define N 500
#define idx(i, j) ((i) * N + (j))
#define ptr(p, i, j) (&(p[(i) * N + (j)]))

int comp(void) {
    int acc = 1024;
    //char buf[256];
    mpfr_t a;
    mpfr_init2(a, acc);

    // mpfi_t hilbert[N][N];
    // mpfi_t **hilbert;
    // hilbert = (mpfi_t **)malloc(N * sizeof(mpfi_t *));
    // for (int i = 0; i < N; i++) {
    //     hilbert[i] = (mpfi_t *)malloc(N * sizeof(mpfi_t));
    // }
    // mpfi_t *hilbert;
    // hilbert = (mpfi_t *)malloc(N * N * sizeof(mpfi_t));
    __mpfr_struct *hilbert;
    hilbert = (__mpfr_struct *)malloc(N * N * sizeof(__mpfr_struct));
    mpfr_t b[N];
    mpfr_t tmp1;
    mpfr_t tmp2;
    mpfr_t tmp;

    // allocate
    for (int i = 0; i < N; i++) {
        mpfr_init2(b[i], acc);
        for (int j = 0; j < N; j++) {
            // mpfi_init2(hilbert[i][j], acc);
            mpfr_init2(ptr(hilbert, i, j), acc);
        }
    }
    mpfr_init2(tmp, acc);
    mpfr_init2(tmp1, acc);
    mpfr_init2(tmp2, acc);

    // initialize
    mpfr_set_str(b[0], "1", 10,MPFR_RNDN);
    for (int i = 1; i < N; i++) {
        mpfr_set_str(b[i], "0", 10,MPFR_RNDN);
    }
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            double r = ((double)rand())/RAND_MAX;
	        mpfr_set_d(ptr(hilbert, i, j), r, MPFR_RNDN);
	        //mpfi_interv_fr(ptr(hilbert,i,j), a, a);
            /*
            mpfr_set_str(tmp1, "1", 10,MPFR_RNDN);
            sprintf(buf, "%d", (i+1)+(j+1)-1);
            mpfr_set_str(tmp2, buf, 10,MPFR_RNDN);
            // mpfr_div(hilbert[i][j], tmp1, tmp2,MPFR_RNDN);
            mpfr_div(ptr(hilbert, i, j), tmp1, tmp2,MPFR_RNDN);
            */
        }
    }
/*
    printf("----- Hilbert Matrix -----\n\n");
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            // printInterval((__mpfr_struct *)&(hilbert[i][j]));
            printInterval((__mpfr_struct *)&(hilbert[idx(i, j)]));
            // printInterval((__mpfr_struct *)&(ptr(hilbert, i, j)));
        }
        printf("\n");
    }
    printf("----- b -----\n\n");
    for (int i = 0; i < N; i++) {
        printInterval((__mpfr_struct *)&(b[i]));
    }
    printf("\n");
*/
    struct timespec ts_start, ts_stop;

    clock_gettime(CLOCK_REALTIME, &ts_start);
    // lu factorization
    for (int k = 0; k < N; k++) {
        for (int i = k+1; i < N; i++) {
            // mpfr_div(hilbert[i][k], hilbert[i][k], hilbert[k][k],MPFR_RNDN);
            // for (int j = k+1; j < N; j++) {
            //     mpfr_mul(tmp, hilbert[i][k], hilbert[k][j],MPFR_RNDN);
            //     mpfr_sub(hilbert[i][j], hilbert[i][j], tmp,MPFR_RNDN);
            // }
            mpfr_div(ptr(hilbert, i, k), ptr(hilbert, i, k), ptr(hilbert, k, k),MPFR_RNDN);
            for (int j = k+1; j < N; j++) {
                mpfr_mul(tmp, ptr(hilbert, i, k), ptr(hilbert, k, j),MPFR_RNDN);
                mpfr_sub(ptr(hilbert, i, j), ptr(hilbert, i, j), tmp,MPFR_RNDN);
            }
        }
    }
    clock_gettime(CLOCK_REALTIME, &ts_stop);
    double t_diff = (ts_stop.tv_sec * (double)1000000000. + ts_stop.tv_nsec)
             - (ts_start.tv_sec * (double)1000000000. + ts_start.tv_nsec);
    //printf("n_interval = %ld, m = %lf, pi = %lf\n", n_intervals, m, pi);
    printf("total time:    %10lf\n", t_diff / 1000000000.);

    // forward substitution
    for (int i = 1; i < N; i++) {
        for (int j = 0; j <= i - 1; j++) {
            // mpfr_mul(tmp, b[j], hilbert[i][j],MPFR_RNDN);
            mpfr_mul(tmp, b[j], ptr(hilbert, i, j),MPFR_RNDN);
            mpfr_sub(b[i], b[i], tmp,MPFR_RNDN);
        }
    }

    // backward substitution
    for (int i = N-1; i >= 0; i--) {
        for (int j = N-1; j > i; j--) {
            // mpfr_mul(tmp, b[j], hilbert[i][j],MPFR_RNDN);
            mpfr_mul(tmp, b[j], ptr(hilbert, i, j),MPFR_RNDN);
            mpfr_sub(b[i], b[i], tmp,MPFR_RNDN);
        }
        // mpfr_div(b[i], b[i], hilbert[i][i],MPFR_RNDN);
        mpfr_div(b[i], b[i], ptr(hilbert, i, i),MPFR_RNDN);
    }

    // print results
    printf("\n");
    //mpfr_exp_t exp;
    for (int i = 0; i < N; i++) {
        /*
        if (i<3||i>N-3){
            mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
             printf("[%sx(%d), ", buf, (int)exp);
            mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
            printf("%sx(%d)]\n", buf, (int)exp);
        }
            */
            //mpfr_printf("%.128RNf\n",b[i]);
    }
    //printf("\n");

    //printInterval(ptr(hilbert,N-1,N-1));
    // deallocate
    for (int i = 0; i < N; i++) {
        mpfr_clear(b[i]);
        for (int j = 0; j < N; j++) {
            // mpfi_clear(hilbert[i][j]);
            // mpfi_clear(hilbert[idx(i, j)]);
            mpfr_clear(ptr(hilbert, i, j));
        }
    }
    mpfr_clear(tmp1);
    mpfr_clear(tmp2);

    return 0;
}

void printInterval(__mpfr_struct *b) {
    for (int i = 0;i < N;i++){
	    mpfr_printf("%.128RNf\n",b[i]);
 	}
    /*
    char buf[256];
    mpfr_exp_t exp;
    mpfr_get_str(buf, &exp, 10, 15,
        // &((__mpfi_struct *)&(b))->left, MPFR_RNDD);
        &(b->left), MPFR_RNDD);
    printf("[%sx(%d), ", buf, (int)exp);
    mpfr_get_str(buf, &exp, 10, 15,
        &(b->right), MPFR_RNDU);
    printf("%sx(%d)]\n", buf, (int)exp);
    */
}

void printtest(void) {
    mpfr_t b[N];
    mpfr_t one;
    mpfr_init2(one, 150);
    mpfr_set_str(one, "1", 10,MPFR_RNDN);
    char buf[2560];
    for (int i = 0; i < N; i++) {
        mpfr_init2(b[i], 150);
        sprintf(buf, "-%d", i+1);
        mpfr_set_str(b[i], buf, 10,MPFR_RNDN);
        mpfr_div(b[i], one, b[i],MPFR_RNDN);
    }
/*
    mpfr_exp_t exp;
    for (int i = 0; i < N; i++) {
        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
        printf("[%sx(%d), ", buf, (int)exp);
        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
        printf("%sx(%d)]\n", buf, (int)exp);
    }
*/
    for (int i = 0; i < N; i++) {
        printInterval((__mpfr_struct *)&(b[i]));
    }
    for (int i = 0; i < N; i++) {
        mpfr_clear(b[i]);
    }
    mpfr_clear(one);
}


int main(void) {
    printf("Hello, World!\n");
    comp();
    // printtest();
    return 0;
}
