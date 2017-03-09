package main

import (
	"fmt"
	"math"
)

func main() {
	var A float64 = 0
	B := math.Pow(float64(101*50), 2)
	for i := 1; i <= 100; i++ {
		A += math.Pow(float64(i), 2)
	}
	fmt.Println(int(B - A))
}
