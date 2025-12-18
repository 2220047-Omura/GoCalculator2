#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <stdbool.h>
#include <float.h>

#include "skyline.h"

int n;
int E; // Number of Elements
int acc = 1024;
char buf[256];
bool boolsk = true;

static const char *mm_filename = NULL;

int size = 0;    // ← 実行時に決まる
int *isk = NULL;

double *Ask;
//int isk[size];
double *Lsk;

double *SUMsk;
double *MULsk;

double tmp1;
double tmp2;
//mpfi_t tmp;


void setMMFilename(const char *fname) {
    mm_filename = fname;
}


int getN(void){
    return E;
}

int getIsk(int c){
    return isk[c];
}

void getNnz(void){
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
    size = cols; 
    E = nnz;

    fclose(fp);
}

int init(void) {
    // define E
    if (boolsk == true){
        /*
        for (int i = 1; i < size; i++) {
        n -= 1;
        if (n < 0) {
            n = i;
        }
        E = E + n + 1;
        }
        */
        getNnz();
    }else{
        for (int i = 1; i < size; i++) {
        E = E + i + 1;
        }
    }
    isk = malloc(size * sizeof(int));
    if (!isk) {
        perror("malloc isk");
        exit(1);
    }
    Ask = (double *)malloc(E * sizeof(double));
    Lsk = (double *)malloc(E * sizeof(double));
    SUMsk = (double *)malloc(E * sizeof(double));
    MULsk = (double *)malloc(E * sizeof(double));
	return 0;
}

void setSkylineOrg(){
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

typedef struct {
    int row;
    double val;
} Entry;


int cmp_row(const void *a, const void *b) {
    return ((Entry *)a)->row - ((Entry *)b)->row;
}

void setMM(){
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
    for (int col = 1; col <= cols; col++) {
        rewind(fp);

    /* 再びヘッダをスキップ */
        do {
            fgets(line, sizeof(line), fp);
        } while (line[0] == '%');
        fgets(line, sizeof(line), fp); /* サイズ行 */

        /* この列の非ゼロ要素を一時保存 */
        Entry *tmp = malloc(nnz * sizeof(Entry));
        int cnt = 0;

        for (int t = 0; t < nnz; t++) {
            int i, j;
            double val;
            fscanf(fp, "%d %d %lf", &i, &j, &val);
            if (i <= j && j == col) {
                tmp[cnt].row = i - 1; /* 0 始まり */
                tmp[cnt].val = val;
                cnt++;
            }
        }

        /* 行番号で昇順ソート */
        qsort(tmp, cnt, sizeof(Entry), cmp_row);

        isk[col-1] = k;
        /* Ask に格納 */
        for (int p = 0; p < cnt; p++) {
            Ask[k++] = tmp[p].val;
        }
        free(tmp);
    }
    fclose(fp);

    /* 確認用出力 
    printf("Ask (column-major, row-sorted):\n");
    for (int i = 0; i < nnz; i++) {
        printf("%10.6f\n", Ask[i]);
    }
    */
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
	}
}

void reset(){
    srand(0);
    if (boolsk == true){
        //setSkyline();
        setMM();
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

void printMatrix3(void) {
    for (int i = E-3; i < E; i++) {
        printf("(%d) = ", i);
        printf("%lf ",Ask[i]);
        printf("\n");
    }
}