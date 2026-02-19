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

var Ngo int
var sum int
var sumSQ int

var wg sync.WaitGroup

//export call1
func call1(k int, i int, N int) {
	C.LUfact1(C.int(k), C.int(i))
	for j := k + 1; j < N; j++ {
		C.LUfact2(C.int(k), C.int(i), C.int(j))
	}
	sum += N-(k+1)
	sumSQ += (N-(k+1))*(N-(k+1))
}

//export call1WG
func call1WG(k int, i int, N int) {
	defer wg.Done()
	C.LUfact1(C.int(k), C.int(i))
	for j := k + 1; j < N; j++ {
		C.LUfact2(C.int(k), C.int(i), C.int(j))
	}
}

//export forkjoinCount
func forkjoinCount(k int, N int) {
	//wg.Add(N - (k + 1))
	for i := k + 1; i < N; i++ {
		//go call1WG(k, i, N)
		call1(k, i, N)
		Ngo += 1
	}
	//wg.Wait()
	fmt.Println("row:", k, "Ngo:", Ngo, "ave:", float64(sum)/float64(Ngo), "var:", (float64(sumSQ)/float64(Ngo))-(float64(sum)/float64(Ngo))*(float64(sum)/float64(Ngo)))
	Ngo = 0
	sum = 0
	sumSQ = 0
}

//export forkjoin
func forkjoin(k int, N int) {
	wg.Add(N - (k + 1))
	for i := k + 1; i < N; i++ {
		go call1WG(k, i, N)
	}
	wg.Wait()
}

func main() {

	fmt.Println("main in golang, which must not be called!")

}
