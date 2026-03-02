#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
//#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
//#include <mpfr.h>
#include <float.h>
#include <math.h>

#include "gauss.h"

static const char *mm_filename = NULL;

//void printInterval(__mpfi_struct *b);
//void comp(void);

//#define DOUBLE
//#define COUNT

int N;
#define ptr(p, i, j) (&(p[(i) * N + (j)]))
int acc = 1024;
char buf[256];
bool boolsk = true;

#ifdef DOUBLE
double *A;
double *b;
double *calc;
double *A2;
#else
__mpfi_struct *A;
__mpfi_struct *b;
__mpfi_struct *calc;
__mpfi_struct *A2;
#endif // DOUBLE

#ifdef COUNT
int *countSub;
int *countMul;
#endif // COUNT

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
    getN();
    //printf("N = %d\n",N);
/*
#ifdef DOUBLE
    double a;
#else
    mpfr_t a;
	mpfr_init2(a, acc);
#endif // DOUBLE
*/
    srand(0);

#ifdef DOUBLE
    A = (double *)calloc(N * N, sizeof(double));
    b = (double *)calloc(N, sizeof(double));
    calc = (double *)calloc(N * N, sizeof(double));
    A2 = (double *)calloc(N * N, sizeof(double));
#else
    A = (__mpfi_struct *)malloc(N * N * sizeof(__mpfi_struct));
    b = (__mpfi_struct *)malloc(N * sizeof(__mpfi_struct));
    calc = (__mpfi_struct *)malloc(N * N * sizeof(__mpfi_struct));
    A2 = (__mpfi_struct *)malloc(N * N * sizeof(__mpfi_struct));
#endif // DOUBLE

#ifdef COUNT
    countSub = (int *)calloc(N * N, sizeof(int));
    countMul = (int *)calloc(N * N, sizeof(int));
#endif // COUNT

    // allocate
    for (int i = 0; i < N; i++) {
#ifdef DOUBLE
#else
        mpfi_init2(&b[i], acc);
#endif // MDIMARRAY
        for (int j = 0; j < N; j++) {
#ifdef DOUBLE
#else
            mpfi_init2(ptr(A, i, j), acc);
            mpfi_init2(ptr(A2, i, j), acc);
            mpfi_init2(ptr(calc, i, j), acc);
#endif // DOUBLE
        }
    }
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
#ifdef DOUBLE
        *ptr(A,i,j) = 0;
        *ptr(A2,i,j) = 0;
#else
        mpfi_set_str(ptr(A,i,j), "0", 10);
        mpfi_set_str(ptr(A2,i,j), "0", 10);
#endif // DOUBLE
        }
    }
	return 0;
}

// typedef struct
// {
//     int row;
//     int col;
//     double val;
// } Entry;

