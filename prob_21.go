package main

import (
	"fmt"
	"math"
)

func main() {
	ans := 0
	N := 10000
	sums := map[int]int{}

	for i := 2; i < N+1; i++ {
		sum := 0
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				sum += j + i/j
			}
		}
		if i%int(math.Sqrt(float64(i))) == 0 {
			sum += i
		}
		sums[i] = sum + 1
	}
	for a := 2; a < N+1; a++ {
		val_b, ok_b := sums[a]
		if ok_b && val_b > 0 {
			tmp, ok_a := sums[val_b]
			if ok_a && tmp > 0 {
				if a == tmp && tmp != val_b {
					fmt.Println(a, val_b)
					ans += val_b + tmp
					sums[val_b] = -sums[val_b]
					sums[a] = -sums[a]
				}
			}
		}
	}
	fmt.Println(ans)
}
