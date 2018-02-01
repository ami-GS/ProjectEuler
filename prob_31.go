package main

import "fmt"

// 1, 2, 5, 10, 20, 50, 100, 200

func calc(target int, coins []int) int {
	if target == 0 || len(coins) == 1 {
		return 1
	}
	biggest := coins[len(coins)-1]
	num := target / biggest
	ans := 0
	for i := 0; i < num+1; i++ {
		ans += calc(target-biggest*i, coins[:len(coins)-1])
	}
	return ans
}

func main() {
	fmt.Println(calc(200, []int{1, 2, 5, 10, 20, 50, 100, 200}))
}
