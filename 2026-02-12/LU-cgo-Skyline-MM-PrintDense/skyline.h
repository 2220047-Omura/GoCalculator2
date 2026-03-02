#ifndef _SKYLINE_
#define _SKYLINE_

#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <stdbool.h>
#include <float.h>
#include <mpfi.h>
#include <mpfi_io.h>

//void printInterval(__mpfi_struct *b);
//void comp(void);

#define DOUBLE
//#define COUNT
#define PRINT

extern int size;     // 行列サイズ（列数）
extern int *Dia;     // 動的配列
extern int *isk;
extern int *prof;
extern int MAXp;
extern int *arrM;

//extern int size;
extern int n;
extern int E;
//#define ptr(p, i, j) (&(p[(i) * size + (j)]))
// extern int acc = 1024;
extern int acc;
extern char buf[256];

#ifdef DOUBLE
extern double *Ask;
extern double *Lsk;

extern double *SUMsk;
extern double *MULsk;

// double *Ask2;
// double *Bsk;
// double *Xsk;
#else
extern mpfi_t *Ask;
extern mpfi_t *Lsk;

extern mpfi_t *SUMsk;
extern mpfi_t *MULsk;

// mpfi_t *Ask2;
// mpfi_t *Bsk;
// mpfi_t *Xsk;
#endif //BOUDLE

#ifdef COUNT
extern int *countAdd;
extern int *countMul;
extern int *counS;
extern int *counS2;
#endif //COUNT

void setMMFilename(const char *fname);

//int def(void);
int getN(void);

int getIsk(int c);

//void printMatrix(__mpfi_struct *array);

int init(void);

//void setSkylineOrg(void);

int cmp_row(const void *a, const void *b);

void setMM(void);

//void setDense(void);

void mulDiagonal(void);

//void setSkylineTest(void);

void reset(void);

/*
void Uset(int i,int j);

void Lset(int j,int i);

void comp(void);
*/

//void Usetsk(int b,int i,int j);
void Usetsk(int i,int j);

int getS(int m);
int getS2(int m);

int getLength(int m, int l);

void cleanCountS(void);

int getprof(int m);

//void printM(__mpfi_struct *b);

void printInterval(__mpfi_struct *b);

void printInterval2(__mpfi_struct *b);

void printMatrix3(void);

void InfoAdd(void);

void InfoMul(void);

void freeArrays(void);

void Norm2(void);

void printSquare(void);

#endif // _SKYLINE_
