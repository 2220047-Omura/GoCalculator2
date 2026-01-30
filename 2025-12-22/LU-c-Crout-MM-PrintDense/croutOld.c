#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
//#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
//#include <mpfr.h>

#include "crout.h"

static const char *mm_filename = NULL;

//void printInterval(__mpfi_struct *b);
//void comp(void);

// #define N 300
int N;
#define ptr(p, i, j) (&(p[(i) * N + (j)]))
int acc = 1024;
char buf[256];
bool boolsk = true;

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

mpfi_t tmp1;
mpfi_t tmp2;
//mpfi_t tmp;

void setMMFilename(const char *fname) {
    mm_filename = fname;
}

int def(void){
    return N;
}

int getN() {
    FILE *fp;
    int rows, cols, nnz;
    char line[256];

    if (!mm_filename) {
        fprintf(stderr, "MM filename is not set\n");
        exit(1);
    }

    fp = fopen(mm_filename, "r");
    if (!fp) {
        perror("fopen");
        exit(1);
    }

    /* コメントスキップ */
    do {
        fgets(line, sizeof(line), fp);
    } while (line[0] == '%');

    sscanf(line, "%d %d %d", &rows, &cols, &nnz);
    //printf("%d %d %d\n", rows, cols, nnz);
    N = cols;

    fclose(fp);
    return N;
}

void printMatrix(__mpfi_struct *array) {
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            printInterval(ptr(array, i, j));
        }
        printf("\n");
    }
}

int init(void) {
    getN();
    printf("N = %d\n",N);
	mpfr_t a;
	mpfr_init2(a, acc);
    srand(0);

#ifdef MDIMARRAY
    //mpfi_init2(b[i], acc);
#else
    hilbert = (__mpfi_struct *)malloc(N * N * sizeof(__mpfi_struct));
    b = (__mpfi_struct *)malloc(N * sizeof(__mpfi_struct));
    SUM = (__mpfi_struct *)malloc(N * N * sizeof(__mpfi_struct));
    MUL = (__mpfi_struct *)malloc(N * N * sizeof(__mpfi_struct));
#endif // MDIMARRAY
    // allocate
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
    //mpfi_init2(tmp, acc);
    //mpfi_init2(MUL, acc);
    mpfi_init2(tmp1, acc);
    mpfi_init2(tmp2, acc);
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
#ifdef MDIMARRAY
        mpfi_set_str(hilbert[i][j], "0", 10);
#else
        mpfi_set_str(ptr(hilbert,i,j), "0", 10);
#endif // MDIMARRAY
        }
    }
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


printf("----- Hilbert Matrix -----\n\n");
printMatrix((__mpfi_struct *)hilbert);

*/
	return 0;
}

void setMM() {
    mpfr_t a;
    mpfr_init2(a,acc);
    FILE *fp;
    int rows, cols, nnz;
    char line[256];

    if (!mm_filename) {
        fprintf(stderr, "MM filename is not set\n");
        exit(1);
    }
    fp = fopen(mm_filename, "r");

    /* ヘッダ・コメント行をスキップ */
    do {
        if (!fgets(line, sizeof(line), fp)) {
            fprintf(stderr, "不正なファイル形式です\n");
            fclose(fp);
        }
    } while (line[0] == '%');

    /* coordinate 形式: 行数 列数 非ゼロ要素数 */
    if (sscanf(line, "%d %d %d", &rows, &cols, &nnz) != 3) {
        fprintf(stderr, "coordinate 形式ではありません\n");
        fclose(fp);
    }

    /* 列ごとに処理 */
    //printf("in setMM\n");
    for (int col = 1; col <= cols; col++) {
        rewind(fp);

    /* 再びヘッダをスキップ */
        do {
            fgets(line, sizeof(line), fp);
        } while (line[0] == '%');
        fgets(line, sizeof(line), fp); /* サイズ行 */

        for (int t = 0; t < N*N; t++) {
            int i, j;
            double val;
            fscanf(fp, "%d %d %lf", &i, &j, &val);
            if (i <= j) {
                mpfr_set_d(a, val, MPFR_RNDN);
#ifdef MDIMARRAY
	            mpfi_interv_fr(hilbert[i-1][j-1], a, a);
                mpfi_interv_fr(hilbert[j-1][i-1], a, a);
#else
                mpfi_interv_fr(ptr(hilbert, i-1, j-1), a, a);
                mpfi_interv_fr(ptr(hilbert, j-1, i-1), a, a);
#endif // MDIMARRAY
            }
        }
        //printf("Dia[%d] = %d\n", col-1,k-1);
        //free(a);
    }
    fclose(fp);

    /* 確認用出力 
    printf("Ask (column-major, row-sorted):\n");
    for (int i = 0; i < k; i++) {
       printf("(i, j, Ask) = %d\n",isk[i]);
        //printf("(i, j) = %d %d",isk2[i], jsk[i]);
        //printInterval((__mpfi_struct *)&(Ask[i]));
    }
    */
}

