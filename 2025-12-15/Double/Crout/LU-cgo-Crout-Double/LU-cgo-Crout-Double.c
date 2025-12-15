#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
//#include <time.h>
#include <float.h>
#include "crout.h"

//void printInterval(__double_struct *b);
void comp(void);

// #define N 300
// #define N 500
#define ptr(p, i, j) (&(p[(i) * N + (j)]))
int acc = 1024;
char buf[256];
bool boolsk = false;

double hilbert[N][N];
double b[N];

double SUM[N][N];
double MUL[N][N];

double tmp1;
double tmp2;
//double tmp;

int def(void){
    return N;
}

void printMatrix(double A[N][N]) {
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            printf("%lf ",A[i][j]);
        }
        printf("\n");
    }
}

/*
int init(void) {
    // allocate
    for (int i = 0; i < N; i++) {
        mpfi_init2(b[i], acc);
        for (int j = 0; j < N; j++) {
            mpfi_init2(hilbert[i][j], acc);
	        mpfi_init2(SUM[i][j], acc);
	        mpfi_init2(MUL[i][j], acc);
        }
    }
    //mpfi_init2(tmp, acc);
    //mpfi_init2(MUL, acc);
    mpfi_init2(tmp1, acc);
    mpfi_init2(tmp2, acc);
    
    // initialize
    mpfi_set_str(b[0], "1", 10);
    for (int i = 1; i < N; i++) {
        mpfi_set_str(b[i], "0", 10);
    }
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {

            mpfi_set_str(SUM[i][j], "0", 10);

	           double r = ((double)rand())/RAND_MAX;
	           mpfr_set_d(a, r, MPFR_RNDN);
	           mpfi_interv_fr(hilbert[i][j], a, a);


            mpfi_set_str(tmp1, "1", 10);
            sprintf(buf, "%d", (i+1)+(j+1)-1);
            mpfi_set_str(tmp2, buf, 10);
            mpfi_div(hilbert[i][j], tmp1, tmp2);

        }
    }
    


printf("----- Hilbert Matrix -----\n\n");
printMatrix((__mpfi_struct *)hilbert);




	   return 0;
}
*/

void setSkyline(){
    int c = 0;
    srand(0);
    for (int i = 0; i < N; i++) {
        hilbert[i][i] = ((double)rand())/RAND_MAX;
        if (i-c<0){
            c = 0;
        }
        for (int j = i-1; j >= c; j--) {
	        hilbert[i][j] = ((double)rand())/RAND_MAX;
            hilbert[j][i] = ((double)rand())/RAND_MAX;
        }
        c += 2;
    }
}

void setDense(void){
    srand(0);
    for (int i = 0; i < N; i++) {
        hilbert[i][i] = ((double)rand())/RAND_MAX;
        for (int j = i+1; j < N; j++) {
	        hilbert[i][j] = ((double)rand())/RAND_MAX;
            hilbert[j][i] = ((double)rand())/RAND_MAX;
        }
    }
}

void mulDiagonal(void){
    for (int i = 0; i < N; i++){
        hilbert[i][i] *= 100;
        for (int j= i+1; j < N; j ++){
            hilbert[i][j] = hilbert[i][j] * 100 / (j - i + 1);
            hilbert[j][i] = hilbert[j][i] * 100 / (j - i + 1);
        }
    }
}

void reset(){
    if (boolsk == true){
        setSkyline();
    }else{
        setDense();
    }
    //mulDiagonal();

    // initialize
    b[0] = 1;
    for (int i = 1; i < N; i++) {
        b[i] = 0;
    }
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            SUM[i][j] = 0;
        }
    }
}

void Uset(int i,int j){
	for (int k = 0; k < i; k++) {
        MUL[i][j] = hilbert[i][k] * hilbert[k][j];
        SUM[i][j] = SUM[i][j] + MUL[i][j];
	}
    hilbert[i][j] = hilbert[i][j] - SUM[i][j];
}

void Lset(int j,int i){
	for (int k = 0; k < i; k++) {
        MUL[j][i] = hilbert[j][k] * hilbert[k][i];
        SUM[j][i] = SUM[j][i] + MUL[j][i];
	}
    SUM[j][i] = hilbert[j][i] - SUM[j][i];
    hilbert[j][i] = SUM[j][i] / hilbert[i][i];
    //printMatrix((__mpfi_struct *)L);
}


void comp(void) {
	double tmp;
    // forward substitution
    for (int i = 1; i < N; i++) {
        for (int j = 0; j <= i - 1; j++) {
            tmp = b[j] * hilbert[i][j];
            b[i] -= tmp;
        }
    }

    // backward substitution
    for (int i = N-1; i >= 0; i--) {
        for (int j = N-1; j > i; j--) {
            tmp = b[j] * hilbert[i][j];
            b[i] -= tmp;
        }
        b[i] = b[i]/hilbert[i][i];
    }

    // print results
	printf("\n");
    for (int i = 0; i < N; i++) {
        printf("(%d)[%lf]",i,b[i]);
	    //mpfr_printf("%.128RNf\n",b[i]);
    }
	printf("\n");
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

/*
// void printMatrix3(__mpfi_struct *array) {
void printMatrix3(void) {
    for (int i = N-3; i < N; i++) {
        for (int j = N-3; j < N; j++) {
            printf("(%d, %d) = ", i, j);
            printInterval((__mpfi_struct *)&(hilbert[i][j]));
//             printInterval(ptr(hilbert, i, j));
        }
        printf("\n");
    }
}
*/

void printMatrix3(void) {
    for (int i = N-3; i < N; i++) {
        for (int j = N-3; j < N; j++) {
            printf("(%d, %d) = ", i, j);
            printf("%lf ",hilbert[i][j]);
        }
        printf("\n");
    }
}