void setMM() {
#ifdef DOUBLE
#else
    mpfr_t a;
    mpfr_init2(a, acc);
#endif //DOUBLE
    FILE *fp;
    int rows, cols, nnz;
    char line[256];

    if (!mm_filename)
    {
        fprintf(stderr, "MM filename is not set\n");
        exit(1);
    }
    fp = fopen(mm_filename, "r");

    /* ヘッダ・コメント行をスキップ */
    do
    {
        if (!fgets(line, sizeof(line), fp))
        {
            fprintf(stderr, "不正なファイル形式です\n");
            fclose(fp);
        }
    } while (line[0] == '%');

    /* coordinate 形式: 行数 列数 非ゼロ要素数 */
    if (sscanf(line, "%d %d %d", &rows, &cols, &nnz) != 3)
    {
        fprintf(stderr, "coordinate 形式ではありません\n");
        fclose(fp);
    }

    /* この列の非ゼロ要素を一時保存 */
    //Entry *tmp = malloc(nnz * sizeof(Entry));

    // printf("in setMM\n");
    for (int n = 0; n < nnz; n++)
    {
        int i, j;
        double val;
        fgets(line, sizeof(line), fp);
        sscanf(line, "%d %d %lf", &i, &j, &val);
#ifdef DOUBLE
        *ptr(A, i-1, j-1) = val;
        *ptr(A, j-1, i-1) = val;
        *ptr(A2, i-1, j-1) = val;
        *ptr(A2, j-1, i-1) = val;
#else
        mpfr_set_d(a, val, MPFR_RNDN);
        mpfi_interv_fr(ptr(A, i-1, j-1), a, a);
        mpfi_interv_fr(ptr(A, j-1, i-1), a, a);
        mpfi_interv_fr(ptr(A2, i-1, j-1), a, a);
        mpfi_interv_fr(ptr(A2, j-1, i-1), a, a);
#endif //DOUBLE
        ////printf("n, i, j = %d,%d,%d\n",n,i-1,j-1);
        // tmp[n].row = i - 1; /* 0 始まり */
        // tmp[n].col = j - 1;
        // tmp[n].val = val;
        ////("%lf\n",val);
    }

//     for (int n = 0; n < nnz; n++) {
// #ifdef DOUBLE
//         *ptr(A,tmp[n].row,tmp[n].col) = tmp[n].val;
//         *ptr(A2,tmp[n].row,tmp[n].col) = tmp[n].val;
// #else
//         mpfr_set_d(a, tmp[n].val, MPFR_RNDN);
//         mpfi_interv_fr(ptr(A, tmp[n].row, tmp[n].col), a, a);
//         mpfi_interv_fr(ptr(A2, tmp[n].row, tmp[n].col), a, a);
// #endif // DOUBLE
//     }
    /*
    for (int n = 0; n < nnz; n++) {
        if (tmp[n].row <= tmp[n].col) {
            mpfr_set_d(a, tmp[n].val, MPFR_RNDN);
#ifdef MDIMARRAY
	        mpfi_interv_fr(A[tmp[n].col][tmp[n].row], a, a);
            mpfi_interv_fr(A[tmp[n].row][tmp[n].col], a, a);
#else
            mpfi_interv_fr(ptr(A, tmp[n].col, tmp[n].row), a, a);
            mpfi_interv_fr(ptr(A, tmp[n].row, tmp[n].col), a, a);
#endif // MDIMARRAY
        }
    }
    */
    fclose(fp);
}

void reset(){
    setMM();

    // initialize
#ifdef DOUBLE
    b[0] = 1;
    for (int i = 1; i < N; i++) {
        b[i] = 0;
    }
    for (int i = 0; i < N; i ++) {
        for (int j = 0; j < N; j ++) {
            *ptr(calc,i,j) = 0;
        }
    }
#else
    mpfi_set_str(&b[0], "1", 10);
    for (int i = 1; i < N; i++) {
        mpfi_set_str(&b[i], "0", 10);
    }
    for (int i = 0; i < N; i ++) {
        for (int j = 0; j < N; j ++) {
            mpfi_set_str(ptr(calc, i, j), "0", 10);
        }
    }
#endif // DOUBLE

#ifdef COUNT
    for (int i = 0; i < N; i ++) {
        for (int j = 0; j < N; j ++) {
            *ptr(countSub,i,j) = 0;
            *ptr(countMul,i,j) = 0;
        }
    }
#endif // COUNT
}

void LUfact1(int k, int i) {
#ifdef DOUBLE
    *ptr(A, i, k) = *ptr(A, i, k) / *ptr(A, k, k);
#else
    mpfi_div(ptr(A, i, k), ptr(A, i, k), ptr(A, k, k));
#endif // DOUBLE
}

