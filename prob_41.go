package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	largest := [10]int{}
	//N := 987654321
	NN := [2]int{4321, 7654321}
	for u := 0; u < len(NN); u++ {
		N := NN[u]
		for i := N; i > N/10; i -= 2 {
			digit := int(math.Log10(float64(i))) + 1
			i_str := strconv.Itoa(i)
			marker := [10]bool{}
			j := 0
			s_len := len(i_str)
			for k := 3; k < int(math.Sqrt(float64(i)))+1; k += 2 {
				if i%k == 0 {
					goto END
				}
			}
			for ; j < s_len; j++ {
				midx := i_str[j] - '0'
				if int(midx) > digit || marker[midx] || midx == 0 {
					goto END
				}
				marker[midx] = true
			}
			fmt.Println(i, digit, largest)

			if j == s_len {
				largest[digit] = i
				i = 0
			}
		END:
		}
	}
	fmt.Println(largest)
}
