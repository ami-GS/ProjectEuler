package main

import (
	"fmt"
	"math"
)

func is_prime(d int) bool {
	if d <= 2 {
		return false
	}
	for i := 3; i < int(math.Sqrt(float64(d)))+1; i += 2 {
		if d%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	A := 1000
	B := 1000

	ans := 0
	ansA := 0
	ansB := 0
	ansN := 0

	for a := -A + 1; a < A; a++ {
		for b := -B + 1; b < B; b += 2 {
			if !is_prime(b) {
				continue
			}
			nSeq := 0
			for n := 0; n != b; n++ {
				p := n*n + a*n + b
				if (n%2 == 1 && a%2 == 0) || !is_prime(p) {
					break
				}
				nSeq++
			}
			if ansN < nSeq {
				ansN = nSeq
				ansA = a
				ansB = b
				ans = ansA * ansB
			}
		}
	}
	fmt.Println(ans, ansA, ansB, ansN)
}
