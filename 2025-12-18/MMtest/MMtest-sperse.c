#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[]) {
    FILE *fp;
    int rows, cols,nnz;
    char line[256];

    if (argc != 2) {
        fprintf(stderr, "使い方: %s matrix.mtx\n", argv[0]);
        return 1;
    }

    fp = fopen(argv[1], "r");
    if (!fp) {
        perror("ファイルを開けません");
        return 1;
    }

    /* ヘッダ・コメント行をスキップ */
    do {
        if (!fgets(line, sizeof(line), fp)) {
            fprintf(stderr, "不正なファイル形式です\n");
            fclose(fp);
            return 1;
        }
    } while (line[0] == '%');

    /* 行数・列数を読み込む */
    if (sscanf(line, "%d %d", &rows, &cols) != 2) {
        fprintf(stderr, "行列サイズを読み取れません\n");
        fclose(fp);
        return 1;
    }

    /* メモリ確保 */
    double *A = (double *)malloc(rows * cols * sizeof(double));
    if (!A) {
        perror("メモリ確保失敗");
        fclose(fp);
        return 1;
    }

    /* Matrix Market の array 形式は列優先（column-major） */
    for (int j = 0; j < cols; j++) {
        for (int i = 0; i < rows; i++) {
            if (fscanf(fp, "%lf", &A[i + j * rows]) != 1) {
                fprintf(stderr, "データ読み込みエラー\n");
                free(A);
                fclose(fp);
                return 1;
            }
        }
    }

    fclose(fp);


    /* 行列を表示（行列の形が分かるように枠付きで表示） */
    printf("Matrix (%d x %d):", rows, cols);


    for (int i = 0; i < rows; i++) {
        printf("| ");
        for (int j = 0; j < cols; j++) {
            printf("%12.6f ", A[i + j * rows]);
        }
        printf("|");
    }


    printf("");


    free(A);
    return 0;
}