package main

import (
	"fmt"
)

func fact(seed int) int {
	ans := 1
	for i := 2; i < seed+1; i++ {
		ans *= i
	}
	return ans
}

func main() {
	list := [2][10]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	num := 1000000
	for i := 9; i >= 0; i-- {
		tmp := fact(i)
		for j := 0; j < i+1; j++ {
			if num-tmp <= 0 {
				l := 0
				for k := 0; k < j+l+1; k++ {
					if list[1][k] == 1 {
						l++
					}
				}
				list[1][j+l] = 1
				fmt.Println(j, list[0][j+l], list, num)
				break
			}
			num -= tmp
		}
	}

}
