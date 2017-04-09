package main

import (
	"fmt"
	"math"
)

func main() {
	N := 28123
	sums := make([]int, N+1)
	ans := 1
	for i := 2; i < N+1; i++ {
		tmp := 0
		sqrt := int(math.Sqrt(float64(i)))
		for j := 2; j <= sqrt; j++ {
			if i%j == 0 {
				tmp += j + i/j
			}
		}
		if sqrt*sqrt == i {
			tmp -= sqrt
		}
		sums[i] = tmp + 1

		ans += i
		for j := 0; j < i; j++ {
			if sums[j] > j && sums[i-j] > i-j {
				ans -= i
				break
			}
		}
	}
	fmt.Println(ans)
}
