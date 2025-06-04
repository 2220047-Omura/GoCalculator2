package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stack, result, nilStack []string

var calc, nilCalc []int

var inputN bool

func judge(j string) []string {
	n, err := strconv.Atoi(j)
	if err == nil {
		n += 0
		inputN = true
	} else {
		inputN = false
	}
	for {
		l := len(stack) - 1
		if stack[l] == "s" && j == "=" {
			break
		} else if stack[l] == "(" && j == ")" {
			stack = append(nilStack, stack[:l]...)
			fmt.Println("stack: ", stack)
			break
		} else if (stack[l] == "s") ||
			(stack[l] == "(" && (j == "+" || j == "-")) || ((stack[l] == "+" || stack[l] == "-") && (j == "*" || j == "/")) ||
			(stack[l] == "(" && (j == "*" || j == "/")) || (j == "(") || (inputN == true) {
			stack = append(stack, j)
			fmt.Println("stack: ", stack)
			break
		} else {
			result = append(result, stack[l])
			fmt.Println("result:", result)
			stack = append(nilStack, stack[:l]...)
			fmt.Println("stack: ", stack)
		}
	}
	return result
}

func calcResult(i int, str string) []int {
	l := len(calc) - 1
	if str == "" {
		calc = append(calc, i)
	} else {
		a := calc[l-1]
		b := calc[l]

		switch {
		case str == "+":
			c := a + b
			calc = append(nilCalc, calc[:l-1]...)
			calc = append(calc, c)
		case str == "-":
			c := a - b
			calc = append(nilCalc, calc[:l-1]...)
			calc = append(calc, c)
		case str == "*":
			c := a * b
			calc = append(nilCalc, calc[:l-1]...)
			calc = append(calc, c)
		case str == "/":
			c := a / b
			calc = append(nilCalc, calc[:l-1]...)
			calc = append(calc, c)
		}
	}
	return calc
}

func main() {
	for {
		stack = append(stack, "s")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "" {
			break
		} else {
			Split := strings.Split(scanner.Text(), " ")
			for _, Split := range Split {
				fmt.Println(Split)
				judge(Split)
			}
			fmt.Println("---------")
			fmt.Println(result)
			fmt.Println("calcResult")
			for _, str := range result {
				n, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println(calcResult(0, str))
				} else {
					fmt.Println(calcResult(n, ""))
				}
			}
			stack = nil
			result = nil
			calc = nil
		}
	}
}
