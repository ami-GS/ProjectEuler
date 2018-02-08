package main

import (
	"fmt"
	"math"
	"strconv"
)

func calc() {
	const N int = 10000
	pFlags := [N]bool{}
	pFlags[2] = true
	pFlags[3] = true

	for i := 3; i < N; i += 2 {
		isPrime := true
		for j := 3; j < int(math.Sqrt(float64(i)))+1; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		pFlags[i] = isPrime
	}

	for d := 1001; d < N; d += 2 {
		if !pFlags[d] {
			continue
		}
		sFlags := [10]bool{}
		for dd := d + 2; 2*dd-d < N; dd += 2 {
			if !pFlags[dd] {
				continue
			}
			if dd-d == 3330 && pFlags[2*dd-d] {
				ds := strconv.Itoa(d)
				dds := strconv.Itoa(dd)
				ddds := strconv.Itoa(2*dd - d)
				for k, _ := range ds {
					sFlags[ds[k]-'0'] = true
				}
				yes := true
				for k := 0; k < 4; k++ {
					if !sFlags[dds[k]-'0'] {
						yes = false
					}
					if !sFlags[ddds[k]-'0'] {
						yes = false
					}
				}
				if yes {
					fmt.Println(d, dd, 2*dd-d, dd-d, 2*dd-d-dd)
				}
			}
		}

	}
}

func main() {
	calc()
}
