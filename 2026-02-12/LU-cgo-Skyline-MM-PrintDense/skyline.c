#include <stdio.h>
#include <stdlib.h>
// #include <time.h>
#include <stdbool.h>
#include <float.h>
#include <mpfi.h>
#include <mpfi_io.h>

#include "skyline.h"

#define DOUBLE
#define COUNT

int n;
int E; // Number of Elements
int acc = 1024;
char buf[256];

static const char *mm_filename = NULL;

int size = 0;     // ← 実行時に決まる 正方行列の一辺の大きさ
int MAXp;
int *Dia = NULL;  // Diagonal. Dia[n]に行(列)番号nの対角要素がAskのどこにあるのかを格納
int *isk = NULL;  // isk[n]にnの行番号を格納
int *jsk = NULL;  // isk[n]にnの列番号を格納
int *prof = NULL; // profile. prof[n]にnの上に非ゼロ要素が何個あるかを格納
int *arrM;

#ifdef DOUBLE
double *Ask;
double *Lsk;

double *SUMsk;
double *MULsk;

// double *Ask2;
// double *Bsk;
// double *Xsk;
#else
mpfi_t *Ask;
mpfi_t *Lsk;

mpfi_t *SUMsk;
mpfi_t *MULsk;

// mpfi_t *Ask2;
// mpfi_t *Bsk;
// mpfi_t *Xsk;
#endif //BOUDLE

#ifdef COUNT
int *countAdd;
int *countMul;
int *countS;
int *countS2;
#endif //COUNT

void setMMFilename(const char *fname) {
    mm_filename = fname;
}

int getN(void) {
    return E;
}

int getIsk(int c) {
    return Dia[c];
}

typedef struct {
    int col;
    int row;
} Entry2;

int cmp_col_major2(const void *a, const void *b) {
    Entry2 *entryA = (Entry2 *)a;
    Entry2 *entryB = (Entry2 *)b;

    // まず列(col)で比較
    if (entryA->col != entryB->col) {
        return entryA->col - entryB->col;
    }
    
    // 列が同じ場合、行(row)で比較
    return entryA->row - entryB->row;
}

void getNnz(void) {
    FILE *fp;
    int rows, cols, nnz;
    int i, j;
    double val;
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
    // printf("%d %d %d\n", rows, cols, nnz);
    size = cols;

    /* この列の非ゼロ要素を一時保存 */
    Entry2 *tmp = malloc(nnz * sizeof(Entry2));

    for (int k = 0; k < nnz; k++) {
        fgets(line, sizeof(line), fp);
        sscanf(line, "%d %d %lf", &i, &j, &val);
        if (i > j) {
            tmp[k].col = i - 1; /* 0 始まり */
            tmp[k].row = j - 1;
        }else{
            tmp[k].row = i - 1; /* 0 始まり */
            tmp[k].col = j - 1;
        }
    }

    qsort(tmp, nnz, sizeof(Entry2), cmp_col_major2);

    int iE = 0;
    int jE = 0;
    for (int k = 0; k < nnz; k++) {
        if (jE != tmp[k].col) {
            jE = tmp[k].col;
            iE = tmp[k].row;
        }
        if (tmp[k].row == tmp[k].col) {
            E += (tmp[k].row + 1) - iE;
            //jE = tmp[k].col;
        }
        //printf("i, j, E = %d, %d, %d\n",tmp[k].row,tmp[k].col, E);
    }
    // printf("E from getNnz = %d\n",E);
    // E = nnz;

    fclose(fp);
}