void setSkyline(){
    int c = 0;
    mpfr_t a;
	mpfr_init2(a, acc);
    srand(0);
    for (int i = 0; i < N; i++) {
        double r = ((double)rand())/RAND_MAX;
	    mpfr_set_d(a, r, MPFR_RNDN);
#ifdef MDIMARRAY
	    mpfi_interv_fr(hilbert[i][i], a, a);
#else
        mpfi_interv_fr(ptr(hilbert, i, i), a, a);
#endif // MDIMARRAY
        if (i-c<0){
            c = 0;
        }
        for (int j = i-1; j >= c; j--) {
	        double r = ((double)rand())/RAND_MAX;
	        mpfr_set_d(a, r, MPFR_RNDN);
#ifdef MDIMARRAY
	        mpfi_interv_fr(hilbert[i][j], a, a);
            mpfi_interv_fr(hilbert[j][i], a, a);
#else
            mpfi_interv_fr(ptr(hilbert, i, j), a, a);
            mpfi_interv_fr(ptr(hilbert, j, i), a, a);
#endif // MDIMARRAY
        }
        c += 2;
    }
}


void setDense(void){
    mpfr_t a;
	mpfr_init2(a, acc);
    srand(0);
    for (int i = 0; i < N; i++) {
        double r = ((double)rand())/RAND_MAX;
	    mpfr_set_d(a, r, MPFR_RNDN);
#ifdef MDIMARRAY
	    mpfi_interv_fr(hilbert[i][i], a, a);
#else
        mpfi_interv_fr(ptr(hilbert, i, i), a, a);
#endif // MDIMARRAY
        for (int j = i+1; j < N; j++) {
	           double r = ((double)rand())/RAND_MAX;
	           mpfr_set_d(a, r, MPFR_RNDN);
#ifdef MDIMARRAY
	            mpfi_interv_fr(hilbert[i][j], a, a);
                mpfi_interv_fr(hilbert[j][i], a, a);
#else
                mpfi_interv_fr(ptr(hilbert, i, j), a, a);
                mpfi_interv_fr(ptr(hilbert, j, i), a, a);
#endif // MDIMARRAY
        }
    }
}

void mulDiagonal(void){
    mpfi_t hdr;
    mpfr_t one;
    mpfr_t toI;
    mpfr_t toJ;
    mpfi_t div;
    mpfi_init2(hdr,acc);
    mpfr_init2(one,acc);
    mpfr_init2(toI,acc);
    mpfr_init2(toJ,acc);
    mpfi_init2(div,acc);
    mpfi_set_str(hdr,"100",10);
    mpfr_set_str(one,"1",10,MPFR_RNDN);
    //int c = 0;
    for (int i = 0; i < N; i++){
#ifdef MDIMARRAY
	    mpfi_mul(hilbert[i][i], hilbert[i][i], hdr);
#else
        mpfi_mul(ptr(hilbert, i, i), ptr(hilbert, i, i), hdr);
#endif // MDIMARRAY
        mpfr_set_si(toI, i, MPFR_RNDN);
        for (int j= i+1; j < N; j ++){
            mpfr_set_si(toJ, j, MPFR_RNDN);
            mpfr_sub(toJ, toJ, toI, MPFR_RNDN);
            mpfr_add(toJ, toJ, one, MPFR_RNDN);
            mpfi_interv_fr(div,toJ,toJ);
            mpfi_div(div, hdr, div);

#ifdef MDIMARRAY
	        mpfi_mul(hilbert[i][j], hilbert[i][j], div);
            mpfi_mul(hilbert[j][i], hilbert[j][i], div);
#else
            mpfi_mul(ptr(hilbert, i, j), ptr(hilbert, i, j), div);
            mpfi_mul(ptr(hilbert, j, i), ptr(hilbert, j, i), div);
#endif // MDIMARRAY
        }
    }
}

void reset(){
    if (boolsk == true){
        //setSkyline();
        setMM();
    }else{
        setDense();
    }
    //mulDiagonal();

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
        }
    }
}

