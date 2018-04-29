package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hasSameDigits(bef, af int) bool {
	befStr := strconv.Itoa(bef)
	afStr := strconv.Itoa(af)
	for _, v := range afStr {
		if !strings.Contains(befStr, string(v)) {
			return false
		}
	}
	return true
}

func run() {
	for i := 123456; i < 165432; i++ {
		for j := 2; j <= 6; j++ {
			if !hasSameDigits(i, i*j) {
				goto NEXT
			}
			if j == 6 {
				fmt.Println(i)
			}
		}
	NEXT:
	}

}

func main() {
	run()
}
