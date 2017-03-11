package main

import (
	"fmt"
	"math"
)

func main() {
	//	var list []bool
	N := 2000000
	list := make([]bool, N)
	list[2] = true
	ans := 2
	for i := 3; i < N; i += 2 {
		flag := true
		for j := 3; j < int(math.Sqrt(float64(i)))+1; j += 2 {
			l := list[j]
			if l && i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			list[i] = true
			ans += i
		}
	}
	fmt.Println(ans)

}