void Uset(int i,int j){
/*
	   mpfi_t SUM;
	   mpfi_t MUL;
    mpfi_init2(SUM, acc);
    mpfi_init2(MUL, acc);
	   mpfi_set_str(SUM,"0",10);
    //mpfi_set_str(MUL,"0",10);
*/
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

void Lset(int j,int i){
/*
	   mpfi_t SUM;
	   mpfi_t MUL;
    mpfi_init2(SUM, acc);
    mpfi_init2(MUL, acc);
    mpfi_set_str(SUM, "0", 10);
*/
	   for (int k = 0; k < i; k++) {
#ifdef MDIMARRAY
	       mpfi_mul(MUL[j][i], hilbert[j][k], hilbert[k][i]);
	       mpfi_add(SUM[j][i], SUM[j][i], MUL[j][i]);
#else
           mpfi_mul(ptr(MUL, j, i), ptr(hilbert, j, k), ptr(hilbert, k, i));
           mpfi_add(ptr(SUM, j, i), ptr(SUM, j, i), ptr(MUL, j, i));
#endif // MDIMARRAY
	   }
#ifdef MDIMARRAY
	   mpfi_sub(SUM[j][i], hilbert[j][i], SUM[j][i]);
	   mpfi_div(hilbert[j][i], SUM[j][i], hilbert[i][i]);
#else
       mpfi_sub(ptr(SUM, j, i), ptr(hilbert, j, i), ptr(SUM, j, i));
       mpfi_div(ptr(hilbert, j, i), ptr(SUM, j, i), ptr(hilbert, i, i));
#endif // MDIMARRAY
    //printMatrix((__mpfi_struct *)L);
}


void comp(void) {

	   mpfi_t tmp;
    mpfi_init2(tmp, acc);
    // forward substitution
    for (int i = 1; i < N; i++) {
        for (int j = 0; j <= i - 1; j++) {
#ifdef MDIMARRAY
            mpfi_mul(tmp, b[j], hilbert[i][j]);
            mpfi_sub(b[i], b[i], tmp);
#else
            mpfi_mul(tmp, &b[j], ptr(hilbert, i, j));
            mpfi_sub(&b[i], &b[i], tmp);
#endif // MDIMARRAY
        }
    }

    // backward substitution
    for (int i = N-1; i >= 0; i--) {
        for (int j = N-1; j > i; j--) {
#ifdef MDIMARRAY
            mpfi_mul(tmp, b[j], hilbert[i][j]);
            mpfi_sub(b[i], b[i], tmp);
#else
            mpfi_mul(tmp, &b[j], ptr(hilbert, i, j));
            mpfi_sub(&b[i], &b[i], tmp);
#endif // MDIMARRAY
        }
#ifdef MDIMARRAY
        mpfi_div(b[i], b[i], hilbert[i][i]);
#else
        mpfi_div(&b[i], &b[i], ptr(hilbert, i, i));
#endif // MDIMARRAY
    }

    // print results
	   printf("\n");
    mpfr_exp_t exp;
    for (int i = 0; i < N; i++) {

        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
        printf("[%sx(%d), ", buf, (int)exp);
        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
        printf("%sx(%d)]\n", buf, (int)exp);

	       //mpfr_printf("%.128RNf\n",b[i]);
    }
	   printf("\n");

    // deallocate
    for (int i = 0; i < N; i++) {
#ifdef MDIMARRAY        
        mpfi_clear(b[i]);
#else
        mpfi_clear(&b[i]);
#endif // MDIMARRAY
        for (int j = 0; j < N; j++) {
#ifdef MDIMARRAY
            mpfi_clear(hilbert[i][j]);
#else
            mpfi_clear(ptr(hilbert, i, j));
#endif // MDIMARRAY
        }
    }
    mpfi_clear(tmp1);
    mpfi_clear(tmp2);
}

void printInterval(__mpfi_struct *b) {
/*
	   for (int i = 0;i < N;i++){
	       mpfi_printf("%.128RNf\n",b[i]);
	   }
*/

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
/*
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

    mpfr_exp_t exp;
    for (int i = 0; i < N; i++) {
        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->left, MPFR_RNDD);
        printf("[%sx(%d), ", buf, (int)exp);
        mpfr_get_str(buf, &exp, 10, 15,
            &((__mpfi_struct *)&(b[i]))->right, MPFR_RNDU);
        printf("%sx(%d)]\n", buf, (int)exp);
    }


    for (int i = 0; i < N; i++) {
        printInterval((__mpfr_struct *)&(b[i]));
    }
    for (int i = 0; i < N; i++) {
        mpfr_clear(b[i]);
    }
    mpfr_clear(one);
}
*/

// void printMatrix3(__mpfi_struct *array) {
void printMatrix3(void) {
    for (int i = N-5; i < N; i++) {
        for (int j = N-5; j < N; j++) {
            printf("(%d, %d) = ", i, j);
#ifdef MDIMARRAY
           printInterval((__mpfi_struct *)&(hilbert[i][j]));
#else
           printInterval(ptr(hilbert, i, j));
#endif // MDIMARRAY
        }
        printf("\n");
    }
}

void allocArrays(int size) {
    N = size;
    hilbert = (__mpfi_struct *)calloc(N * N, sizeof(__mpfi_struct));
    b = (__mpfi_struct *)calloc(N, sizeof(__mpfi_struct));
    SUM = (__mpfi_struct *)calloc(N * N, sizeof(__mpfi_struct));
    MUL = (__mpfi_struct *)calloc(N * N, sizeof(__mpfi_struct));
}
