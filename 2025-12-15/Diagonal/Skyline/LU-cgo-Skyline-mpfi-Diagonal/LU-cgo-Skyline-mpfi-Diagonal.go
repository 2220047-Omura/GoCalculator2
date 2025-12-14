package main

// #cgo CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/opt/homebrew/lib  -lmpfi -lmpfr -lgmp
//
//// #cgo CFLAGS: -I/usr/local/GMP-6.3.0/include -I/usr/local/MPFR-4.2.2/include -I/usr/local/MPFI-1.5.4/include
//// #cgo LDFLAGS: -L/usr/local/GMP-6.3.0/lib -L/usr/local/MPFR-4.2.2/lib -L/usr/local/MPFI-1.5.4/lib -Wl,-rpath,/usr/local/MPFI-1.5.4/lib -lmpfi -lmpfr -lgmp
//
// #include "SkMpfi.h"
//
import "C"

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func call(b int, i int, j int) {
	C.Usetsk(C.int(b), C.int(i), C.int(j))
}

func callWG(b int, i int, j int, wg *sync.WaitGroup) {
	defer wg.Done()
	C.Usetsk(C.int(b), C.int(i), C.int(j))
}

func main() {
	//fmt.Println("【スカイライン法】")
	var ts, te time.Time
	var i, j, c, isk int
	var wg sync.WaitGroup
	
	C.init()
	N := int(C.getN())
	//fmt.Println(N)

	//fmt.Println("-----逐次-----")
	C.reset()
	ts = time.Now()
	for a := 1; a < N; a++ {
		c = 1
		isk = int(C.getIsk(C.int(c)))
		for b := 1; b < N; b++ {
			i = c - (isk - b)
			j = c
			if i == a {
				//fmt.Println("(i, j)=", i, j)
				call(b, i, j)
			}
			if b == isk {
				c += 1
				isk = int(C.getIsk(C.int(c)))
			}
		}
	}
	te = time.Now()
	fmt.Println("逐次：", te.Sub(ts), "\n")
	C.result()
	//PrintArr(Ask)

	//fmt.Println("-----並列-----")
	C.reset()
	ts = time.Now()
	for a := 1; a < N; a++ {
		c = 1
		isk = int(C.getIsk(C.int(c)))
		for b := 1; b < N; b++ {
			i = c - (isk - b)
			j = c
			if i == a {
				//fmt.Println("(i, j)=", i, j)
				wg.Add(1)
				go callWG(b, i, j, &wg)
			}
			if b == isk {
				c += 1
				isk = int(C.getIsk(C.int(c)))
			}
		}
		wg.Wait()
	}
	te = time.Now()
	fmt.Println("並列：", te.Sub(ts), "\n")
	C.result()
	//PrintArr(Ask)

	C.clear()
}
