#!/bin/sh

TITLE="[simple]"
echo ${TITLE}

CC=gcc

TARGET=simple
SRC=${TARGET}.c
OBJ=${TARGET}.o

set -x 

${CC} \
 -O3 -g -Wall \
 -I/usr/local/GMP-6.3.0/include \
 -I/usr/local/MPFR-4.2.2/include \
 -I/usr/local/MPFI-1.5.4/include \
 -c ${SRC} -o ${OBJ}



${CC} \
 -O3 -g -Wall \
 -L/usr/local/GMP-6.3.0/lib \
 -L/usr/local/MPFR-4.2.2/lib \
 -L/usr/local/MPFI-1.5.4/lib \
 -Wl,-rpath,/usr/local/MPFI-1.5.4/lib \
 ${OBJ} -o ${TARGET} \
 -lmpfi -lmpfr -lgmp 


