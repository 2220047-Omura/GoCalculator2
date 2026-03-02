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
var Ngo int
var s int

//var cgo int32

//export defE
func defE(e int) {
	E = e
}

var isk []int

var prof []int

var Dia []int

//export makeIsk
func makeIsk(n int) {
	isk = append(isk, n)
}

//export makeProf
func makeProf(n int) {
	prof = append(prof, n)
}

var wg sync.WaitGroup

func UcallWG(m int, l int) {
	defer wg.Done()
	C.Usetsk(C.int(m), C.int(l))
}

//export forkjoinCount
func forkjoinCount(a int, l int) {
	//C.cleanCountS()
	for m := l; m < E; m++ {
		if isk[m] == a {
			/*
			profM := C.getprof(C.int(m))
			profL := C.getprof(C.int(l))
			s = min(int(profM), int(profL))
			if s != 0 {
				//wg.Add(1)
				//go UcallWG(m, l)
				C.Usetsk(C.int(m), C.int(l))
				Ngo += 1
			}
			*/
			C.Usetsk(C.int(m),C.int(l))
			Ngo += 1
		}
	}
	//wg.Wait()
	sum := C.getS(C.int(a))
	sumSQ := C.getS2(C.int(a))
	//fmt.Println("row:", a, "S:", sgo, "cgo:", cgo, "cgo2:",cgo2)
	//fmt.Println("row:", a, "Ngo:", Ngo, "ave:", float64(sum)/float64(Ngo), "var:", (float64(sumSQ)/float64(Ngo))-(float64(sum)/float64(Ngo))*(float64(sum)/float64(Ngo)))
	fmt.Println(a, ", ", Ngo, ", ", float64(sum)/float64(Ngo), ", ", (float64(sumSQ)/float64(Ngo))-(float64(sum)/float64(Ngo))*(float64(sum)/float64(Ngo)))
	Ngo = 0
}

//export forkjoin
func forkjoin(a int, l int, E2 int) {
	for m := l; m < E2; m++ {
		if isk[m] == a {
			wg.Add(1)
			go UcallWG(m, l)
		}
	}
	wg.Wait()
}

var profM int
var profL int

//export forkjoin2
func forkjoin2(a int, l int, E2 int) {
	for m := l; m < E2; m++ {
		if isk[m] == a {
			//profM := C.getprof(C.int(m))
			//profL := C.getprof(C.int(l))
			//profM = prof[m]
			//profL = prof[l]
			//s = min(int(profM), int(profL))
			if (int(C.getLength(C.int(m), C.int(l))) != 0) {
				wg.Add(1)
				go UcallWG(m, l)
			}
		}
	}
	wg.Wait()
}

func main() {

	fmt.Println("main in golang, which must not be called!")

}
