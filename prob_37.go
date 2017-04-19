package main

import (
	"fmt"
	"math"
	"strconv"
)

const N int = 1000000
const ansNum int = 11

func main() {
	ans := 0
	ansList := [ansNum]int{}
	plist := [N]bool{}
	plist[2] = true
	for i := 3; i < N; i += 2 {
		plist[i] = true
		for j := 3; j < int(math.Sqrt(float64(i)))+1; j += 2 {
			if plist[j] && i%j == 0 {
				plist[i] = false
				break
			}
		}
	}

	ansIdx := 0
	for i := 13; ansIdx < ansNum; i += 2 {
		if plist[i] {
			strprime := strconv.Itoa(i)
			strlen := len(strprime)
			for j := 1; j < strlen; j++ {
				val, _ := strconv.Atoi(strprime[:strlen-j])
				if !plist[val] {
					goto END
				}
				val, _ = strconv.Atoi(strprime[j:strlen])
				if !plist[val] {
					goto END
				}
			}
			ans += i
			ansList[ansIdx] = i
			ansIdx++
		}
	END:
	}
	fmt.Println(ans, ansList)
}
