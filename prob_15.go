package main

import (
	"fmt"
)

func main() {
	a := 1
	//list := [20]bool{}
	lIdx := 20
	for i := 40; i >= 21; i-- {
		a *= i
		for ; lIdx != 0 && a%lIdx == 0; lIdx-- {
			a /= lIdx
		}
	}
	fmt.Println(a)
}
