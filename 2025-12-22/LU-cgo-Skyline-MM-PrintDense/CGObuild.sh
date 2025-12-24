
#!/bin/sh

INCLUDEPATH="-I/opt/homebrew/include"
LIBPATH="-L. -L/opt/homebrew/lib -Wl,-rpath,/opt/homebrew/lib -Wl,-rpath,."

set -x

go mod init main
go mod tidy
gcc -c ${CFLAGS} ${INCLUDEPATH} -fPIC skyline.c
if [ $? -ne 0 ]; then exit 1; fi
gcc ${LIBPATH} -shared -o libskyline.so skyline.o -lmpfi -lmpfr -lgmp
if [ $? -ne 0 ]; then exit 1; fi
#go build .
#go build -buildmode=c-shared -o liblu.so *.go
go build -buildmode=c-shared -o libcgoskyline.so .
if [ $? -ne 0 ]; then exit 1; fi
#cc -c ${CFLAGS} ${INCLUDEPATH} -fPIC cmain.c
cc -c ${CFLAGS} ${INCLUDEPATH} cgomain.c
if [ $? -ne 0 ]; then exit 1; fi
cc cgomain.o ${LIBPATH} -lcgoskyline -o cgomain
if [ $? -ne 0 ]; then exit 1; fi

