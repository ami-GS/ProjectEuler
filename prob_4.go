package main

import (
	"fmt"
	"strconv"
)

func main() {
	largest := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			var flag bool = true
			tmp := i * j
			str := strconv.Itoa(tmp)
			l := len(str)
			for k := 0; k < l/2; k++ {
				if str[k] != str[l-k-1] {
					flag = false
					break
				}
			}
			if flag && largest < tmp {
				largest = tmp
			}
		}
	}
	fmt.Println(largest)
}
