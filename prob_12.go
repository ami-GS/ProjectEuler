package main

import (
	"fmt"
	"math"
)

const LIMIT int = 500

func main() {
	fac_nums := make([]int, 1)
	fac_nums[0] = 1
	tri_num := 1

	for n := 2; len(fac_nums) < 500; n++ {
		fac_nums = make([]int, 0)
		tri_num += n //start from 3
		for j := 2; j < int(math.Sqrt(float64(tri_num)))+1; j++ {
			//fmt.Println(tri_num, j)
			if tri_num%j == 0 {
				fac_nums = append(fac_nums, j)
			}
		}
		for j := 0; j < len(fac_nums); j++ {
			val := tri_num / fac_nums[j]
			ok := 0
			for k := 0; k < len(fac_nums); k++ { // check existence
				if val == fac_nums[k] {
					ok = 1
					break
				}
			}

			if ok == 0 {
				fac_nums = append(fac_nums, val)
			}
		}
		fmt.Println(tri_num, len(fac_nums), n)
	}

}
