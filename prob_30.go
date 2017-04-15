package main

import (
	"fmt"
)

func main() {
	ans := 0
	memo := [10]int{0, 1, //memo
		2 * 2 * 2 * 2 * 2,
		3 * 3 * 3 * 3 * 3,
		4 * 4 * 4 * 4 * 4,
		5 * 5 * 5 * 5 * 5,
		6 * 6 * 6 * 6 * 6,
		7 * 7 * 7 * 7 * 7,
		8 * 8 * 8 * 8 * 8,
		9 * 9 * 9 * 9 * 9}
	for e := 0; e < 10; e++ {
		e10 := e * 100000
		for i := 0; i < 10; i++ {
			tmp := e10 + i*10000
			for j := 0; j < 10; j++ {
				tmp10 := tmp + j*1000
				for k := 0; k < 10; k++ {
					tmp20 := tmp10 + k*100
					for l := 0; l < 10; l++ {
						tmp30 := tmp20 + l*10
						for m := 0; m < 10; m++ {
							if memo[e]+memo[i]+memo[j]+memo[k]+memo[l]+memo[m] == tmp30+m {
								ans += tmp30 + m
								fmt.Println(tmp30 + m)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