int init(void) {
    // define E
    getNnz();

    Dia = malloc(size * sizeof(int));
    isk = malloc(E * sizeof(int));
    jsk = malloc(E * sizeof(int));
    prof = malloc(E * sizeof(int));
    if (!isk) {
        perror("malloc isk");
        exit(1);
    }

#ifdef DOUBLE
    Ask = (double *)calloc(E, sizeof(double));
    Lsk = (double *)calloc(E, sizeof(double));
    SUMsk = (double *)calloc(E, sizeof(double));
    MULsk = (double *)calloc(E, sizeof(double));
    // Ask2 = (double *)malloc(E, sizeof(double));
    // Bsk = (double *)malloc(size, sizeof(double));
    // Xsk = (double *)malloc(size, sizeof(double));
#else
    Ask = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    Lsk = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    SUMsk = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    MULsk = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    // Ask2 = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    // Bsk = (mpfi_t *)malloc(size * sizeof(mpfi_t));
    // Xsk = (mpfi_t *)malloc(size * sizeof(mpfi_t));

    for (int i = 0; i < E; i++) {
        mpfi_init2(Ask[i], acc);
        mpfi_init2(Lsk[i], acc);
        mpfi_init2(SUMsk[i], acc);
        mpfi_init2(MULsk[i], acc);
        // mpfi_init2(Ask2[i], acc);
    }
    /*
    for (int i = 0; i < size; i++) {
        mpfi_init2(Bsk[i], acc);
        mpfi_init2(Xsk[i], acc);
    }
    */
#endif //DOUBLE

#ifdef COUNT
    countAdd = (int *)calloc(E, sizeof(int));
    countMul = (int *)calloc(E, sizeof(int));
    countS = (int *)calloc(size, sizeof(int));
    countS2 = (int *)calloc(size, sizeof(int));
#endif // COUNT

    return 0;
}

typedef struct {
    int row;
    int col;
    double val;
} Entry;

int cmp_col_major(const void *a, const void *b) {
    Entry *entryA = (Entry *)a;
    Entry *entryB = (Entry *)b;

    // まず列(col)で比較
    if (entryA->col != entryB->col) {
        return entryA->col - entryB->col;
    }
    
    // 列が同じ場合、行(row)で比較
    return entryA->row - entryB->row;
}

void setMM() {
#ifdef DOUBLE
#else
    mpfr_t a;
    mpfr_init2(a, acc);
#endif //DOUBLE
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

    /* この列の非ゼロ要素を一時保存 */
    Entry *tmp = malloc(nnz * sizeof(Entry));

    /*
    // printf("in setMM\n");
    for (int n = 0; n < nnz; n++) {
        int i, j;
        double val;
        fgets(line, sizeof(line), fp);
        sscanf(line, "%d %d %lf", &i, &j, &val);
        //printf("n, i, j = %d,%d,%d\n",n,i-1,j-1);
        tmp[n].row = i - 1;
        tmp[n].col = j - 1;
        tmp[n].val = val;
        //("%lf\n",val);
    }
    */

    int i, j;
    for (int n = 0; n < nnz; n++) {
        double val;
        fgets(line, sizeof(line), fp);
        sscanf(line, "%d %d %lf", &i, &j, &val);
        if (i > j) {
            tmp[n].col = i - 1; /* 0 始まり */
            tmp[n].row = j - 1;
        }else{
            tmp[n].row = i - 1; /* 0 始まり */
            tmp[n].col = j - 1;
        }
        tmp[n].val = val;
        //("%lf\n",val);
    }
    
    qsort(tmp, nnz, sizeof(Entry), cmp_col_major);

    /*
    FILE *fp2;
    fp2 = fopen("copyMM.mtx", "w");

    for (int n = 0; n < nnz; n++) {
        fprintf(fp2, "%d %d %20.15le\n", tmp[n].row + 1, tmp[n].col + 1,tmp[n].val );
    }
    fclose(fp2);
    */

    /* Ask に格納 */
    int k = 0;
    int m = 0;
    int p = 0;
    int z1 = 0;
    int cnt = 0;
    MAXp = 0;
#ifdef DOUBLE
#else
    mpfr_t zero;
    mpfr_init2(zero, acc);
    mpfr_set_str(zero, "0",10, MPFR_RNDN);
#endif // DOUBLE

    for (int n = 0; n < nnz; n++)
    {
        if (tmp[n].row <= tmp[n].col)
        {
            for (int m = z1 +1; m < tmp[n].row; m++) {
#ifdef DOUBLE
                Ask[k] = 0;
#else
                mpfi_interv_fr(Ask[k], zero, zero);
#endif // DOUBLE
                isk[k] = m;
                jsk[k] = tmp[n].col;
                prof[k] = p;
                p++;
                k++;
            }
#ifdef DOUBLE
            Ask[k] = tmp[n].val;
#else
            mpfr_set_d(a, tmp[n].val, MPFR_RNDN);
            mpfi_interv_fr(Ask[k], a, a);
            // mpfi_interv_fr(Ask2[k], a, a);
#endif // DOUBLE
            isk[k] = tmp[n].row;
            z1 = tmp[n].row;
            jsk[k] = tmp[n].col;
            prof[k] = p;
            p++;
            if (tmp[n].row == tmp[n].col) {
                cnt++;
                Dia[m] = k;
                if (MAXp < p) {
                    MAXp = p;
                }
                m++;
                p = 0;
            }
            k++;
        }
    }
    if (cnt != cols) {
        printf("対角要素がありません\n");
        exit(1);
    }
    arrM = (int *)malloc((MAXp + 1) * sizeof(int));
    free(tmp);
    fclose(fp);
}