void LUfact2(int k, int i, int j) {
#ifdef DOUBLE
    *ptr(calc, i, j) = (*ptr(A, i, k)) * (*ptr(A, k, j));
    *ptr(A, i, j) = (*ptr(A, i, j)) - (*ptr(calc, i, j));
#else
    mpfi_mul(ptr(calc, i, j), ptr(A, i, k), ptr(A, k, j));
    mpfi_sub(ptr(A, i, j), ptr(A, i, j), ptr(calc, i, j));
#endif // DOUBLE

#ifdef COUNT
    *ptr(countMul, i, j) += 1;
    *ptr(countSub, i, j) += 1;
#endif // COUNT
}
/*
void comp(void) {
	mpfi_t tmp;
    mpfi_init2(tmp, acc);
    // forward substitution
    for (int i = 1; i < N; i++) {
        for (int j = 0; j <= i - 1; j++) {
#ifdef MDIMARRAY
            mpfi_mul(tmp, b[j], A[i][j]);
            mpfi_sub(b[i], b[i], tmp);
#else
            mpfi_mul(tmp, &b[j], ptr(A, i, j));
            mpfi_sub(&b[i], &b[i], tmp);
#endif // MDIMARRAY
        }
    }

    // backward substitution
    for (int i = N-1; i >= 0; i--) {
        for (int j = N-1; j > i; j--) {
#ifdef MDIMARRAY
            mpfi_mul(tmp, b[j], A[i][j]);
            mpfi_sub(b[i], b[i], tmp);
#else
            mpfi_mul(tmp, &b[j], ptr(A, i, j));
            mpfi_sub(&b[i], &b[i], tmp);
#endif // MDIMARRAY
        }
#ifdef MDIMARRAY
        mpfi_div(b[i], b[i], A[i][i]);
#else
        mpfi_div(&b[i], &b[i], ptr(A, i, i));
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
            mpfi_clear(A[i][j]);
#else
            mpfi_clear(ptr(A, i, j));
#endif // MDIMARRAY
        }
    }
}
*/
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

void printMatrix3(void) {
    for (int i = N-5; i < N; i++) {
        for (int j = N-5; j < N; j++) {
            printf("(%d, %d) = ", i, j);
#ifdef DOUBLE
            printf("%lf", *ptr(A, i, j));
#else
            printInterval(ptr(A, i, j));
#endif // DOUBLE
        }
        printf("\n");
    }
}

void Norm2() {
#ifdef DOUBLE
    double *y1;
    double *x1;
    //double *y2;
    double *b1;
    double mul;
    double sum;
    //double y;
    //double x;
    double L;
    //double tmp;
    double norm = 0.0;
    y1 = (double *)calloc(N, sizeof(double));
    x1 = (double *)calloc(N, sizeof(double));
    //y2 = (double *)calloc(N, sizeof(double));
    b1 = (double *)calloc(N, sizeof(double));

    for (int i = 0; i < N; i ++) {
        for (int j = 0; j <= i; j ++) {
            L = (*ptr(A, j, i)) / (*ptr(A, j, j));
            mul = L * y1[j];
            sum += mul;
        }
        y1[i] = b[i] - sum;
        //printf("y1[%d] = %f\n", i, y1[i]);
        sum = 0;
    }

    for (int i = N -1; i >= 0; i--) {
        for (int j = i; j < N; j++) {
            mul = (*ptr(A, i, j)) * x1[j];
            sum += mul;
        }
        x1[i] = (y1[i] - sum) / (*ptr(A, i, i));
        //printf("x1[%d] = %f, A = %f\n", i, x1[i], (*ptr(A, i, i)));
        sum = 0;
    }

    /*
    for (int i = N -1; i >= 0; i--) {
        for (int j = i; j < N; j++) {
            mul = (*ptr(A, i, j)) * x1[j];
            sum += mul;
        }
        y1[i] = sum;
        sum = 0;
    }

    for (int i = 0; i < N; i ++) {
        for (int j = 0; j <= i; j ++) {
            L = (*ptr(A, j, i)) / (*ptr(A, j, j));
            //printf("L[%d][%d] = %f\n", i, j, L);
            mul = L * y1[j];
            sum += mul;
        }
        b1[i] = sum;
        sum = 0;
    }
    */

    /*
    for (int i = 0; i < N; i++) {
        for (int j = i; j < N; j++) {
            mul = (*ptr(A, i, j)) * x1[j];
            sum += mul;
        }
        y1[i] = sum;
        sum = 0;
    }

    for (int i = 0; i < N; i ++) {
        for (int j = 0; j <= i; j ++) {
            L = (*ptr(A, j, i)) / (*ptr(A, j, j));
            mul = L * y1[j];
            sum += mul;
        }
        b1[i] = sum;
        sum = 0;
    }
    */

    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            mul = (*ptr(A2, i, j)) * x1[j];
            sum += mul;
        }
        b1[i] = sum;
        sum = 0;
    }

    for (int i = 0; i < N; i++) {
        //printf("x1[%d] = %f\n", i, x1[i]);
        //printf("y1[%d] = %f\n", i, y1[i]);
        //printf("b1[%d] = %f\n", i, b1[i]);
        norm += (b[i] - b1[i]) * (b[i] - b1[i]);
    }
    printf("Norm : %f\n ", sqrt(norm));

