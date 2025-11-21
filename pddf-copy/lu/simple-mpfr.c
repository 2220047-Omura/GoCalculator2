#include <stdio.h>
#include <mpfr.h>

//void printInterval(__mpfi_struct *b);

#define N 8

int comp(void) {
    int acc = 2000;
    char buf[256];

    mpfr_t hilbert[N][N];
    mpfr_t b[N];
    mpfr_t tmp1;
    mpfr_t tmp2;
    mpfr_t tmp;

    // allocate
    for (int i = 0; i < N; i++) {
        mpfr_init2(b[i], acc);
        for (int j = 0; j < N; j++) {
            mpfr_init2(hilbert[i][j], acc);
        }
    }
    mpfr_init2(tmp, acc);
    mpfr_init2(tmp1, acc);
    mpfr_init2(tmp2, acc);

    // initialize
    mpfr_set_str(b[0], "1", 10, MPFR_RNDN);
    for (int i = 1; i < N; i++) {
        mpfr_set_str(b[i], "0", 10, MPFR_RNDN);
    }
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            mpfr_set_str(tmp1, "1", 10, MPFR_RNDN);
            sprintf(buf, "%d", (i+1)+(j+1)-1);
            mpfr_set_str(tmp2, buf, 10, MPFR_RNDN);
            mpfr_div(hilbert[i][j], tmp1, tmp2, MPFR_RNDN);
        }
    }

    printf("----- Hilbert Matrix -----\n\n");
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
//            printInterval((__mpfi_struct *)&(hilbert[i][j]));
            mpfr_printf ("%.128RNf\n", hilbert[i][j]);
        }
        printf("\n");
    }
    printf("----- b -----\n\n");
    for (int i = 0; i < N; i++) {
//        printInterval((__mpfi_struct *)&(b[i]));
        mpfr_printf ("%.128RNf\n", b[i]);
    }
    printf("\n");

    // lu factorization
    for (int k = 0; k < N; k++) {
        for (int i = k+1; i < N; i++) {
            mpfr_div(hilbert[i][k], hilbert[i][k], hilbert[k][k], MPFR_RNDN);
            for (int j = k+1; j < N; j++) {
                mpfr_mul(tmp, hilbert[i][k], hilbert[k][j], MPFR_RNDN);
                mpfr_sub(hilbert[i][j], hilbert[i][j], tmp, MPFR_RNDN);
            }
        }
    }

    // forward substitution
    for (int i = 1; i < N; i++) {
        for (int j = 0; j <= i - 1; j++) {
            mpfr_mul(tmp, b[j], hilbert[i][j], MPFR_RNDN);
            mpfr_sub(b[i], b[i], tmp, MPFR_RNDN);
        }
    }

    // backward substitution
    for (int i = N-1; i >= 0; i--) {
        for (int j = N-1; j > i; j--) {
            mpfr_mul(tmp, b[j], hilbert[i][j], MPFR_RNDN);
            mpfr_sub(b[i], b[i], tmp, MPFR_RNDN);
        }
        mpfr_div(b[i], b[i], hilbert[i][i], MPFR_RNDN);
    }

    // print results
//    mpfr_exp_t exp;
    for (int i = 0; i < N; i++) {
//        mpfr_get_str(buf, &exp, 10, 15,
//            &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
//        printf("[%sx(%d), ", buf, (int)exp);
//        mpfr_get_str(buf, &exp, 10, 15,
//            &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
//        printf("%sx(%d)]\n", buf, (int)exp);
          mpfr_printf ("%.128RNf\n", b[i]);
    }

    // deallocate
    for (int i = 0; i < N; i++) {
        mpfr_clear(b[i]);
        for (int j = 0; j < N; j++) {
            mpfr_clear(hilbert[i][j]);
        }
    }
    mpfr_clear(tmp1);
    mpfr_clear(tmp2);
    return 0;
}

/*
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
*/

void printtest(void) {
    mpfr_t b[N];
    mpfr_t one;
    mpfr_init2(one, 150);
    mpfr_set_str(one, "1", 10, MPFR_RNDN);
    char buf[2560];
    for (int i = 0; i < N; i++) {
        mpfr_init2(b[i], 150);
        sprintf(buf, "-%d", i+1);
        mpfr_set_str(b[i], buf, 10, MPFR_RNDN);
        mpfr_div(b[i], one, b[i], MPFR_RNDN);
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
//        printInterval((__mpfi_struct *)&(b[i]));
        mpfr_printf ("%.128RNf\n", b[i]);
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
