CC=gcc

CFLAGS=-O3
LDFLAGS=-lmpfi -lmpfr -lgmp
INCLUDEPATH=-I. -I/opt/homebrew/include
LIBPATH=-L. -L/opt/homebrew/lib -Wl,-rpath,/opt/homebrew/lib -Wl,-rpath,.
#INCLUDEPATH=-I. -I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include
#LIBPATH=-L. -L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib -Wl,-rpath,.

GOSRC=forkjoin.go
LIBNAME=crout
LIBHDR=${LIBNAME}.h
LIBSRC=${LIBNAME}.c
LIBOBJ=lib${LIBNAME}.so
LIBCGONAME=cgocrout
LIBCGOOBJ=lib${LIBCGONAME}.so
LIBCGOHDR=lib${LIBCGONAME}.h
LIBTMP=${LIBNAME}.o
EXETMP=${EXE}.o
EXE=cgomain
GOMOD=go.mod

cmain: ${LIBCGOOBJ} ${EXETMP}
	${CC} ${LIBPATH} -o $@ ${EXETMP} -l${LIBCGONAME}

${LIBOBJ}: ${LIBSRC} ${LIBHDR}
	${CC} ${CFLAGS} -c -O3 ${INCLUDEPATH} -fPIC ${LIBSRC}
	${CC} ${LIBPATH} -shared -o $@ ${LDFLAGS}

${LIBCGOOBJ}: ${LIBOBJ} go.mod
	go build -buildmode=c-shared -o $@ .

.c.o:
	${CC} ${CFLAGS} -c ${INCLUDEPATH} $<

go.mod: ${GOSRC}
	rm -fr $@
	go mod init main
	go mod tidy

clean:
	rm -rf *.o *.so ${EXE} ${LIBOBJ} ${GOMOD} ${LIBCGOHDR}
