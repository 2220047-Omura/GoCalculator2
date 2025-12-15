#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
//#include <time.h>
#include <float.h>
#include "gauss.h"

//void printInterval(__mpfi_struct *b);
void comp(void);

// #define N 300
// #define N 500
//#define ptr(p, i, j) (&(p[(i) * N + (j)]))
int acc = 1024;
char buf[256];
bool boolsk = false;

double hilbert[N][N];
double b[N];
double calc[N][N];

int def(void){
    return N;
}

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
            calc[i][j] = 0;
        }
    }
}

void LUfact1(int k, int i){
    // lu factorization
    hilbert[i][k] = hilbert[i][k]/hilbert[k][k];
}

void LUfact2(int k, int i, int j){
    // lu factorization
    calc[i][j] = hilbert[i][k]*hilbert[k][j];
    hilbert[i][j] -= calc[i][j];
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

// void printMatrix3(__mpfi_struct *array) {
void printMatrix3(void) {
    for (int i = N-3; i < N; i++) {
        for (int j = N-3; j < N; j++) {
            printf("(%d, %d) = ", i, j);
            printf("%lf ",hilbert[i][j]);
        }
        printf("\n");
    }
}