#else
    __mpfi_struct *y1;
    __mpfi_struct *x1;
    __mpfi_struct *b1;
    mpfi_t mul;
    mpfi_t sum;
    mpfi_t L;
    mpfi_t norm;
    mpfi_t tmp;
    mpfi_t zero;

    y1 = (__mpfi_struct *)malloc(N * sizeof(__mpfi_struct));
    x1 = (__mpfi_struct *)malloc(N * sizeof(__mpfi_struct));
    b1 = (__mpfi_struct *)malloc(N * sizeof(__mpfi_struct));

    mpfi_init2(zero, acc);
    mpfi_set_str(zero, "0", 10);

    for (int i = 0; i < N; i ++) {
        mpfi_init2(&y1[i], acc);
        mpfi_set(&y1[i], zero);
        mpfi_init2(&x1[i], acc);
        mpfi_set(&x1[i], zero);
        mpfi_init2(&b1[i], acc);
        mpfi_set(&b1[i], zero);
    }

    mpfi_init2(mul, acc);
    mpfi_init2(sum, acc);
    mpfi_init2(L, acc);
    mpfi_init2(norm, acc);
    mpfi_init2(tmp, acc);

    mpfi_set(sum, zero);
    mpfi_set(norm, zero);

    for (int i = 0; i < N; i ++) {
        for (int j = 0; j <= i; j ++) {
            //L = (*ptr(A, j, i)) / (*ptr(A, j, j));
            mpfi_div(L, ptr(A, j, i), ptr(A, j, j));
            //mul = L * y1[j];
            mpfi_mul(mul, L, &y1[j]);
            //sum += mul;
            mpfi_add(sum, sum, mul);
        }
        //y1[i] = b[i] - sum;
        mpfi_sub(&y1[i], &b[i], sum);
        //printInterval(&y1[i]);
        //sum = 0;
        mpfi_set(sum, zero);
    }

    for (int i = N -1; i >= 0; i--) {
        for (int j = i; j < N; j++) {
            //mul = (*ptr(A, i, j)) * x1[j];
            mpfi_mul(mul, ptr(A, i, j), &x1[j]);
            //sum += mul;
            mpfi_add(sum, sum, mul);
        }
        //x1[i] = (y1[i] - sum) / (*ptr(A, i, i));
        mpfi_sub(tmp, &y1[i], sum);
        mpfi_div(&x1[i], tmp, ptr(A, i, i));
        //sum = 0;
        mpfi_set(sum, zero);
    }

    /*
    for (int i = N -1; i >= 0; i--) {
        for (int j = i; j < N; j++) {
            //mul = (*ptr(A, i, j)) * x1[j];
            mpfi_mul(mul, ptr(A, i, j), &x1[j]);
            //sum += mul;
            mpfi_add(sum, sum, mul);
        }
        //y1[i] = sum;
        mpfi_set(&y1[i], sum);
        //sum = 0;
        mpfi_set(sum, zero);
    }

    for (int i = 0; i < N; i ++) {
        for (int j = 0; j <= i; j ++) {
            //L = (*ptr(A, j, i)) / (*ptr(A, j, j));
            mpfi_div(L, ptr(A, j, i), ptr(A, j, j));
            //printf("L[%d][%d] = %f\n", i, j, L);
            //mul = L * y1[j];
            mpfi_mul(mul, L, &y1[j]);
            //sum += mul;
            mpfi_add(sum, sum, mul);
        }
        //b1[i] = sum;
        mpfi_set(&b1[i], sum);
        //sum = 0;
        mpfi_set(sum, zero);
    }
    */

    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            //mul = (*ptr(A2, i, j)) * x1[j];
            mpfi_mul(mul, ptr(A2, i, j), &x1[j]);
            //sum += mul;
            mpfi_add(sum, sum, mul);
        }
        //b1[i] = sum;
        mpfi_set(&b1[i], sum);
        //sum = 0;
        mpfi_set(sum, zero);
    }

    for (int i = 0; i < N; i++) {
        //printf("x1[%d] = %f\n", i, x1[i]);
        //printf("y1[%d] = %f\n", i, y1[i]);
        //printf("b1[%d] = %f\n", i, b1[i]);
        printInterval(&b1[i]);
        //norm += (b[i] - b1[i]);
        mpfi_sub(tmp, &b[i], &b1[i]);
        mpfi_mul(tmp, tmp, tmp);
        mpfi_add(norm, norm, tmp);
        //printInterval(norm);
    }
    mpfi_sqrt(norm, norm);
    printf("Norm :");
    printInterval(norm);

    free(y1);
    free(x1);
    free(b1);
