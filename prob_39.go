package main

import (
	"fmt"
)

func main() {

	lists := map[int][][3]int{}
	maxnum := 0
	maxp := 0
	for p := 3; p <= 1000; p++ {
		li := [][3]int{}
		for a := 1; a < p; a++ {
			for b := a; ; b++ {
				c := p - a - b
				AA := a * a
				BB := b * b
				CC := c * c
				if AA+BB == CC {
					li = append(li, [3]int{a, b, c})
				} else if BB > CC-AA {
					break
				}
			}
		}
		lilen := len(li)
		if lilen != 0 {
			lists[p] = li
			if maxnum < lilen {
				maxnum = lilen
				maxp = p
			}
		}
	}
	fmt.Println(maxp, maxnum)
}
