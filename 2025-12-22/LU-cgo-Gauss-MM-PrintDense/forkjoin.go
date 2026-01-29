package main

// #cgo CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
//// #cgo CFLAGS: -I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include
//// #cgo LDFLAGS: -L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib -lmpfi -lmpfr -lgmp
//
// #include "gauss.h"
//
import "C"

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//export call1
func call1(k int, i int, N int) {
	C.LUfact1(C.int(k), C.int(i))
	for j := k + 1; j < N; j++ {
		call2(k, i, j)
	}
}

func call2(k int, i int, j int) {
	C.LUfact2(C.int(k), C.int(i), C.int(j))
}

func call3(k int, i int, N int) {
	defer wg.Done()
	C.LUfact1(C.int(k), C.int(i))
	for j := k + 1; j < N; j++ {
		call2(k, i, j)
	}
}

//export forkjoin
func forkjoin(k int, N int) {
	wg.Add(N - (k + 1))
	for i := k + 1; i < N; i++ {
		go call3(k, i, N)
	}
	wg.Wait()
}

func main() {

	fmt.Println("main in golang, which must not be called!")

}
