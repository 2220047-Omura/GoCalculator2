CC=gcc

CFLAGS=-O3
LDFLAGS=-lgmp
INCLUDEPATH=-I/opt/homebrew/include
LIBPATH=-L/opt/homebrew/lib -Wl,-rpath,/opt/homebrew/lib
#INCLUDEPATH=-I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include
#LIBPATH=-L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib

GOSRC=LU-cgo-Skyline-Double.go
LIBNAME=SkDouble
LIBHDR=${LIBNAME}.h
#LIBSRC=${LIBNAME}.c
LIBSRC=${GOSRC}.c
LIBOBJ=lib${LIBNAME}.so
#LIBTMP=${LIBNAME}.o
LIBTMP=${GOSRC}.o
EXE=main
GOMOD=go.mod

main: ${LIBOBJ} go.mod ${HDRS} ${GOSRC}
	go build .

${LIBOBJ}: ${LIBSRC} ${LIBHDR}
	${CC} ${CFLAGS} -c -O3 ${INCLUDEPATH} -fPIC ${LIBSRC}
	${CC} ${LIBPATH} -shared -o $@ ${LDFLAGS}

go.mod: ${GOSRC}
	rm -fr $@
	go mod init main
	go mod tidy

clean:
	rm -rf *.o ${EXE} ${LIBOBJ} ${LIBTMP} ${GOMOD}
