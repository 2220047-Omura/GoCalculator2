#ifndef _CROUT_
#define _CROUT_

#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
//#include <mpfr.h>
#include <float.h>

//void printInterval(__mpfi_struct *b);
//void comp(void);

extern int N;
#define DOUBLE
#define COUNT

#define ptr(p, i, j) (&(p[(i) * N + (j)]))
// extern int acc = 1024;
extern int acc;
extern char buf[256];

#ifdef DOUBLE
extern double *A;
extern double *b;
extern double *SUM;
extern double *MUL;
#else
extern __mpfi_struct *A;
extern __mpfi_struct *b;
extern __mpfi_struct *SUM;
extern __mpfi_struct *MUL;
#endif // DOUBLE

#ifdef COUNT
extern int *countAdd;
extern int *countMul;
#endif // COUNT


void setMMFilename(const char *fname);

int def(void);

int getN(void);

// void printMatrix(__mpfi_struct *array);

int init(void);

void reset(void);

void Uset(int i,int j);

void Lset(int j,int i);

// void comp(void);

//void printM(__mpfi_struct *b);

void printInterval(__mpfi_struct *b);

void printMatrix3(void);

void InfoAdd(void);

void InfoMul(void);

void allocArrays(int size);

#endif // _CROUT_
