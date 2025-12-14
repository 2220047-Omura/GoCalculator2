#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
//#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
//#include <mpfr.h>

#include "SkMpfi.h"

void printInterval(__mpfi_struct *b);
//void comp(void);

//#define size 500
int n;
int E;
int acc = 1024;
char buf[256];
bool boolsk = true;

mpfi_t *Ask;
int isk[size];
mpfi_t *Lsk;

mpfi_t *SUMsk;
mpfi_t *MULsk;

mpfi_t tmp1;
mpfi_t tmp2;
//mpfi_t tmp;

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
	mpfr_t a;
	mpfr_init2(a, acc);

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
    Ask = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    Lsk = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    SUMsk = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    MULsk = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    
    for (int i = 0; i < E; i++) {
        mpfi_init2(Ask[i], acc);
        mpfi_init2(Lsk[i], acc);
        mpfi_init2(SUMsk[i], acc);
        mpfi_init2(MULsk[i], acc);
    }
    //printInterval((__mpfi_struct *)&(Ask[0]));
    //mpfi_init2(tmp, acc);
    //mpfi_init2(MUL, acc);
    mpfi_init2(tmp1, acc);
    mpfi_init2(tmp2, acc);
	return 0;
}

void setSkyline(){
    mpfr_t a;
 	mpfr_init2(a, acc);

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
        double r = ((double)rand())/RAND_MAX;
	    mpfr_set_d(a, r, MPFR_RNDN);
	    mpfi_interv_fr(Ask[i], a, a);
    }
}

void setDense(){
    mpfr_t a;
 	mpfr_init2(a, acc);

	isk[0] = 0;
	for (int i = 1; i < size; i++) {
		isk[i] = isk[i-1] + i + 1;
	}
    for (int i = 0; i < E; i ++){
        double r = ((double)rand())/RAND_MAX;
	    mpfr_set_d(a, r, MPFR_RNDN);
        mpfi_interv_fr(Ask[i], a, a);
    }
}


void mulDiagonal(){
    mpfi_t hdr;
    mpfr_t one;
    mpfr_t toIsk;
    mpfr_t toI;
    mpfi_t div;
    mpfi_init2(hdr,acc);
    mpfr_init2(one,acc);
    mpfr_init2(toIsk,acc);
    mpfr_init2(toI,acc);
    mpfi_init2(div,acc);
    mpfi_set_str(hdr,"100",10);
    mpfr_set_str(one,"1",10,MPFR_RNDN);
    int c = 0;
    for (int i = 0; i < E; i++){
        
        /*
        if (i <3){
            printf("isk[c]-i:%d\n",isk[c]-i);
        }
        */

        mpfr_set_si(toIsk, isk[c], MPFR_RNDN);
        mpfr_set_si(toI, i, MPFR_RNDN);
        mpfr_sub(toIsk, toIsk, toI, MPFR_RNDN);
        mpfr_add(toIsk, toIsk, one, MPFR_RNDN);
        mpfi_interv_fr(div,toIsk,toIsk);
        mpfi_div(div, hdr, div);

        mpfi_mul(Ask[i], Ask[i], div);
        if (i == isk[c]) {
            c += 1;
        }
    }
}

void setSkylineTest(){
    int AskTest[10] = {2, 1, 3, 0, 4, 7, 8, 2, 3, 5};
	//N = 10;
	for (int i = 0; i < 10; i++) {
        mpfi_set_si(Ask[i],AskTest[i]);
        //printInterval((__mpfi_struct *)&(Ask[i]));
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
        mpfi_set_str(Lsk[i], "0", 10);
        mpfi_set_str(SUMsk[i], "0", 10);
        mpfi_set_str(MULsk[i], "0", 10);
	}
}

void Usetsk(int b,int i,int j){
    int s;
    if (isk[j]-isk[j-1]-(j-i)-1 < isk[i]-isk[i-1]-1) {
		s = isk[j] - isk[j-1] - (j - i) - 1;
	} else {
		s = isk[i] - isk[i-1] - 1;
	}

    for (int k = 0; k < s; k++) {
        mpfi_div(Lsk[b],Ask[isk[i]-(s-k)],Ask[isk[i-(s-k)]]);
        mpfi_mul(MULsk[b],Lsk[b],Ask[isk[j]-(j-i)-(s-k)]);
        mpfi_add(SUMsk[b],SUMsk[b],MULsk[b]);
    }
    mpfi_sub(Ask[b], Ask[b], SUMsk[b]);
    //printInterval((__mpfi_struct *)&(Ask[b]));
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

void result(){
    for (int i = E-3; i < E; i++) {
        printInterval((__mpfi_struct *)&(Ask[i]));
    }
}

void clear(){
    //printf("clear");
    memset(isk,0,sizeof(isk));
    free(Ask);
    free(Lsk);
    free(SUMsk);
    free(MULsk);
}

/*
void printM(__mpfi_struct *b) {

}
*/


