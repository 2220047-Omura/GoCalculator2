package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
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
		fmt.Println("insert：", n)
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

func (f *fib) printTreeGo() {
	wg.Done()
	wg.Wait()
	fmt.Println(f)
	if f.left != nil {
		wg.Add(1)
		go f.left.printTreeGo()
	}
	if f.right != nil {
		wg.Add(1)
		go f.right.printTreeGo()
	}
}

func (f *fib) nodeCalc() *fib {
	fmt.Println("nodeCalc")
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
			} else {
				f = f.fibTree(n)
				f.printTree()
				fmt.Println("printTreeGo")
				wg.Add(1)
				f.printTreeGo()
				/*
					for {
						f = f.nodeCalc()
						if contFib == false {
							break
						}
					}
					f.printTree()
				*/
			}
		}
		l = 0
	}
}
