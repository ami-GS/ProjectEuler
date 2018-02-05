package main

import (
	"fmt"
	"math"
)

func done(s [10]int) bool {
	for _, d := range s {
		if d == -1 {
			return false
		}
	}
	return true
}

var ans int = 0

func buildNum(used [10]int, depth int) (prev int) {
	for i, d := range used {
		if d == depth-1 {
			prev += i * 10
		}
		if d == depth-2 {
			prev += i * 100
		}
	}
	return
}

func calc(used [10]int, depth int) {
	if done(used) {
		now, tmp := 0, 0
		for now != 10 {
			for i, d := range used {
				if d == now {
					now++
					tmp += i * int(math.Pow(10, float64(9-d)))
					fmt.Print(i)
				}
			}
		}
		fmt.Println()
		ans += tmp
		return
	}
	next := used
	switch depth {
	case 0:
		for i, d := range used {
			if i != 0 && d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	case 1:
		for i, d := range used {
			if d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	case 2:
		for i, d := range used {
			if d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	case 3:
		for i, d := range used {
			if i%2 == 0 && d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	case 4:
		prev := buildNum(used, depth)
		for i, d := range used {
			if (prev+i)%3 == 0 && d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	case 5:
		for i, d := range used {
			if i%5 == 0 && d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	case 6:
		prev := buildNum(used, depth)
		for i, d := range used {
			if (prev+i)%7 == 0 && d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	case 7:
		prev := buildNum(used, depth)
		for i, d := range used {
			if (prev+i)%11 == 0 && d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	case 8:
		prev := buildNum(used, depth)
		for i, d := range used {
			if (prev+i)%13 == 0 && d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	case 9:
		prev := buildNum(used, depth)
		for i, d := range used {
			if (prev+i)%17 == 0 && d == -1 {
				next[i] = depth
				calc(next, depth+1)
				next[i] = -1
			}
		}
	}
}

func main() {
	//fmt.Println(calc())
	flags := [10]int{}
	for i := 0; i < 10; i++ {
		flags[i] = -1
	}
	calc(flags, 0)
	fmt.Println(ans)
}
