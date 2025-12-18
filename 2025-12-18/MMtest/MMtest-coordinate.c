#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    FILE *fp;
    int rows, cols, nnz;
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

    /* coordinate 形式: 行数 列数 非ゼロ要素数 */
    if (sscanf(line, "%d %d %d", &rows, &cols, &nnz) != 3) {
        fprintf(stderr, "coordinate 形式ではありません\n");
        fclose(fp);
        return 1;
    }

    /* 密行列として確保（初期値 0） */
    double *A = calloc(rows * cols, sizeof(double));
    if (!A) {
        perror("メモリ確保失敗");
        fclose(fp);
        return 1;
    }

    /* 非ゼロ要素を読み込む */
    for (int k = 0; k < nnz; k++) {
        int i, j;
        double val;
        if (fscanf(fp, "%d %d %lf", &i, &j, &val) != 3) {
            fprintf(stderr, "データ読み込みエラー\n");
            free(A);
            fclose(fp);
            return 1;
        }
        /* Matrix Market は 1 始まり */
        printf("(i, j) = (%d, %d)\n", i, j);
        A[(i - 1) + (j - 1) * rows] = val;
    }

    fclose(fp);

    /* 行列を表示 */
    printf("Matrix (%d x %d):\n\n", rows, cols);
    for (int i = 0; i < rows; i++) {
        printf("| ");
        for (int j = 0; j < cols; j++) {
            printf("%10.4f ", A[i + j * rows]);
        }
        printf("|\n");
    }

    free(A);
    return 0;
}
