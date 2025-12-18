#!/bin/sh

CFLAGS="-O3 -Wall"

INCLUDEPATH="-I. -I/opt/homebrew/include"
#LIBPATH="-L. -L/opt/homebrew/lib -Wl,-rpath,/opt/homebrew/lib:."
LIBPATH="-L. -L/opt/homebrew/lib -Wl,-rpath,/opt/homebrew/lib -Wl,-rpath,."
#INCLUDEPATH="-I. -I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include"
#LIBPATH="-L. -L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib -Wl,-rpath,."

set -x

go mod init main
go mod tidy
gcc -c ${CFLAGS} ${INCLUDEPATH} -fPIC skyline.c
if [ $? -ne 0 ]; then exit 1; fi
gcc ${LIBPATH} -shared -o libskyline.so skyline.o -lmpfi -lmpfr -lgmp
if [ $? -ne 0 ]; then exit 1; fi
#go build .
#go build -buildmode=c-shared -o libcrout.so *.go
go build -buildmode=c-shared -o libcgoskyline.so .
if [ $? -ne 0 ]; then exit 1; fi
#cc -c ${CFLAGS} ${INCLUDEPATH} -fPIC cmain.c
cc -c ${CFLAGS} ${INCLUDEPATH} cmain.c
if [ $? -ne 0 ]; then exit 1; fi
cc ${LIBPATH} -o cmain.o -lcgoskyline -o cgomain
