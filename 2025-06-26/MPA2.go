package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func MakeNum(DigitA string, DigitB string) ([]int, []int) {
	var NumListA, NumListB []int
	DA, _ := strconv.Atoi(DigitA)
	DB, _ := strconv.Atoi(DigitB)
	for i := 0; i < DA; i++ {
		NumListA = append(NumListA, rand.Intn(10))
	}
	for i := 0; i < DB; i++ {
		NumListB = append(NumListB, rand.Intn(10))
	}
	fmt.Println(NumListA, NumListB)
	return NumListA, NumListB
}

func ListPlus(NumListA []int, NumListB []int) []int {
	lA := len(NumListA)
	lB := len(NumListB)
	if lB > lA {
		for i := 0; i < lB-lA; i++ {
			NumListA = append([]int{0}, NumListA...)
		}
	}
	lA = len(NumListA)
	if lA > lB {
		for i := 0; i < lA-lB; i++ {
			NumListB = append([]int{0}, NumListB...)
		}
	}
	lB = len(NumListB)
	var AnsListP1 []int
	for i := 1; i <= lA; i++ {
		AnsListP1 = append(AnsListP1, NumListA[lA-i]+NumListB[lB-i])
	}
	return AnsListP1
}

func AnsListPlus(ListP1 []int) []int {
	var ListP2, AnsListP []int
	var q int
	l1 := len(ListP1)
	for i := 0; i < l1; i++ {
		ListP2 = append(ListP2, (q+ListP1[i])%10)
		q = (q + ListP1[i]) / 10
	}
	for {
		if q > 0 {
			ListP2 = append(ListP2, q%10)
			q = q / 10
		} else {
			break
		}
	}
	fmt.Println(ListP2)
	l2 := len(ListP2)
	for i := 1; i <= l2; i++ {
		AnsListP = append(AnsListP, ListP2[l2-i])
	}
	return AnsListP
}

func ListTiems(NumListA []int, NumListB []int) []int {
	var AnsListT1, SubListT []int
	lA := len(NumListA)
	lB := len(NumListB)
	for i := 1; i <= lB; i++ {
		for j := 1; j <= lA; j++ {
			SubListT = append([]int{NumListA[lA-j] * NumListB[lB-i]}, SubListT...)
		}
		//AnsListの各桁+SubListの各桁の和をAnsListに格納
		if i == 1 {
			for k := 0; k < len(SubListT); k++ {
				AnsListT1 = append(AnsListT1, 0)
			}
		} else {
			for k := 0; k < i-1; k++ {
				SubListT = append(SubListT, 0)
			}
		}
		fmt.Println(AnsListT1, "\n", SubListT)
		AnsListT1 = AnsListPlus(ListPlus(AnsListT1, SubListT))
		fmt.Println("a", AnsListT1)
		SubListT = nil
	}
	return AnsListT1
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		print("(桁数)(演算子)(桁数)：")
		scanner.Scan()
		if scanner.Text() == "" {
			break
		} else {
			BoolPlus := strings.Contains(scanner.Text(), "+")
			if BoolPlus == true {
				SplitPlus := strings.Split(scanner.Text(), "+")
				NumListA, NumListB := MakeNum(SplitPlus[0], SplitPlus[1])
				ListP1 := ListPlus(NumListA, NumListB)
				fmt.Println(ListP1)
				AnsListP := AnsListPlus(ListP1)
				fmt.Println("Answer:", AnsListP, "\n")
			}

			BoolTimes := strings.Contains(scanner.Text(), "*")
			if BoolTimes == true {
				SplitTimes := strings.Split(scanner.Text(), "*")
				NumListA, NumListB := MakeNum(SplitTimes[0], SplitTimes[1])
				ListT1 := ListTiems(NumListA, NumListB)
				//fmt.Println(ListT1)
				//AnsListT := AnsListPlus(ListT1)
				fmt.Println("Answer:", ListT1, "\n")
			}
		}
	}
}