void reset() {
    setMM();
#ifdef DOUBLE
    for (int i = 0; i < E; i++) {
        Lsk[i] = 0;
        SUMsk[i] = 0;
        MULsk[i] = 0;
    }
    /*
    for (int i = 0; i < size; i++) {
        mpfi_set_str(Bsk[i], "0", 10);
        mpfi_set_str(Xsk[i], "0", 10);
    }
    mpfi_set_str(Bsk[0], "1", 10);
    mpfi_set_str(Xsk[0], "1", 10);
    */
#else
    for (int i = 0; i < E; i++) {
        mpfi_set_str(Lsk[i], "0", 10);
        mpfi_set_str(SUMsk[i], "0", 10);
        mpfi_set_str(MULsk[i], "0", 10);
    }
    /*
    for (int i = 0; i < size; i++) {
        mpfi_set_str(Bsk[i], "0", 10);
        mpfi_set_str(Xsk[i], "0", 10);
    }
    mpfi_set_str(Bsk[0], "1", 10);
    mpfi_set_str(Xsk[0], "1", 10);
    */
#endif // DOUBLE

#ifdef COUNT
    for (int i = 0; i < size; i ++) {
        countS[i] = 0;
        countS2[i] = 0;
    }
    for (int i = 0; i < E; i ++) {
        countAdd[i] = 0;
        countMul[i] = 0;
    }
#endif //COUNT
}

void Usetsk(int m, int l) {
    int s = (prof[m] < prof[l]) ? prof[m] : prof[l];

    
    if (s == 0) {
        return;
    }
    

#ifdef DOUBLE
    for (int k = 0; k < s; k++) {
        Lsk[m] = Ask[l - (s - k)] / Ask[Dia[isk[l - (s - k)]]];
        MULsk[m] = Lsk[m] * Ask[m - (s - k)];
        SUMsk[m] = SUMsk[m] + MULsk[m]; 
    }
    Ask[m] = Ask[m] - SUMsk[m];
#else
    for (int k = 0; k < s; k++) {
        // printf("l-(s-k)/isk[l-(s-k)] = %d/%d\n",l-(s-k),Dia[isk[l-(s-k)]]);
        mpfi_div(Lsk[m], Ask[l - (s - k)], Ask[Dia[isk[l - (s - k)]]]);
        mpfi_mul(MULsk[m], Lsk[m], Ask[m - (s - k)]);
        mpfi_add(SUMsk[m], SUMsk[m], MULsk[m]);
    }
    mpfi_sub(Ask[m], Ask[m], SUMsk[m]);
#endif //DOUBLE

#ifdef COUNT
    countAdd[m] += s;
    countMul[m] += s;
    countS[isk[m]] += s;
    countS2[isk[m]] += s*s;
    if (isk[m] == 1) {
        printf("s :%d\n", s);
    }
#endif //COUNT
}

int getS(int m) {
    return countS[m];
}

int getS2(int m) {
    return countS2[m];
}

void cleanCountS() {
#ifdef COUNT
    for (int i = 0; i < size; i ++) {
        countS[i] = 0;
        countS2[i] = 0;
    }
#endif //COUNT
}

