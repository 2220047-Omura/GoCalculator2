#!/bin/sh

INCLUDEPATH="-I/opt/homebrew/include"
LIBPATH="-L/opt/homebrew/lib -Wl,-rpath,/opt/homebrew/lib"

set -x

go mod init main
go mod tidy
gcc -c -O3 ${INCLUDEPATH} -fPIC LU-cgo-Gauss-Double.c
gcc ${LIBPATH} -shared -o libgauss.so LU-cgo-Gauss-Double.o -lmpfi -lmpfr -lgmp
go build .
