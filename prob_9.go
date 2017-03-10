package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	num := 1000
	for a = 1; a < num; a++ {
		for b = a + 1; b < num+1; b++ {
			c = 1000 - a - b
			if c < b {
				break
			}
			if int(math.Pow(float64(a), 2)+math.Pow(float64(b), 2)) == int(math.Pow(float64(c), 2)) {
				fmt.Println(a, b, c, a*b*c)
				return
			}
		}
	}
}