#endif //DOUBLE
}

void InfoSub(void) {
#ifdef COUNT
    printf("-----InfoSub-----\n");
    double ave = 0.0;
    for (int i = 0; i < N; i ++) {
        for (int j = 0; j < N; j ++) {
            ave += *ptr(countSub, i, j);
        }
    } 
    ave /= (double)(N*N);

    double var = 0.0;
    for (int i = 0; i < N; i ++) {
        for (int j = 0; j < N; j ++) {
            var += ((double)(*ptr(countSub, i, j)))*((double)(*ptr(countSub, i, j)));
        }
    }
    var = (var / (N * N)) - (ave * ave);
    printf("average:%f\n", ave);
    printf("variance:%f\n", var);
#endif //COUNT
}

void InfoMul(void) {
#ifdef COUNT
    printf("-----InfoMul-----\n");
    double ave = 0.0;
    for (int i = 0; i < N; i ++) {
        for (int j = 0; j < N; j ++) {
            ave += *ptr(countMul, i, j);
        }
    } 
    ave /= (double)(N*N);

    double var = 0.0;
    for (int i = 0; i < N; i ++) {
        for (int j = 0; j < N; j ++) {
            var += ((double)(*ptr(countMul, i, j)))*((double)(*ptr(countMul, i, j)));
        }
    }
    var = (var / (N * N)) - (ave * ave);
    printf("average:%f\n", ave);
    printf("variance:%f\n", var);
#endif //COUNT
}

void allocArrays(int size) {
    N = size;
#ifdef DOUBLE
    A = (double *)calloc(N * N, sizeof(double));
    b = (double *)calloc(N, sizeof(double));
    calc = (double *)calloc(N * N, sizeof(double));
#else
    A = (__mpfi_struct *)calloc(N * N, sizeof(__mpfi_struct));
    b = (__mpfi_struct *)calloc(N, sizeof(__mpfi_struct));
    calc = (__mpfi_struct *)calloc(N * N, sizeof(__mpfi_struct));
#endif //DOUBLE
}
