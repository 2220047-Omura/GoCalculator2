#include <stdio.h>
#include <mpfi.h>
#include <mpfi_io.h>

void printInterval(__mpfi_struct *b);

#define N 8

int comp(void) {
    int acc = 2000;
    char buf[256];

    mpfi_t hilbert[N][N];
    mpfi_t b[N];
    mpfi_t tmp1;
    mpfi_t tmp2;
    mpfi_t tmp;

    // allocate
    for (int i = 0; i < N; i++) {
        mpfi_init2(b[i], acc);
        for (int j = 0; j < N; j++) {
            mpfi_init2(hilbert[i][j], acc);
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
            mpfi_set_str(tmp1, "1", 10);
            sprintf(buf, "%d", (i+1)+(j+1)-1);
            mpfi_set_str(tmp2, buf, 10);
            mpfi_div(hilbert[i][j], tmp1, tmp2);
        }
    }

    printf("----- Hilbert Matrix -----\n\n");
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            printInterval((__mpfi_struct *)&(hilbert[i][j]));
        }
        printf("\n");
    }
    printf("----- b -----\n\n");
    for (int i = 0; i < N; i++) {
        printInterval((__mpfi_struct *)&(b[i]));
    }
    printf("\n");

    // lu factorization
    for (int k = 0; k < N; k++) {
        for (int i = k+1; i < N; i++) {
            mpfi_div(hilbert[i][k], hilbert[i][k], hilbert[k][k]);
            for (int j = k+1; j < N; j++) {
                mpfi_mul(tmp, hilbert[i][k], hilbert[k][j]);
                mpfi_sub(hilbert[i][j], hilbert[i][j], tmp);
            }
        }
    }

    // forward substitution
    for (int i = 1; i < N; i++) {
        for (int j = 0; j <= i - 1; j++) {
            mpfi_mul(tmp, b[j], hilbert[i][j]);
            mpfi_sub(b[i], b[i], tmp);
        }
    }

    // backward substitution
    for (int i = N-1; i >= 0; i--) {
        for (int j = N-1; j > i; j--) {
            mpfi_mul(tmp, b[j], hilbert[i][j]);
            mpfi_sub(b[i], b[i], tmp);
        }
        mpfi_div(b[i], b[i], hilbert[i][i]);
    }

    // print results
    mpfr_exp_t exp;
    for (int i = 0; i < N; i++) {
        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
        printf("[%sx(%d), ", buf, (int)exp);
        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
        printf("%sx(%d)]\n", buf, (int)exp);
    }

    // deallocate
    for (int i = 0; i < N; i++) {
        mpfi_clear(b[i]);
        for (int j = 0; j < N; j++) {
            mpfi_clear(hilbert[i][j]);
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
