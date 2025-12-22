#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <stdbool.h>
#include <float.h>
#include <mpfi.h>
#include <mpfi_io.h>

#include "skyline.h"

int n;
int E; // Number of Elements
int acc = 1024;
char buf[256];

static const char *mm_filename = NULL;

int size = 0;    // ← 実行時に決まる
int *Dia = NULL;
int *isk = NULL;
int *jsk = NULL;
int *prof = NULL;

mpfi_t *Ask;
//int isk[size];
mpfi_t *Lsk;

mpfi_t *SUMsk;
mpfi_t *MULsk;

mpfi_t *Ask2;
mpfi_t *Bsk;
mpfi_t *Xsk;

void setMMFilename(const char *fname) {
    mm_filename = fname;
}

int getN(void){
    return E;
}

int getIsk(int c){
    return Dia[c];
}

void getNnz(void){
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
    //printf("%d %d %d\n", rows, cols, nnz);
    size = cols;
    for (int k = 0; k<nnz; k++){
        fscanf(fp, "%d %d %lf", &i, &j, &val);
        if (i <= j) {
            E += 1;
        }
    }
    //printf("E from getNnz = %d\n",E);
    //E = nnz;

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
    Ask = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    Lsk = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    SUMsk = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    MULsk = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    Ask2 = (mpfi_t *)malloc(E * sizeof(mpfi_t));
    Bsk = (mpfi_t *)malloc(size * sizeof(mpfi_t));
    Xsk = (mpfi_t *)malloc(size * sizeof(mpfi_t));
    for (int i = 0; i < E; i++){
        mpfi_init2(Ask[i],acc);
        mpfi_init2(Lsk[i],acc);
        mpfi_init2(SUMsk[i],acc);
        mpfi_init2(MULsk[i],acc);
        mpfi_init2(Ask2[i],acc);
    }
    for (int i = 0; i < size; i++){
        mpfi_init2(Bsk[i],acc);
        mpfi_init2(Xsk[i],acc);
    }
	return 0;
}

typedef struct {
    int row;
    int col;
    double val;
} Entry;

int cmp_row(const void *a, const void *b) {
    return ((Entry *)a)->row - ((Entry *)b)->row;
}

void setMM(){
    mpfr_t a;
    mpfr_init2(a,acc);
    FILE *fp;
    int rows, cols, nnz;
    char line[256];
    int k = 0;

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

        /* この列の非ゼロ要素を一時保存 */
        Entry *tmp = malloc(E * sizeof(Entry));
        int cnt = 0;

        for (int t = 0; t < E; t++) {
            int i, j;
            double val;
            fscanf(fp, "%d %d %lf", &i, &j, &val);
            if (i <= j && j == col) {
                tmp[cnt].row = i - 1; /* 0 始まり */
                tmp[cnt].col = j - 1;
                tmp[cnt].val = val;
                cnt++;
            }
        }

        /* 行番号で昇順ソート */
        qsort(tmp, cnt, sizeof(Entry), cmp_row);

        /* Ask に格納 */
        for (int p = 0; p < cnt; p++) {
            mpfr_set_d(a, tmp[p].val, MPFR_RNDN);
            mpfi_interv_fr(Ask[k], a, a);
            mpfi_interv_fr(Ask2[k], a, a);
            isk[k] = tmp[p].row;
            jsk[k] = tmp[p].col;
            prof[k] = p;
            //printf("prof %d ",prof[k]);
            //printInterval((__mpfi_struct *)&(Ask[k]));
            //printf("col=%d cnt=%d k=%d E=%d\n", col, cnt, k, E);
            k += 1;
        }
        Dia[col-1] = k-1;
        //printf("Dia[%d] = %d\n", col-1,k-1);
        free(tmp);
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
        mpfr_set_si(toIsk, Dia[c], MPFR_RNDN);
        mpfr_set_si(toI, i, MPFR_RNDN);
        mpfr_sub(toIsk, toIsk, toI, MPFR_RNDN);
        mpfr_add(toIsk, toIsk, one, MPFR_RNDN);
        mpfi_interv_fr(div,toIsk,toIsk);
        mpfi_div(div, hdr, div);

        mpfi_mul(Ask[i], Ask[i], div);
        if (i == Dia[c]) {
            c += 1;
        }
    }
}

