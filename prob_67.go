package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fp, err := os.Open("p067_triangle.txt")
	defer fp.Close()
	if err != nil {
		panic(err)
	}

	data := make([][]int, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		ddd := make([]int, len(tmp))
		for i := 0; i < len(tmp); i++ {
			val, _ := strconv.Atoi(tmp[i])
			ddd[i] = val
		}
		data = append(data, ddd)
	}

	last := make([]int, len(data))
	last[0] = data[0][0]
	for i := 1; i < len(data); i++ {
		bet := 0
		for j := 0; j < i; j++ {
			tmp := last[j] + data[i][j]
			if bet > tmp {
				tmp = bet
			}
			bet = last[j] + data[i][j+1]
			last[j] = tmp
		}
		last[i] = bet
	}
	ans := 0
	for i := 0; i < len(last); i++ {
		if ans < last[i] {
			ans = last[i]
		}
	}
	fmt.Println(ans, last)
}
