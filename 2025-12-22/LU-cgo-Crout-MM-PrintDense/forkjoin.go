package main

// #cgo CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
//// #cgo CFLAGS: -I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include
//// #cgo LDFLAGS: -L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib -lmpfi -lmpfr -lgmp
//
// #include "crout.h"
//
import "C"

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func UcallWG(i int, j int, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	C.Uset(C.int(i), C.int(j))
}

func LcallWG(j int, i int, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	C.Lset(C.int(j), C.int(i))
}

//export forkjoin
func forkjoin(i int, N int) {
	wg.Add(N - i)
	for j := i; j < N; j++ {
		go UcallWG(i, j, &wg)
	}
	wg.Wait()
	wg.Add(N - i - 1)
	for j := i + 1; j < N; j++ {
		go LcallWG(j, i, &wg)
	}
	wg.Wait()
}

func main() {

	fmt.Println("main in golang, which must not be called!")

}
