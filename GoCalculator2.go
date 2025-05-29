package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var slise []int //式中の数を一時的に格納するリスト

var contCalc bool //nodeCalcを先頭から繰り返し呼び出すためのbool値

var p int //演算子の優先度

type Formula struct {
	priority    int      //pの比較対象となる優先度
	operator    string   //演算子
	stateO      bool     //演算子の有無
	number      int      //数
	stateN      bool     //数の有無
	left, right *Formula //子ノード(二つ)
}

func (f *Formula) makeTree(str string) *Formula {
	int, err := strconv.Atoi(str)
	if err == nil { //数値だった場合
		slise = append(slise, int)
	} else { //演算子だった場合
		switch { //優先度の付与
		case str == "(":
			p += 2
		case str == ")":
			p -= 2
		case str == "*" || str == "/":
			p += 1
			f = f.insertOperator(str)
			p -= 1
		case str == "+" || str == "-":
			f = f.insertOperator(str)
		default:
			break
		}
		fmt.Println("inmakeAST：", f)
		if str != "(" && str != ")" && slise != nil {
			fmt.Println("slise：", slise)
			f = f.insertNum(slise[0])
			slise = append(slise[:0], slise[1:]...)
			fmt.Println("inmakeAST：", f)
			for {
				f = f.nodeCalc()
				if contCalc == false {
					break
				}
			}
		}
	}
	return f
}

func (f *Formula) insertOperator(o string) *Formula {
	if f == nil {
		fmt.Println("insertO")
		return &Formula{priority: p, operator: o, stateO: true}
	} else if f.priority < p {
		if f.left == nil {
			fmt.Println("LeftO")
			f.left = f.left.insertOperator(o)
		} else if f.left.stateN == false {
			fmt.Println("LeftO")
			f.left = f.left.insertOperator(o)
		} else {
			fmt.Println("RightO")
			f.right = f.right.insertOperator(o)
		}
	} else {
		fmt.Println("insertO")
		return &Formula{priority: p, operator: o, stateO: true, left: f}
	}
	return f
}

func (f *Formula) insertNum(n int) *Formula {
	if f == nil {
		fmt.Println("insertN")
		return &Formula{number: n, stateN: true}
	}
	if f.left == nil {
		fmt.Println("LeftN")
		f.left = f.left.insertNum(n)
	} else if f.left.stateN == false {
		fmt.Println("LeftN")
		f.left = f.left.insertNum(n)
	} else {
		fmt.Println("RightN")
		f.right = f.right.insertNum(n)
	}
	return f
}

func (f *Formula) nodeCalc() *Formula {
	fmt.Println("nodeCalc")
	contCalc = false
	if f.left != nil && f.right != nil {
		if f.left.stateN == true && f.right.stateN == true {
			switch {
			case f.operator == "+":
				contCalc = true
				return &Formula{number: f.left.number + f.right.number, stateN: true, left: nil, right: nil}
			case f.operator == "-":
				contCalc = true
				return &Formula{number: f.left.number - f.right.number, stateN: true, left: nil, right: nil}
			case f.operator == "*":
				contCalc = true
				return &Formula{number: f.left.number * f.right.number, stateN: true, left: nil, right: nil}
			case f.operator == "/":
				contCalc = true
				return &Formula{number: f.left.number / f.right.number, stateN: true, left: nil, right: nil}
			}
		}
	}
	if f.left != nil {
		if f.left.stateO == true || f.left.stateN == false {
			f.left = f.left.nodeCalc()
		}
	}
	if f.right != nil {
		if f.right.stateO == true || f.right.stateN == false {
			f.right = f.right.nodeCalc()
		}
	}
	return f
}

func main() {
	for {
		var f *Formula
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "" {
			break
		} else {
			Split := strings.Split(scanner.Text(), " ")
			for _, Split := range Split {
				fmt.Println(Split)
				f = f.makeTree(Split)
			}
			fmt.Println("Answer：", f.number)
		}
	}
}
