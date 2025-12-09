#ifndef _SKYLINE_
#define _SKYLINE_

#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
//#include <mpfr.h>

void printInterval(__mpfi_struct *b);
//void comp(void);

#define size 300
//extern int size;
extern int n;
extern int N;
//#define ptr(p, i, j) (&(p[(i) * size + (j)]))
// extern int acc = 1024;
extern int acc;
extern char buf[256];

/*
#ifdef MDIMARRAY
extern mpfi_t hilbert[N][N];
extern mpfi_t b[N];
extern mpfi_t SUM[N][N];
extern mpfi_t MUL[N][N];
#else
extern __mpfi_struct *hilbert;
extern __mpfi_struct *b;
extern __mpfi_struct *SUM;
extern __mpfi_struct *MUL;
#endif // MDIMARRAY
*/

mpfi_t *Ask;
int isk[size];
mpfi_t *Lsk;

mpfi_t *SUMsk;
mpfi_t *MULsk;

extern mpfi_t tmp1;
extern mpfi_t tmp2;
//extern mpfi_t tmp;

//int def(void);
int getN(void);

int getIsk(int c);

void printMatrix(__mpfi_struct *array);

int init(void);

void setSkyline(void);

void setSkylineTest(void);

void reset(void);

/*
void Uset(int i,int j);

void Lset(int j,int i);

void comp(void);
*/

void Usetsk(int b,int i,int j);

//void printM(__mpfi_struct *b);

void printInterval(__mpfi_struct *b);

//void printMatrix3(void);

void allocArrays();

#endif // _SKYLINE_