int getprof(int m) {
    return prof[m];
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

void printInterval2(__mpfi_struct *b) {
    char buf[256];
    mpfr_exp_t exp;
    mpfr_get_str(buf, &exp, 10, 15,
                 // &((__mpfi_struct *)&(b))->left, MPFR_RNDD);
                 &(b->left), MPFR_RNDD);
    printf("[%sx(%d), ", buf, (int)exp);
    mpfr_get_str(buf, &exp, 10, 15,
                 &(b->right), MPFR_RNDU);
    printf("%sx(%d)] ", buf, (int)exp);
}

void allocArrays() {
    free(Ask);
    free(Lsk);
    free(SUMsk);
    free(MULsk);
    // free(Ask2);
    // free(Bsk);
    // free(Xsk);
}

void printMatrix3() {
    for (int i = E-3; i < E; i++) {
#ifdef DOUBLE
        printf("%f",Ask[i]);
#else
        printInterval((__mpfi_struct *)&(Ask[i]));
#endif //DOUBLE
    }
}

void InfoAdd(void) {
#ifdef COUNT
    printf("-----InfoAdd-----\n");
    double ave = 0.0;
    for (int i = 0; i < E; i ++) {
        ave += countAdd[i];
    } 
    ave /= (double)(E);

    double var = 0.0;
    for (int i = 0; i < E; i ++) {
        var += countAdd[i]*countAdd[i];
    }
    var = (var / (E)) - (ave * ave);
    printf("average:%f\n", ave);
    printf("variance:%f\n", var);
#endif //COUNT
}

void InfoMul(void) {
#ifdef COUNT
    printf("-----InfoMul-----\n");
    double ave = 0.0;
    for (int i = 0; i < E; i ++) {
        ave += countMul[i];
    } 
    ave /= (double)(E);

    double var = 0.0;
    for (int i = 0; i < E; i ++) {
        var += countMul[i]*countMul[i];
    }
    var = (var / (E)) - (ave * ave);
    printf("average:%f\n", ave);
    printf("variance:%f\n", var);
#endif //COUNT
}

/*
void Norm() {
    mpfi_t tmp;
    mpfi_init2(tmp, acc);
    mpfi_set_str(tmp, "0", 10);

    // forward substitution
    int l;
    for (int a = 1; a < E; a++)
    {
        if (isk[a] != jsk[a])
        {
            printf("a = %d ", a);
            // printf("Ask[%d] / Ask[%d] \n",Ask[n],Ask[Dia[isk[n]]]);
            // printInterval((__mpfi_struct *)&(Ask[n]));
            // printInterval((__mpfi_struct *)&(Ask[Dia[isk[n]]]));
            mpfi_div(tmp, Ask[a], Ask[Dia[isk[a]]]);
            printInterval((__mpfi_struct *)&(tmp));
            mpfi_mul(tmp, Xsk[isk[a]], tmp);
            // mpfi_sub(Xsk[jsk[a]], Bsk[jsk[a]], tmp);
            mpfi_sub(Xsk[jsk[a]], Xsk[jsk[a]], tmp);
            // printf("\n");
        }
    }

    // backward substitution
    for (int a = size - 1; a >= 0; a--)
    {
        l = Dia[a];
        for (int m = E - 1; m >= l; m--)
        {
            if (isk[m] == a)
            {
                mpfi_mul(tmp, Xsk[jsk[m]], Ask[m]);
                // mpfi_sub(Xsk[isk[m]],Bsk[isk[m]],tmp);
                mpfi_sub(Xsk[isk[m]], Xsk[isk[m]], tmp);
            }
        }
        // mpfi_div(Xsk[a],Xsk[a],Ask[l]);
        mpfi_div(Xsk[a], Xsk[a], Ask[l]);
    }

    for (int i = 0; i < size; i++)
    {
        printInterval((__mpfi_struct *)&(Xsk[i]));
    }
    for (int i = 0; i < E; i++)
    {
        // printf("%d\n",prof[i]);
        // printInterval((__mpfi_struct *)&(Ask[i]));
    }

    // Ax
    for (int a = 0; a < size; a++)
    {
        if (isk[a] == 0)
        {
            // j = m - Dia[a-1] - 1;
            mpfi_mul(tmp, Xsk[a], Ask2[a]);
            mpfi_add(Xsk[0], Xsk[0], tmp);
        }
    }
    for (int a = 1; a < size; a++)
    {
        l = Dia[a];
        for (int n = Dia[a - 1] + 1; n < l; n++)
        {
            mpfi_div(tmp, Ask2[n], Ask2[Dia[isk[n]]]);
            mpfi_mul(tmp, Xsk[isk[n]], tmp);
            mpfi_add(Xsk[a], Xsk[a], tmp);
        }
        for (int m = l; m < E; m++)
        {
            if (isk[m] == a)
            {
                mpfi_mul(tmp, Xsk[jsk[m]], Ask2[m]);
                mpfi_add(Xsk[a], Xsk[a], tmp);
            }
        }
    }

    mpfi_t norm;
    mpfi_init2(norm, acc);
    mpfi_set_str(norm, "0", 10);
    for (int i = 0; i < size; i++)
    {
        mpfi_sub(Xsk[i], Xsk[i], Bsk[i]);
        mpfi_mul(Xsk[i], Xsk[i], Xsk[i]);
        mpfi_add(norm, norm, Xsk[i]);
    }
    mpfi_sqrt(norm, norm);

    printf("norm = ");
    printInterval((__mpfi_struct *)&(norm));
}
*/
void printSquare() {
    int N = 5; // 表示するサイズの大きさ
    int j, predj;
#ifdef DOUBLE
    double l = 0;
#else
    mpfi_t zero, l;
    mpfi_init2(zero, acc);
    mpfi_init2(l, acc);
    mpfi_set_str(zero, "0", 10);
    mpfi_set_str(l, "0", 10);
#endif // DOUBLE

#ifdef DOUBLE
    int p = 0; // 1行に表示するL要素の個数
    for (int a = size - N; a < size; a++) { // aの値が行列Aの何行目について処理するか示す
        printf("| ");

        // L要素の表示
        for (int n = 0; n < p - prof[Dia[a]]; n++) {
            printf("%f ",0.0);
        }
        int m = p - prof[Dia[a]];
        if (m < 0) {
            m = 0;
        }
        for (int n = m; n < p; n++) {
            l = Ask[Dia[a] - (p - n)] / Ask[Dia[isk[Dia[a] - (p - n)]]];
            printf("%f ",l);
        }
        p += 1;

        // U要素の表示
        predj = jsk[Dia[a]];
        for (int n = Dia[a]; n < E; n++) {
            if (isk[n] == a) {
                j = jsk[n];
                for (int m = 0; m < j - predj; m++) {
                    printf("%f ",0.0);
                }
                printf("%f ",Ask[n]);
                predj = j + 1;
            }
        }
        printf(" |\n");
    }
#else
    int p = 0; // 1行に表示するL要素の個数
    for (int a = size - N; a < size; a++) { // aの値が行列Aの何行目について処理するか示す
        printf("| ");

        // L要素の表示
        for (int n = 0; n < p - prof[Dia[a]]; n++) {
            printInterval2((__mpfi_struct *)&(zero));
        }
        int m = p - prof[Dia[a]];
        if (m < 0) {
            m = 0;
        }
        for (int n = m; n < p; n++) {
            mpfi_div(l, Ask[Dia[a] - (p - n)], Ask[Dia[isk[Dia[a] - (p - n)]]]);
            printInterval2((__mpfi_struct *)&(l));
        }
        p += 1;

        // U要素の表示
        predj = jsk[Dia[a]];
        for (int n = Dia[a]; n < E; n++) {
            if (isk[n] == a) {
                j = jsk[n];
                for (int m = 0; m < j - predj; m++) {
                    printInterval2((__mpfi_struct *)&(zero));
                }
                printInterval2((__mpfi_struct *)&(Ask[n]));
                predj = j + 1;
            }
        }
        printf(" |\n");
    }
#endif // DOUBLE
}