package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var contFib bool
var l int

type fib struct {
	number      int
	level       int
	calc        bool
	left, right *fib
}

func (f *fib) fibTree(n int) *fib {
	if f == nil {
		if n != 1 && n != 0 {
			f = &fib{number: n, level: l, calc: true}
		} else {
			f = &fib{number: n, level: l}
		}
		if n > 1 {
			l = f.level + 1
			f.left = f.left.fibTree(n - 1)
			l = f.level + 1
			f.right = f.right.fibTree(n - 2)
		}
	}
	return f
}

func (f *fib) printTree() {
	fmt.Println(f)
	if f.left != nil {
		f.left.printTree()
	}
	if f.right != nil {
		f.right.printTree()
	}
}

func (f *fib) printTreeGo(n *sync.WaitGroup, c chan int) {
	var wg sync.WaitGroup
	if f.left != nil && f.right != nil {
		c1 := make(chan int, 1)
		c2 := make(chan int, 1)
		if f.left != nil {
			wg.Add(1)
			go f.left.printTreeGo(&wg, c1)
		}
		if f.right != nil {
			wg.Add(1)
			go f.right.printTreeGo(&wg, c2)
		}
		wg.Wait()
		v1 := <-c1
		v2 := <-c2
		c <- v1 + v2
	} else {
		c <- f.number
	}
	//fmt.Println(f)
	n.Done()
}

func (f *fib) nodeCalc() *fib {
	contFib = false
	if f.left != nil && f.right != nil {
		if f.left.calc == false && f.right.calc == false {
			contFib = true
			return &fib{number: f.left.number + f.right.number, calc: false, left: nil, right: nil}
		}
	}
	if f.left != nil {
		if f.left.calc == true {
			f.left = f.left.nodeCalc()
		}
	}
	if f.right != nil {
		if f.right.calc == true {
			f.right = f.right.nodeCalc()
		}
	}
	return f
}

func main() {
	for {
		var f *fib
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "" {
			break
		} else {
			n, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("数値を入力してください")
				fmt.Println(n)
			} else {
				var wg sync.WaitGroup
				c := make(chan int, 1)
				f = f.fibTree(n)
				//f.printTree()
				fmt.Println("printTreeGo")
				wg.Add(1)
				t1 := time.Now()
				f.printTreeGo(&wg, c)
				wg.Wait()
				t2 := time.Now()
				v := <-c
				fmt.Println(v)
				fmt.Println(t2.Sub(t1))
				t3 := time.Now()
				for {
					f = f.nodeCalc()
					if contFib == false {
						break
					}
				}
				t4 := time.Now()
				fmt.Println(f.number)
				fmt.Println(t4.Sub(t3))
			}
		}
		l = 0
	}
}
