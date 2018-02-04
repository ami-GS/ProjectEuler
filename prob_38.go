package main

import (
	"fmt"
	"strconv"
)

func calc() {
	//startFrom := 91
	starts := []int{91, 9182}
	ends := []int{100, 10000}
	ans := 0
	for i, base := range starts {
		for num := base; num < ends[i]; num++ {
			numStr := ""
			for j := 1; ; j++ {
				numStr += strconv.Itoa(num * j)
				if len(numStr) == 9 {
					flag := [9]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0}
					tmp := []byte(numStr)
					for k := 0; k < 9; k++ {
						if tmp[k] == '0' || flag[tmp[k]-'1'] == 1 {
							goto OUT
						}
						flag[tmp[k]-'1'] = 1
					}
					concatNum, _ := strconv.Atoi(numStr)
					if concatNum > ans {
						ans = concatNum
					}
				} else if len(numStr) > 9 {
					break
				}
			}
		OUT:
		}
	}
	fmt.Println(ans)
}

func main() {
	calc()
}
