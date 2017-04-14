package main

import (
	"fmt"
)

func main() {
	ans := 1
	last_ru := 1

	for i := 1; i < 501; i++ {
		rd := last_ru + 2*i
		ans += rd
		for j := 1; j < 4; j++ {
			fmt.Println(rd)
			rd += 2 * i
			ans += rd
		}
		fmt.Println(" ", rd)
		last_ru = rd
	}
	fmt.Println(ans)
}
