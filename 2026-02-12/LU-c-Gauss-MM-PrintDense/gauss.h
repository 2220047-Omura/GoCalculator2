#ifndef _GAUSS_
#define _GAUSS_

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
//#define DOUBLE
//#define COUNT

#define ptr(p, i, j) (&(p[(i) * N + (j)]))
// extern int acc = 1024;
extern int acc;
extern char buf[256];

#ifdef DOUBLE
extern double *A;
extern double *b;
extern double *calc;
#else
extern __mpfi_struct *A;
extern __mpfi_struct *b;
extern __mpfi_struct *calc;
#endif // DOUBLE

#ifdef COUNT
extern int *countSub;
extern int *countMul;
#endif // COUNT


void setMMFilename(const char *fname);

int def(void);

int getN(void);

// void printMatrix(__mpfi_struct *array);

int init(void);

void reset(void);

void LUfact1(int k,int i);

void LUfact2(int k,int i,int j);

// void comp(void);

//void printM(__mpfi_struct *b);

void printInterval(__mpfi_struct *b);

void printMatrix3(void);

void InfoSub(void);

void InfoMul(void);

//void Norm(void);

void Norm2(void);

void allocArrays(int size);

#endif // _GAUSS_
