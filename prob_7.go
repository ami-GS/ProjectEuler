package main

import (
	"fmt"
	"math"
)

func main() {
	num := 2
	ans := 0
	for i := 5; num != 10001; i += 2 {
		var flag bool = true
		for j := 3; j < int(math.Sqrt(float64(i)))+1; j += 2 {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			num++
			ans = i
		}
	}
	fmt.Println(ans)
}
