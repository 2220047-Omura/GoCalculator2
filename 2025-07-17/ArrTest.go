package main

import (
	"fmt"
)

type Lst struct {
	c    chan float64
	flag bool
}

func main() {
	var Larr [3][3]Lst
	fmt.Println(Larr)
	fmt.Println(Larr[1][1].flag)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i < j {
				Larr[j][i].flag = true
			}
		}
	}

	fmt.Println(Larr)
	fmt.Println(Larr[2][1].flag, Larr[1][1].flag, Larr[2][2].flag)
}
