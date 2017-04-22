package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ans := 0
	ansList := []string{}

	fp, err := os.Open("p042_words.txt")
	defer fp.Close()
	if err != nil {
		panic(err)
	}

	var data []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		data = strings.Split(scanner.Text(), ",")
	}

	maxLen := 0
	for i := 0; i < len(data); i++ {
		if maxLen < len(data[i]) {
			maxLen = len(data[i])
		}
	}
	N := maxLen * 26
	isTri := make([]bool, N)
	for i := 1; i*(i+1)/2 < N; i++ {
		isTri[i*(i+1)/2] = true
	}
	for i := 0; i < len(data); i++ {
		stmp := data[i][1 : len(data[i])-1]
		sum := 0
		for j := 0; j < len(stmp); j++ {
			sum += int(stmp[j] - 'A' + 1)
		}
		if isTri[sum] {
			ans++
			ansList = append(ansList, stmp)
		}

	}

	fmt.Println(ans, ansList)
}
