#ifndef _SKDOUBLE_
#define _SKDOUBLE_

#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <float.h>

//void printInterval(__mpfi_struct *b);
//void comp(void);

#define size 500
//#define ptr(p, i, j) (&(p[(i) * N + (j)]))
// extern int acc = 1024;
extern int n;
extern int E;
extern int acc;
extern char buf[256];

extern double *Ask;
extern int isk[size];
extern double *Lsk;

extern double *SUMsk;
extern double *MUL;


int getN(void);

int getIsk(int c);

//void printMatrix(__mpfi_struct *array);

int init(void);

void setSkyline(void);

void setDense(void);

void setSkylineTest(void);

void reset(void);

void Usetsk(int b,int i,int j);

void printMatrix3(void);

/*
void result(void);

void clear(void);

void Uset(int i,int j);

void Lset(int j,int i);

void comp(void);

void printM(__mpfi_struct *b);

void printInterval(__mpfi_struct *b);
*/

#endif // _SKDOUBLE_
