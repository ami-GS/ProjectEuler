package main

import (
	"fmt"
)

func main() {
	var prv, cur, ans, tmp int = 1, 2, 0, 0
	for prv < 4000000 {
		tmp = cur
		cur += prv
		prv = tmp
		if prv%2 == 0 {
			ans += prv
		}
	}
	fmt.Println(ans)
}
