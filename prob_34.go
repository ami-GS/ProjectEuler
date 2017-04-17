package main

import (
	"fmt"
	"strconv"
)

func main() {
	N := 10000000
	ans := 0
	memo := [10]int{}
	memo[0] = 1
	memo[1] = 1
	for i := 2; i < 10; i++ {
		memo[i] = memo[i-1] * i
	}
	fmt.Println(ans, memo)

	for i := 3; i < N; i++ {
		is := strconv.Itoa(i)
		tmp := 0
		for j := 0; j < len(is); j++ {
			tmp += memo[is[j]-'0']
		}
		if tmp == i {
			ans += i
			fmt.Println(i)
		}
	}
	fmt.Println(ans)
}
