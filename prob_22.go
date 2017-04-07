package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fp, err := os.Open("p022_names.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fp)

	var names []string
	for scanner.Scan() {
		txt := scanner.Text()
		names = strings.Split(txt, ",")
	}

	sort.Slice(names, func(i, j int) bool {
		II := len(names[i])
		JJ := len(names[j])
		min := II
		if JJ < II {
			min = JJ
		}
		for idx := 1; idx < min-1; idx++ {
			if byte(names[i][idx]) == byte(names[j][idx]) {
				continue
			}
			return byte(names[i][idx]) < byte(names[j][idx])
		}
		if II < JJ {
			return true
		}
		return false
	})
	fmt.Println(names)

	ans := 0
	for i := 0; i < len(names); i++ {
		tmp := 0
		for j := 1; j < len(names[i])-1; j++ {
			tmp += int(byte(names[i][j]) - 'A' + 1)
		}
		ans += tmp * (i + 1)
	}
	fmt.Println(ans)
}
