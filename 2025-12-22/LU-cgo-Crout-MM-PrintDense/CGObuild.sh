
#!/bin/sh

INCLUDEPATH="-I/opt/homebrew/include"
LIBPATH="-L. -L/opt/homebrew/lib -Wl,-rpath,/opt/homebrew/lib -Wl,-rpath,."

set -x

go mod init main
go mod tidy
gcc -c ${CFLAGS} ${INCLUDEPATH} -fPIC crout.c
if [ $? -ne 0 ]; then exit 1; fi
gcc ${LIBPATH} -shared -o libcrout.so crout.o -lmpfi -lmpfr -lgmp
if [ $? -ne 0 ]; then exit 1; fi
#go build .
#go build -buildmode=c-shared -o liblu.so *.go
go build -buildmode=c-shared -o libcgocrout.so .
if [ $? -ne 0 ]; then exit 1; fi
#cc -c ${CFLAGS} ${INCLUDEPATH} -fPIC cmain.c
cc -c ${CFLAGS} ${INCLUDEPATH} cgomain.c
if [ $? -ne 0 ]; then exit 1; fi
cc cgomain.o ${LIBPATH} -lcgocrout -o cgomain
if [ $? -ne 0 ]; then exit 1; fi

