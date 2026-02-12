package main

// #cgo CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
//// #cgo CFLAGS: -I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include
//// #cgo LDFLAGS: -L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib -lmpfi -lmpfr -lgmp
//
// #include "skyline.h"
//
import "C"

import (
	"fmt"
	"sync"
)

var E int
var sgo int

//export defE
func defE(e int) {
	E = e
}


var isk []int

//export makeIsk
func makeIsk(n int) {
	isk = append(isk, n)
}

var wg sync.WaitGroup

func UcallWG(m int, l int) {
	defer wg.Done()
	C.Usetsk(C.int(m),C.int(l))
}

//export forkjoin
func forkjoin(a int, l int) {
	for m := l; m < E; m++ {
		if isk[m] == a {
			wg.Add(1)
			go UcallWG(m, l)
			sgo += 1
		}
	}
	wg.Wait()
	fmt.Println("row:",a, "S:",sgo)
	sgo = 0
}

func main() {

	fmt.Println("main in golang, which must not be called!")

}
