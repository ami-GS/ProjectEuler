package main

import (
	"fmt"
)

func main() {
	ans := 0
	for i := 3; i <= 999; i += 3 {
		ans += i
	}
	for i := 5; i <= 995; i += 5 {
		if i%3 != 0 {
			ans += i
		}
	}
	fmt.Printf("%d\n", ans)
}
