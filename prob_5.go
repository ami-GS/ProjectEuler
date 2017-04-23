package main

import (
	"fmt"
)

func main() {
	for i := 20; ; i += inc {
		j := 3
		for ; j <= 20; j++ {
			if i%j != 0 {
				break
			}
		}
		if j == 21 {
			fmt.Println(i)
			break
		}
	}
}