void reset(){
    setMM();
    //mulDiagonal();
	for (int i = 0; i < E; i++) {
        mpfi_set_str(Lsk[i], "0", 10);
        mpfi_set_str(SUMsk[i], "0", 10);
        mpfi_set_str(MULsk[i], "0", 10);
	}
    for (int i = 0; i < size; i++){
        mpfi_set_str(Bsk[i], "0",10);
        mpfi_set_str(Xsk[i], "0",10);
    }
    mpfi_set_str(Bsk[0], "1",10);
}

void Usetsk(int m,int l) {
    //printf("Hello from (%d)\n",j);
    int s;
    if (prof[m] < prof[l]) {
		s = prof[m];
	} else {
		s = l;
	}

    for (int k = 0; k < s; k++) {
        mpfi_div(Lsk[m],Ask[l-(s-k)],Ask[l]);
        mpfi_mul(MULsk[m],Lsk[m],Ask[m-(s-k)]);
        mpfi_add(SUMsk[m],SUMsk[m],MULsk[m]);
    }
    mpfi_sub(Ask[m], Ask[m], SUMsk[m]);
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

void allocArrays() {
    //memset(isk,0,sizeof(isk));
    free(Ask);
    free(Lsk);
    free(SUMsk);
    free(MULsk);
    free(Ask2);
    free(Bsk);
    free(Xsk);
}

void printMatrix3() {
    for (int i = 0; i < 5; i++) {
        printInterval((__mpfi_struct *)&(Ask[i]));
    }
}

void Norm(){
    mpfi_t tmp;
    mpfi_init2(tmp,acc);
    mpfi_set_str(tmp, "0", 10);

    // forward substitution
    for (int a = 1; a < E; a++) {
        mpfi_div(Lsk[a], Ask[a], Ask[jsk[a]]);
        mpfi_mul(tmp, Bsk[isk[a]], Lsk[a]);
        mpfi_sub(Xsk[jsk[a]], Bsk[jsk[a]], tmp);
    }
    
    // backward substitution
    int l;
    for (int a = size-1; a >= 0; a--) {
        l = Dia[a];
        for (int m = E-1; m >= l; m--) {
		    if (isk[m] == a) {
			    mpfi_mul(tmp,Bsk[jsk[m]],Ask[m]);
                mpfi_sub(Xsk[isk[m]],Bsk[isk[m]],tmp);
		    }
        }
        mpfi_div(Xsk[a],Xsk[a],Ask[l]);
	}

    for (int i = 0; i < size; i++){
        //printInterval((__mpfi_struct *)&(Xsk[i]));
    }
    for (int i = 0; i < E; i++){
        //printf("%d\n",prof[i]);
        //printInterval((__mpfi_struct *)&(Ask2[i]));
    }
   
    //Ax
    for (int a = 0; a < size; a++) {
        if (isk[a] == 0) {
            //j = m - Dia[a-1] - 1;
            mpfi_mul(MULsk[a], Xsk[a], Ask2[a]);
            mpfi_add(Xsk[0], Xsk[0], MULsk[a]);
		}
    }
    for (int a = 1; a < size; a++) {
        l = Dia[a];
        for (int n = Dia[a-1] + 1; n < l; n++) {
            mpfi_div(Lsk[n], Ask2[n], Ask2[Dia[isk[n]]]);
            mpfi_mul(MULsk[n], Xsk[isk[n]], Lsk[n]);
            mpfi_add(Xsk[a], Xsk[a], MULsk[n]);
        }
        for (int m = l; m < E; m++){
		    if (isk[m] == a) {
                mpfi_mul(MULsk[m], Xsk[jsk[m]], Ask2[m]);
                mpfi_add(Xsk[a], Xsk[a], MULsk[m]);
		    }
        }
	}

    mpfi_t norm;
    mpfi_init2(norm, acc);
    mpfi_set_str(norm, "0", 10);
    for (int i = 0; i < size; i++){
        mpfi_sub(Xsk[i],Xsk[i],Bsk[i]);
        mpfi_mul(Xsk[i],Xsk[i],Xsk[i]);
        mpfi_add(norm, norm, Xsk[i]);
    }
    mpfi_sqrt(norm,norm);

    printf("norm = ");
    printInterval((__mpfi_struct *)&(norm));
    
}
