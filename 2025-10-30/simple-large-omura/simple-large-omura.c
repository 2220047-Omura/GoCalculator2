#include <stdio.h>
#include <mpfi.h>
#include <mpfi_io.h>
#include <stdlib.h>
#include <time.h>

void printInterval(__mpfi_struct *b);

#define N 500
#define idx(i, j) ((i) * N + (j))
#define ptr(p, i, j) (&(p[(i) * N + (j)]))

int comp(void) {
    int acc = 1024;
    char buf[256];
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
    __mpfi_struct *hilbert;
    hilbert = (__mpfi_struct *)malloc(N * N * sizeof(__mpfi_struct));
    mpfi_t b[N];
    mpfi_t tmp1;
    mpfi_t tmp2;
    mpfi_t tmp;

    // allocate
    for (int i = 0; i < N; i++) {
        mpfi_init2(b[i], acc);
        for (int j = 0; j < N; j++) {
            // mpfi_init2(hilbert[i][j], acc);
            mpfi_init2(ptr(hilbert, i, j), acc);
        }
    }
    mpfi_init2(tmp, acc);
    mpfi_init2(tmp1, acc);
    mpfi_init2(tmp2, acc);

    // initialize
    mpfi_set_str(b[0], "1", 10);
    for (int i = 1; i < N; i++) {
        mpfi_set_str(b[i], "0", 10);
    }
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            double r = ((double)rand())/RAND_MAX;
	        mpfr_set_d(a, r, MPFR_RNDN);
	        mpfi_interv_fr(ptr(hilbert,i,j), a, a);
            /*
            mpfi_set_str(tmp1, "1", 10);
            sprintf(buf, "%d", (i+1)+(j+1)-1);
            mpfi_set_str(tmp2, buf, 10);
            // mpfi_div(hilbert[i][j], tmp1, tmp2);
            mpfi_div(ptr(hilbert, i, j), tmp1, tmp2);
            */
        }
    }
/*
    printf("----- Hilbert Matrix -----\n\n");
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            // printInterval((__mpfi_struct *)&(hilbert[i][j]));
            printInterval((__mpfi_struct *)&(hilbert[idx(i, j)]));
            // printInterval((__mpfi_struct *)&(ptr(hilbert, i, j)));
        }
        printf("\n");
    }
    printf("----- b -----\n\n");
    for (int i = 0; i < N; i++) {
        printInterval((__mpfi_struct *)&(b[i]));
    }
    printf("\n");
*/
    struct timespec ts_start, ts_stop;

    clock_gettime(CLOCK_REALTIME, &ts_start);
    // lu factorization
    for (int k = 0; k < N; k++) {
        for (int i = k+1; i < N; i++) {
            // mpfi_div(hilbert[i][k], hilbert[i][k], hilbert[k][k]);
            // for (int j = k+1; j < N; j++) {
            //     mpfi_mul(tmp, hilbert[i][k], hilbert[k][j]);
            //     mpfi_sub(hilbert[i][j], hilbert[i][j], tmp);
            // }
            mpfi_div(ptr(hilbert, i, k), ptr(hilbert, i, k), ptr(hilbert, k, k));
            for (int j = k+1; j < N; j++) {
                mpfi_mul(tmp, ptr(hilbert, i, k), ptr(hilbert, k, j));
                mpfi_sub(ptr(hilbert, i, j), ptr(hilbert, i, j), tmp);
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
            // mpfi_mul(tmp, b[j], hilbert[i][j]);
            mpfi_mul(tmp, b[j], ptr(hilbert, i, j));
            mpfi_sub(b[i], b[i], tmp);
        }
    }

    // backward substitution
    for (int i = N-1; i >= 0; i--) {
        for (int j = N-1; j > i; j--) {
            // mpfi_mul(tmp, b[j], hilbert[i][j]);
            mpfi_mul(tmp, b[j], ptr(hilbert, i, j));
            mpfi_sub(b[i], b[i], tmp);
        }
        // mpfi_div(b[i], b[i], hilbert[i][i]);
        mpfi_div(b[i], b[i], ptr(hilbert, i, i));
    }

    // print results
    mpfr_exp_t exp;
    for (int i = 0; i < N; i++) {
        if (i<3||i>N-3){
            mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
             printf("[%sx(%d), ", buf, (int)exp);
            mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
            printf("%sx(%d)]\n", buf, (int)exp);
        }
    }

    //printInterval(ptr(hilbert,N-1,N-1));
    // deallocate
    for (int i = 0; i < N; i++) {
        mpfi_clear(b[i]);
        for (int j = 0; j < N; j++) {
            // mpfi_clear(hilbert[i][j]);
            // mpfi_clear(hilbert[idx(i, j)]);
            mpfi_clear(ptr(hilbert, i, j));
        }
    }
    mpfi_clear(tmp1);
    mpfi_clear(tmp2);

    return 0;
}

void printInterval(__mpfi_struct *b) {
    char buf[256];
    mpfr_exp_t exp;
    mpfr_get_str(buf, &exp, 10, 15,
        // &((__mpfi_struct *)&(b))->left, MPFR_RNDD);
        &(b->left), MPFR_RNDD);
    printf("[%sx(%d), ", buf, (int)exp);
    mpfr_get_str(buf, &exp, 10, 15,
        &(b->right), MPFR_RNDU);
    printf("%sx(%d)]\n", buf, (int)exp);
}

void printtest(void) {
    mpfi_t b[N];
    mpfi_t one;
    mpfi_init2(one, 150);
    mpfi_set_str(one, "1", 10);
    char buf[2560];
    for (int i = 0; i < N; i++) {
        mpfi_init2(b[i], 150);
        sprintf(buf, "-%d", i+1);
        mpfi_set_str(b[i], buf, 10);
        mpfi_div(b[i], one, b[i]);
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
        printInterval((__mpfi_struct *)&(b[i]));
    }
    for (int i = 0; i < N; i++) {
        mpfi_clear(b[i]);
    }
    mpfi_clear(one);
}


int main(void) {
    printf("Hello, World!\n");
    comp();
    // printtest();
    return 0;
}
