package main

import (
	"fmt"
	"math"
	"strconv"
)

func calc() {
	ans := 0
	for i := 1; i < 1000000; i += 2 {
		pow := int(math.Log10(float64(i)))
		if (i/int(math.Pow(10, float64(pow))))%2 == 0 {
			i += int(math.Pow(10, float64(pow)))
		}

		missDec := false
		for j := 0; j < pow; j++ {
			lm := (i / int(math.Pow(10, float64(pow-j)))) % 10
			rm := (i / int(math.Pow(10, float64(j)))) % 10
			if lm != rm {
				missDec = true
				break
			}
		}
		if missDec {
			continue
		}

		missBin := false
		iBin := strconv.FormatInt(int64(i), 2)
		for j := 0; j < len(iBin)/2; j++ {
			if iBin[len(iBin)-j-1] != iBin[j] {
				missBin = true
				break
			}
		}
		if missBin {
			continue
		}
		fmt.Println(i, iBin)
		ans += i
	}
	fmt.Println(ans)
}

func calc2() {
	// use strconv
	ans := 0
	for i := 1; i < 1000000; i += 2 {
		iStr := strconv.Itoa(i)

		missDec := false
		for j := 0; j < len(iStr); j++ {
			if iStr[j] != iStr[len(iStr)-1-j] {
				missDec = true
				break
			}
		}
		if missDec {
			continue
		}

		missBin := false
		iBin := strconv.FormatInt(int64(i), 2)
		for j := 0; j < len(iBin)/2; j++ {
			if iBin[len(iBin)-j-1] != iBin[j] {
				missBin = true
				break
			}
		}
		if missBin {
			continue
		}
		fmt.Println(i, iBin)
		ans += i
	}
	fmt.Println(ans)
}

func main() {
	calc()
	calc2()
}
