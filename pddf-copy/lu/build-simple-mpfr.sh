#!/bin/sh

TITLE="[simple-mpfr]"
echo ${TITLE}

CC=gcc

TARGET=simple-mpfr
SRC=${TARGET}.c
OBJ=${TARGET}.o

set -x 

CC=gcc

${CC} -O3 -g -Wall \
 -I/usr/local/GMP-6.3.0/include \
 -I/usr/local/MPFR-4.2.2/include \
 -c ${SRC} -o ${OBJ} \
 -pthread


${CC} -O3 -g -Wall \
 -L/usr/local/GMP-6.3.0/lib \
 -L/usr/local/MPFR-4.2.2/lib \
 ${OBJ} -o ${TARGET} \
 -lmpfr -lgmp -pthread \
 -static

# -lgmp -lmpfr -pthread \

