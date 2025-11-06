#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <mpfr.h>

void printInterval(__mpfr_struct *b);
void comp(void);

#define N 2
int acc = 1024;
char buf[256];

mpfr_t hilbert[N][N];
mpfr_t b[N];
mpfr_t L[N][N];
mpfr_t U[N][N];
/*
mpfr_t SUM[N][N];
mpfr_t MUL;
*/
mpfr_t tmp1;
mpfr_t tmp2;
//mpfr_t tmp;

int def(void){
    return N;
}

int init(void) {
	   mpfr_t a;
	   mpfr_init2(a, acc);

    // allocate
    for (int i = 0; i < N; i++) {
        mpfr_init2(b[i], acc);
        for (int j = 0; j < N; j++) {
            mpfr_init2(hilbert[i][j], acc);
	           //mpfr_init2(SUM[i][j], acc);
	           mpfr_init2(L[i][j], acc);
	           mpfr_init2(U[i][j], acc);
        }
    }
    //mpfr_init2(tmp, acc);
    //mpfr_init2(MUL, acc);
    mpfr_init2(tmp1, acc);
    mpfr_init2(tmp2, acc);

    // initialize
    mpfr_set_str(b[0], "1", 10, MPFR_RNDN);
    for (int i = 1; i < N; i++) {
        mpfr_set_str(b[i], "0", 10, MPFR_RNDN);
        mpfr_set_str(L[i][i], "1", 10, MPFR_RNDN);
    }
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
/*
	           double r = ((double)rand())/RAND_MAX;
	           mpfr_set_d(hilbert[i][j], r, MPFR_RNDN);
*/
            if (i != j){
                mpfr_set_str(L[i][j], "0", 10, MPFR_RNDN);
	           }
            mpfr_set_str(U[i][j], "0", 10, MPFR_RNDN);

            mpfr_set_str(tmp1, "1", 10, MPFR_RNDN);
            sprintf(buf, "%d", (i+1)+(j+1)-1);
            mpfr_set_str(tmp2, buf, 10, MPFR_RNDN);
            mpfr_div(hilbert[i][j], tmp1, tmp2, MPFR_RNDN);
        }
    }


    printf("----- Hilbert Matrix -----\n\n");
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            printInterval((__mpfr_struct *)&(hilbert[i][j]));
        }
        printf("\n");
    }

    printf("----- L -----\n\n");
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            printInterval((__mpfr_struct *)&(L[i][j]));
        }
        printf("\n");
    }

    printf("----- U -----\n\n");
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            printInterval((__mpfr_struct *)&(U[i][j]));
        }
        printf("\n");
    }

    printf("----- b -----\n\n");
    for (int i = 0; i < N; i++) {
        printInterval((__mpfr_struct *)&(b[i]));
    }
    printf("\n");


	   return 0;
}

void Uset(int i,int j){
	   mpfr_t SUM;
	   mpfr_t MUL;
    mpfr_init2(SUM, acc);
    mpfr_init2(MUL, acc);
	   for (int k = 0; k < N; k++) {
	       if (k != i){
	           mpfr_mul(MUL, L[i][k], U[k][j], MPFR_RNDN);
	           mpfr_add(SUM, SUM, MUL, MPFR_RNDN);
	       }
           printM(U[i][j])
	   }
	   mpfr_sub(U[i][j], hilbert[i][j], SUM, MPFR_RNDN);
	   //mpfr_printf("%.128RNf\n",U[i][j]);
}

void Lset(int j,int i){
	   mpfr_t SUM;
	   mpfr_t MUL;
    mpfr_init2(SUM, acc);
    mpfr_init2(MUL, acc);
	   for (int k = 0; k < N; k++) {
	       if (k != i){
	           mpfr_mul(MUL, L[j][k], U[k][i], MPFR_RNDN);
	           mpfr_add(SUM, SUM, MUL, MPFR_RNDN);
	       }
           printM(L[i][j])
	   }
	   mpfr_sub(SUM, hilbert[j][i], SUM, MPFR_RNDN);
	   mpfr_div(L[j][i], SUM, U[i][i], MPFR_RNDN);
	   mpfr_printf("%.128RNf\n",L[j][i]);
}


void comp(void) {
    for (int i = 1; i < N; i++) {
	       mpfr_set(hilbert[i][i], U[i][i], MPFR_RNDN);
        for (int j = i + 1; j < N; j++) {
	           mpfr_set(hilbert[i][j], U[i][j], MPFR_RNDN);
	           mpfr_set(hilbert[j][i], L[j][i], MPFR_RNDN);
        }
    }

	   mpfr_t tmp;
    mpfr_init2(tmp, acc);
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
	   printf("\n");
    mpfr_exp_t exp;
    for (int i = 0; i < N; i++) {
/*
        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
        printf("[%sx(%d), ", buf, (int)exp);
        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
        printf("%sx(%d)]\n", buf, (int)exp);
*/
	       mpfr_printf("%.128RNf\n",b[i]);
    }
	   printf("\n");

    // deallocate
    for (int i = 0; i < N; i++) {
        mpfr_clear(b[i]);
        for (int j = 0; j < N; j++) {
            mpfr_clear(hilbert[i][j]);
        }
    }
    mpfr_clear(tmp1);
    mpfr_clear(tmp2);
}

void printM(mpfr_t *array[a][b]) {
    for (int i =0; i <N; i++){
        for (int j =0; j<N;j++){
            mpfr_printf("%.128RNf\n",*array[i][j]);
        }
        printf("\n");
    }
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
        printInterval((__mpfr_struct *)&(b[i]));
    }
    for (int i = 0; i < N; i++) {
        mpfr_clear(b[i]);
    }
    mpfr_clear(one);
}



