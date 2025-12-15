#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <stdbool.h>
#include <float.h>

#include "skyline.h"

//void printInterval(__mpfi_struct *b);

//#define size 500
//int size;
//#define ptr(p, i, j) (&(p[(i) * N + (j)]))
//#define ptr(p, i, j) (&(p[(i) * size + (j)]))
int n;
int E; // Number of Elements
int acc = 1024;
char buf[256];
bool boolsk = true;
/*
#ifdef MDIMARRAY
mpfi_t hilbert[N][N];
mpfi_t b[N];
#else
__mpfi_struct *hilbert;
__mpfi_struct *b;
#endif // MDIMARRAY

#ifdef MDIMARRAY
mpfi_t SUM[N][N];
mpfi_t MUL[N][N];
#else
__mpfi_struct *SUM;
__mpfi_struct *MUL;
#endif // MDIMARRAY
*/

double *Ask;
int isk[size];
double *Lsk;

double *SUMsk;
double *MULsk;

double tmp1;
double tmp2;
//mpfi_t tmp;

int getN(void){
    return E;
}

int getIsk(int c){
    return isk[c];
}
/*
void printMatrix(__mpfi_struct *array) {
    for (int i = E-3; i < E; i++) {
        printf(array[i]);
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

    // allocate
    /*
    for (int i = 0; i < N; i++) {
#ifdef MDIMARRAY
        mpfi_init2(b[i], acc);
#else
        mpfi_init2(&b[i], acc);
#endif // MDIMARRAY
        for (int j = 0; j < N; j++) {
#ifdef MDIMARRAY
            mpfi_init2(hilbert[i][j], acc);
	           mpfi_init2(SUM[i][j], acc);
	           mpfi_init2(MUL[i][j], acc);
#else
            mpfi_init2(ptr(hilbert, i, j), acc);
               mpfi_init2(ptr(SUM, i, j), acc);
               mpfi_init2(ptr(MUL, i, j), acc);
#endif // MDIMARRAY
        }
    }
    */
    Ask = (double *)malloc(E * sizeof(double));
    Lsk = (double *)malloc(E * sizeof(double));
    SUMsk = (double *)malloc(E * sizeof(double));
    MULsk = (double *)malloc(E * sizeof(double));
    /*
    for (int i = 0; i < E; i++) {  
        mpfi_init2(Ask[i], acc);
        mpfi_init2(Lsk[i], acc);
        mpfi_init2(SUMsk[i], acc);
        mpfi_init2(MULsk[i], acc);
    }
    //mpfi_init2(tmp, acc);
    //mpfi_init2(MUL, acc);
    mpfi_init2(tmp1, acc);
    mpfi_init2(tmp2, acc);
    */
    /*
    // initialize
#ifdef MDIMARRAY
    mpfi_set_str(b[0], "1", 10);
#else
    mpfi_set_str(&b[0], "1", 10);
#endif // MDIMARRAY
    for (int i = 1; i < N; i++) {
#ifdef MDIMARRAY
        mpfi_set_str(b[i], "0", 10);
#else
        mpfi_set_str(&b[i], "0", 10);
#endif // MDIMARRAY
    }
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
#ifdef MDIMARRAY
            mpfi_set_str(SUM[i][j], "0", 10);
#else
            mpfi_set_str(ptr(SUM, i, j), "0", 10);
#endif // MDIMARRAY
	           double r = ((double)rand())/RAND_MAX;
	           mpfr_set_d(a, r, MPFR_RNDN);
#ifdef MDIMARRAY
	           mpfi_interv_fr(hilbert[i][j], a, a);
#else
               mpfi_interv_fr(ptr(hilbert, i, j), a, a);
#endif // MDIMARRAY


            mpfi_set_str(tmp1, "1", 10);
            sprintf(buf, "%d", (i+1)+(j+1)-1);
            mpfi_set_str(tmp2, buf, 10);
#ifdef MDIMARRAY
            mpfi_div(hilbert[i][j], tmp1, tmp2);
#else
            mpfi_div(ptr(hilbert, i, j), tmp1, tmp2);
#endif // MDIMARRAY

        }
    }
    */
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
        /*
#ifdef MDIMARRAY
        mpfi_set_si(Ask[i],AskTest[i]);
        //printInterval((__mpfi_struct *)&(Ask[i]));
#else
        mpfi_set_si(&Ask[i],&AskTest[i]);
#endif //MDIMARRAY
*/
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
        /*
#ifdef MDIMARRAY
        mpfi_set_str(Lsk[i], "0", 10);
        mpfi_set_str(SUMsk[i], "0", 10);
        mpfi_set_str(MULsk[i], "0", 10);
#else
        mpfi_set_str(&Lsk[i], "0", 10);
        mpfi_set_str(&SUMsk[i], "0", 10);
        mpfi_set_str(&MULsk[i], "0", 10);
#endif //MDIMARRAY
*/
	}
}
/*
void Uset(int i,int j){

	   mpfi_t SUM;
	   mpfi_t MUL;
    mpfi_init2(SUM, acc);
    mpfi_init2(MUL, acc);
	   mpfi_set_str(SUM,"0",10);
    //mpfi_set_str(MUL,"0",10);

	   for (int k = 0; k < i; k++) {
#ifdef MDIMARRAY
	       mpfi_mul(MUL[i][j], hilbert[i][k], hilbert[k][j]);
	       mpfi_add(SUM[i][j], SUM[i][j], MUL[i][j]);
#else
           mpfi_mul(ptr(MUL, i, j), ptr(hilbert, i, k), ptr(hilbert, k, j));
           mpfi_add(ptr(SUM, i, j), ptr(SUM, i, j), ptr(MUL, i, j));
#endif // MDIMARRAY
	   }
#ifdef MDIMARRAY
	   mpfi_sub(hilbert[i][j], hilbert[i][j], SUM[i][j]);
#else
       mpfi_sub(ptr(hilbert, i, j), ptr(hilbert, i, j), ptr(SUM, i, j));
#endif // MDIMARRAY
    //printMatrix((__mpfi_struct *)U);
}
*/

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