package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	anssum := 2
	ans := 1
	N := 1000000
	plist := [1000000]bool{}
	plist[2] = true
	for i := 3; i < N; i += 2 {
		f := true
		for j := 3; j < int(math.Sqrt(float64(i)))+1; j += 2 {
			if plist[j] && i%j == 0 {
				f = false
				break
			}
		}
		plist[i] = f
	}

	for i := 3; i < N; i += 2 {
		if plist[i] {
			str := strconv.Itoa(i)
			lll := len(str)
			// not good here
			same := true
			if lll == 1 {
				same = false
			} else {
				for j := 0; j < lll-1; j++ {
					if str[j] != str[j+1] {
						same = false
						break
					}
				}
			}

			circle_prime := true
			ansList := make([]int, lll)
			ansList[0] = i
			strbuff := make([]byte, lll)
			for j := 1; j < lll; j++ {
				tmp := str[lll-1]
				for k := 0; k < lll-1; k++ {
					if (str[k]-'0')%2 == 0 {
						circle_prime = false
						break
					}
					strbuff[k+1] = str[k]
				}
				if !circle_prime {
					break
				}

				strbuff[0] = tmp
				val, _ := strconv.Atoi(string(strbuff))
				if !plist[val] {
					circle_prime = false
					break
				}
				ansList[j] = val
				str = string(strbuff)
			}
			if circle_prime {
				for j := 0; j < lll; j++ {
					anssum += ansList[j]
					ans++
					plist[ansList[j]] = false
				}
				if same {
					anssum -= ansList[0] * (lll - 1)
					ans--
				}
				fmt.Println(str, ansList)
			}
		}
	}

	fmt.Println(anssum, ans)

}
