package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MakeListP(NumListA []string, NumListB []string) []int {
	nA := len(NumListA)
	nB := len(NumListB)
	var AnsListP1 []int
	if nA >= nB {
		for i := nA; i > 0; i-- {
			if nB <= 0 {
				//fmt.Println(NumListA[nA-1], 0)
				AnsA, _ := strconv.Atoi(NumListA[i-1])
				AnsListP1 = append(AnsListP1, AnsA)
			} else {
				//fmt.Println(NumListA[nA-1], NumListB[nB-1])
				AnsA, _ := strconv.Atoi(NumListA[i-1])
				AnsB, _ := strconv.Atoi(NumListB[nB-1])
				AnsListP1 = append(AnsListP1, AnsA+AnsB)
				nB--
			}
		}
	} else {
		for j := nB; j > 0; j-- {
			if nA <= 0 {
				//fmt.Println(0, NumListB[nB-1])
				AnsB, _ := strconv.Atoi(NumListA[j-1])
				AnsListP1 = append(AnsListP1, AnsB)
			} else {
				//fmt.Println(NumListA[nA-1], NumListB[nB-1])
				AnsA, _ := strconv.Atoi(NumListA[nA-1])
				AnsB, _ := strconv.Atoi(NumListB[j-1])
				AnsListP1 = append(AnsListP1, AnsA+AnsB)
				nA--
			}
		}
	}
	return AnsListP1
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
				NumListA := strings.Split(SplitPlus[0], "")
				NumListB := strings.Split(SplitPlus[1], "")
				AnsListP1 := MakeListP(NumListA, NumListB)
				fmt.Println(AnsListP1)

				var AnsListP2, AnsListP3 []int
				var q int
				l1 := len(AnsListP1)
				for i := 0; i < l1; i++ {
					AnsListP2 = append(AnsListP2, (q+AnsListP1[i])%10)
					q = (q + AnsListP1[i]) / 10
				}
				for {
					if q > 0 {
						AnsListP2 = append(AnsListP2, q%10)
						q = q / 10
					} else {
						break
					}
				}
				fmt.Println(AnsListP2)
				l2 := len(AnsListP2)
				for i := 1; i <= l2; i++ {
					AnsListP3 = append(AnsListP3, AnsListP2[l2-i])
				}
				fmt.Println("Answer:", AnsListP3, "\n")
			}

			/*
				BoolTimes := strings.Contains(scanner.Text(), "*")
				if BoolTimes == true {
					SplitTimes := strings.Split(scanner.Text(), "*")
					NumListA := strings.Split(SplitTimes[0], "")
					NumListB := strings.Split(SplitTimes[1], "")
					nA := len(NumListA)
					nB := len(NumListB)
					//fmt.Println(nA, nB)
					var AnsListT1 []int
					for i := 0; i < nB; i++ {
						var AnsListT2 []int
						AnsA, _ := strconv.Atoi(NumListA[i])
						for j := 0; j < nA; j++ {
							AnsB, _ := strconv.Atoi(NumListB[j])
							AnsListT2 = append(AnsListT2, AnsA*AnsB)
						}
						if i == 0{
							AnsListT1 = append(AnsListT1, AnsListT2...)
						}else{
							for k := 0; k < i; k ++{
								AnsListT2 = append(AnsListT2, 0)
							}
						}
					}
				}
			*/
		}
	}
}
