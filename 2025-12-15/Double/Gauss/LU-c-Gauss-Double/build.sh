#!/bin/sh

#CC="gcc"
#CC="clang"

CC="/opt/homebrew/opt/llvm/bin/clang"
CFLAGS="-O3 -Wall -fopenmp"
INCLUDEPATH="-I/opt/homebrew/include"
LIBPATH="-L. -L/opt/homebrew/lib -Wl,-rpath,/opt/homebrew/lib -Wl,-rpath,."
#INCLUDEPATH="-I. -I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include"
#LIBPATH="-L. -L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib -Wl,-rpath,."

set -x

# go mod init main
# go mod tidy
${CC} -c ${CFLAGS} ${INCLUDEPATH} -fPIC gauss.c
if [ $? -ne 0 ]; then exit 1; fi
${CC} ${LIBPATH} -shared -o libgauss.so gauss.o -lgmp -fopenmp
if [ $? -ne 0 ]; then exit 1; fi
# go build -buildmode=c-shared -o libcgolu.so .
# if [ $? -ne 0 ]; then exit 1; fi
${CC} -c ${CFLAGS} ${INCLUDEPATH} forkjoin.c -fopenmp
if [ $? -ne 0 ]; then exit 1; fi
${CC} -c ${CFLAGS} ${INCLUDEPATH} cmain.c -fopenmp
if [ $? -ne 0 ]; then exit 1; fi
#cc ${LIBPATH} -o cmain.o forkjoin.o -lcgolu -o cmain
${CC} ${LIBPATH} -o cmain cmain.o forkjoin.o -lgauss -fopenmp
