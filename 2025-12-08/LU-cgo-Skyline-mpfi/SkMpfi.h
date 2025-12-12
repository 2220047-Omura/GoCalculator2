#ifndef _SKMPFI_
#define _SKMPFI_

#include <stdio.h>
#include <stdlib.h>
//#include <time.h>
#include <mpfi.h>
#include <mpfi_io.h>
//#include <mpfr.h>

void printInterval(__mpfi_struct *b);
//void comp(void);

#define size 10
//#define ptr(p, i, j) (&(p[(i) * N + (j)]))
// extern int acc = 1024;
extern int n;
extern int E;
extern int acc;
extern char buf[256];

extern mpfi_t *Ask;
extern int isk[size];
extern mpfi_t *Lsk;

extern mpfi_t *SUMsk;
extern mpfi_t *MUL;

extern mpfi_t tmp1;
extern mpfi_t tmp2;
//extern mpfi_t tmp;

int getN(void);

int getIsk(int c);

//void printMatrix(__mpfi_struct *array);

int init(void);

void setSkyline(void);

void setSkylineTest(void);

void reset(void);

void Usetsk(int b,int i,int j);

void result(void);

void clear(void);
/*
void Uset(int i,int j);

void Lset(int j,int i);

void comp(void);

void printM(__mpfi_struct *b);

void printInterval(__mpfi_struct *b);

void printMatrix3(void);
*/

#endif // _SKMPFI_
