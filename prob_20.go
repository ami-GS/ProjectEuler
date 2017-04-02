package main

import (
	"fmt"
)

func main() {
	myInt := make([]int, 50)
	myInt[0] = 1
	idxMax := 0
	for i := 2; i <= 100; i++ {
		up := make([]int, idxMax+2)
		for j := 0; j < idxMax+1; j++ {
			myInt[j] *= i
			if myInt[j] >= 1000000000 {
				up[j] = myInt[j] / 1000000000
				myInt[j] %= 1000000000
			}
		}
		for j := 0; j < idxMax+1; j++ {
			myInt[j+1] += up[j]
		}
		if myInt[idxMax+1] != 0 {
			idxMax++
		}
	}
	ans := 0
	for i := 0; i < idxMax+1; i++ {
		if myInt[i] == 0 {
			continue
		}
		for j := 100000000; j >= 1; j /= 10 {
			ans += myInt[i] / j
			myInt[i] %= j
		}
	}
	fmt.Println(ans)
}
