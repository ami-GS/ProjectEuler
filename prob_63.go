package main

import (
	"fmt"
	"math"
)

func run() int {
	out := 0
	for i := 1; i < 10; i++ {
		val := uint64(i)
		out++ // for i**1
		for j := 2; j <= 22; j++ {
			val *= uint64(i)
			// 9**21 is missed due to cast error, could be solve by using strconv.Itoa
			if j == int(math.Log10(float64(val)))+1 {
				fmt.Println(i, j, val)
				out++
			}
		}
	}
	return out + 1 // includes 9**9
}

func main() {
	fmt.Println(run())
}
