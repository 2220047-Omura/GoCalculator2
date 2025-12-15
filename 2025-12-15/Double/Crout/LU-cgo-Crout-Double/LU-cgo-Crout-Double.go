package main

// #cgo CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/opt/homebrew/lib  -lmpfi -lmpfr -lgmp
//
//// #cgo CFLAGS: -I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include
//// #cgo LDFLAGS: -L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib -lmpfi -lmpfr -lgmp
//
// #include "crout.h"
//
import "C"

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Ucall(i int, j int) {
	C.Uset(C.int(i), C.int(j))
}

func Lcall(j int, i int) {
	C.Lset(C.int(j), C.int(i))
}

func UcallWG(i int, j int) {
	defer wg.Done()
	C.Uset(C.int(i), C.int(j))
}

func LcallWG(j int, i int) {
	defer wg.Done()
	C.Lset(C.int(j), C.int(i))
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
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			//fmt.Println("Uset", i, j)
			Ucall(i, j)
		}
		for j := i + 1; j < N; j++ {
			//fmt.Println("Lset", j, i)
			Lcall(j, i)
		}
	}
	te = time.Now()
	fmt.Println("逐次：", te.Sub(ts), "\n")
	//C.comp()
	C.printMatrix3()

	//fmt.Println("-----並列-----")
	C.reset()

	ts = time.Now()
	for i := 0; i < N; i++ {
		wg.Add(N - i)
		for j := i; j < N; j++ {
			go UcallWG(i, j)
		}
		wg.Wait()
		wg.Add(N - i - 1)
		for j := i + 1; j < N; j++ {
			go LcallWG(j, i)
		}
		wg.Wait()
	}
	te = time.Now()
	fmt.Println("並列：", te.Sub(ts), "\n")
	//	C.comp()
	C.printMatrix3()
}
