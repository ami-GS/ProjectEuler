package main

import (
	"fmt"
	"math"
)

var N = 600851475143

func main() {
	SqrtN := int(math.Sqrt(float64(N)))
	if SqrtN%2 == 0 {
		SqrtN -= 1
	}
	list := make([]uint8, SqrtN)
	list[2] = 1
	list[3] = 1

	for i := 5; i < SqrtN; i += 2 {
		flag := 0
		for j := 3; j < SqrtN; j += 2 {
			if list[j] == 0 {
				continue
			}
			if (i % j) == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			list[i] = 1
		}
	}

	for i := SqrtN; i > 1; i -= 2 {
		if list[i] == 1 && N%i == 0 {
			fmt.Println(i)
			break
		}
	}
}
