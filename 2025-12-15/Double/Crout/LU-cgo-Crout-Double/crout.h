#ifndef _CROUT_
#define _CROUT_

#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <float.h>

//void printInterval(__double_struct *b);
void comp(void);

#define N 500
#define ptr(p, i, j) (&(p[(i) * N + (j)]))
// extern int acc = 1024;
extern int acc;
extern char buf[256];

extern double hilbert[N][N];
extern double b[N];

extern double SUM[N][N];
extern double MUL[N][N];

extern double tmp1;
extern double tmp2;
//extern mpfi_t tmp;

int def(void);

void printMatrix(double A[N][N]);

//int init(void);

void reset(void);

void Uset(int i,int j);

void Lset(int j,int i);

void comp(void);

//void printM(__double_struct *b);

//void printInterval(__double_struct *b);

void printMatrix3(void);


#endif // _CROUT_
