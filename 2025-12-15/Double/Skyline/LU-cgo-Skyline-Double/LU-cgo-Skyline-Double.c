#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
//#include <time.h>
#include <float.h>

#include "SkDouble.h"

//void printInterval(__mpfi_struct *b);
//void comp(void);

//#define size 500
int n;
int E;
int acc = 1024;
char buf[256];
bool boolsk = true;

double *Ask;
int isk[size];
double *Lsk;

double *SUMsk;
double *MULsk;


int getN(void){
    return E;
}

int getIsk(int c){
    return isk[c];
}

/*
void printMatrix(__mpfi_struct *array) {
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            printInterval(ptr(array, i, j));
        }
        printf("\n");
    }
}
*/

int init(void) {
    // define E
    if (boolsk == true){
        for (int i = 1; i < size; i++) {
        n -= 1;
        if (n < 0) {
            n = i;
        }
        E = E + n + 1;
        }
    }else{
        for (int i = 1; i < size; i++) {
        E = E + i + 1;
        }
    }
    //printf("N:%d\n",N);

    // allocate
    Ask = (double *)malloc(E * sizeof(double));
    Lsk = (double *)malloc(E * sizeof(double));
    SUMsk = (double *)malloc(E * sizeof(double));
    MULsk = (double *)malloc(E * sizeof(double));
    
	return 0;
}

void setSkyline(){
    n = 0;
	isk[0] = 0;
	for (int i = 1; i < size; i++) {
		n -= 1;
		if (n < 0) {
			n = i;
		}
		isk[i] = isk[i-1] + n + 1;
	}
    for (int i = 0; i < E; i ++){
        Ask[i] = ((double)rand())/RAND_MAX;
    }
}

void setDense(){
	isk[0] = 0;
	for (int i = 1; i < size; i++) {
		isk[i] = isk[i-1] + i + 1;
	}
    for (int i = 0; i < E; i ++){
        Ask[i] = ((double)rand())/RAND_MAX;
    }
}


void mulDiagonal(){
    int c = 0;
    for (int i = 0; i < E; i++){
        Ask[i] = Ask[i]*100/(isk[c]-i+1);
        if (i == isk[c]) {
            c += 1;
        }
    }
}

void setSkylineTest(){
    int AskTest[10] = {2, 1, 3, 0, 4, 7, 8, 2, 3, 5};
	//N = 10;
	for (int i = 0; i < 10; i++) {
        Ask[i] = AskTest[i];
	}
	int iskTest[5] = {0, 1, 4, 5, 9};
	for (int i = 0; i < 5; i++) {
		isk[i] = iskTest[i];
        //printf("isk : %d\n",isk[i]);
	}
}

void reset(){
    srand(0);
    if (boolsk == true){
        setSkyline();
    }else{
        setDense();
    }
    mulDiagonal();
	//setSkylineTest();
	for (int i = 0; i < E; i++) {
        Lsk[i] = 0;
        SUMsk[i] = 0;
        MULsk[i] = 0;
	}
}

void Usetsk(int b,int i,int j) {
    //printf("Hello from (%d)\n",j);
    int s;
    if (isk[j]-isk[j-1]-(j-i)-1 < isk[i]-isk[i-1]-1) {
		s = isk[j] - isk[j-1] - (j - i) - 1;
	} else {
		s = isk[i] - isk[i-1] - 1;
	}

    for (int k = 0; k < s; k++) {
        Lsk[b] = Ask[isk[i]-(s-k)]/Ask[isk[i-(s-k)]];
        MULsk[b] = Lsk[b] * Ask[isk[j]-(j-i)-(s-k)];
        SUMsk[b] += MULsk[b];
    }
    Ask[b] -= SUMsk[b];
}
/*
void printInterval(__mpfi_struct *b) {

	   for (int i = 0;i < N;i++){
	       mpfi_printf("%.128RNf\n",b[i]);
	   }


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

void allocArrays() {
    
    N = size;
    hilbert = (__mpfi_struct *)calloc(N * N, sizeof(__mpfi_struct));
    b = (__mpfi_struct *)calloc(N, sizeof(__mpfi_struct));
    SUM = (__mpfi_struct *)calloc(N * N, sizeof(__mpfi_struct));
    MUL = (__mpfi_struct *)calloc(N * N, sizeof(__mpfi_struct));
    
    memset(isk,0,sizeof(isk));
    free(Ask);
    free(Lsk);
    free(SUMsk);
    free(MULsk);
}
*/

void printMatrix3(void) {
    for (int i = E-3; i < E; i++) {
        printf("(%d) = ", i);
        printf("%lf ",Ask[i]);
        printf("\n");
    }
}