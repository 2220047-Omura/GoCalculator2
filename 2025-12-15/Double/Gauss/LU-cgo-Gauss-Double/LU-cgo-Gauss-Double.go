package main

// #cgo CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/opt/homebrew/lib  -lmpfi -lmpfr -lgmp
//
//// #cgo CFLAGS: -I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include
//// #cgo LDFLAGS: -L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib -lmpfi -lmpfr -lgmp
//
// #include "gauss.h"
//
import "C"

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup


func call1(k int, i int, N int) {

	C.LUfact1(C.int(k), C.int(i))

	for j := k + 1; j < N; j++ {
		call2(k, i, j)
	}
}

func call2(k int, i int, j int) {

	C.LUfact2(C.int(k), C.int(i), C.int(j))
}

func call3(k int, i int, N int, wg *sync.WaitGroup) {
	defer wg.Done()

	C.LUfact1(C.int(k), C.int(i))

	for j := k + 1; j < N; j++ {
		call2(k, i, j)
	}
}

func main() {
	//fmt.Println("【クラウト法】")
	//C.init()
	N := int(C.def())
	var ts, te time.Time

	//fmt.Println("-----逐次-----")
	C.reset()
	//SimpleA(&A)
	ts = time.Now()
	for k := 0; k < N; k++ {
		for i := k + 1; i < N; i++ {
			//fmt.Println(k, i)
			call1(k, i, N)
		}
	}
	te = time.Now()
	fmt.Println("逐次：", te.Sub(ts), "\n")
	//	C.comp()
	C.printMatrix3()

	//fmt.Println("-----並列-----")
	C.reset()

	ts = time.Now()
	for k := 0; k < N; k++ {
		for i := k + 1; i < N; i++ {
			//fmt.Println(k, i)
			wg.Add(1)
			go call3(k, i, N, &wg)
		}
		wg.Wait()
	}
	te = time.Now()
	fmt.Println("並列：", te.Sub(ts), "\n")
	//	C.comp()
	C.printMatrix3()
}
