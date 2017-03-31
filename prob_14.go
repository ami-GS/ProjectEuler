package main

import (
	"fmt"
)

const N = 1000000

func Getlen(n int, memo map[int]int) int {
	val, ok := memo[n]
	if ok {
		return val
	} else if n == 1 {
		return 1
	} else if n%2 == 0 {
		return 1 + Getlen(n/2, memo)
	} else {
		return 1 + Getlen(3*n+1, memo)
	}
}

func main() {
	memo := make(map[int]int)
	maxlen := 0
	ans := 0
	for i := 1; i < N+1; i++ {
		tmp := Getlen(i, memo)
		if tmp > maxlen {
			maxlen = tmp
			ans = i
		}
	}
	fmt.Println(maxlen, ans)
}
