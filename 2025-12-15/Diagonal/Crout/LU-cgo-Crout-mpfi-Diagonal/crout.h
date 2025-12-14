#ifndef _CROUT_
#define _CROUT_

#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
//#include <mpfr.h>

void printInterval(__mpfi_struct *b);
void comp(void);

#define N 500
#define ptr(p, i, j) (&(p[(i) * N + (j)]))
// extern int acc = 1024;
extern int acc;
extern char buf[256];

extern mpfi_t hilbert[N][N];
extern mpfi_t b[N];

extern mpfi_t SUM[N][N];
extern mpfi_t MUL[N][N];

extern mpfi_t tmp1;
extern mpfi_t tmp2;
//extern mpfi_t tmp;

int def(void);

void printMatrix(__mpfi_struct *array);

int init(void);

void reset(void);

void Uset(int i,int j);

void Lset(int j,int i);

void comp(void);

void printM(__mpfi_struct *b);

void printInterval(__mpfi_struct *b);

void printMatrix3(void);


#endif // _CROUT_
