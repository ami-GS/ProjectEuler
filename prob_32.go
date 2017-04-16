package main

import (
	"fmt"
	"math"
	"strconv"
)

func hasZero(i int) bool {
	ia := strconv.Itoa(i)
	for ii := 0; ii < len(ia); ii++ {
		if ia[ii] == '0' {
			return true
		}
	}
	return false
}

func hasSame(i int) bool {
	ia := strconv.Itoa(i)
	li := [10]bool{}
	for ii := 0; ii < len(ia); ii++ {
		tmp := ia[ii] - '0'
		if li[tmp] {
			return true
		}
		li[tmp] = true
	}
	return false
}

func digitList(i int, list *[10]bool) bool {
	ia := strconv.Itoa(i)
	for ii := 0; ii < len(ia); ii++ {
		tmp := ia[ii] - '0'
		if list[tmp] {
			return false
		}
		list[tmp] = true
	}
	return true
}

func main() {
	ans := 0

	for i := 1; i < 100; i++ {
		if i%10 == 0 || i%10 == i/10 {
			continue
		}
		idigit := int(math.Log10(float64(i))) + 1
		li := [10]bool{}
		I := i
		for ii := idigit - 1; ii >= 0; ii-- {
			li[(I / int(math.Pow(float64(10), float64(ii))))] = true
			I = I % 10
		}
		for j := 1; j < 10000; j++ {
			if hasZero(j) || hasSame(j) {
				continue
			}

			prod := i * j
			if hasZero(prod) || hasSame(prod) {
				continue
			}

			pdigit := int(math.Log10(float64(prod))) + 1
			jdigit := int(math.Log10(float64(j))) + 1
			if pdigit+idigit+jdigit == 9 {
				list := li
				if !digitList(j, &list) {
					continue

				}
				if !digitList(prod, &list) {
					continue
				}

				all := true
				for k := 1; k < 10; k++ {
					if !list[k] {
						all = false
						break
					}
				}
				if all {
					ans += prod
					fmt.Println(i, j, prod)
				}
			}
		}
	}
	fmt.Println(ans)
}
