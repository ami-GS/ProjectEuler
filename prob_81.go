package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const N = 80

var dp [N][N]int64

func run(mat [][]int64) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			//dp[i+1][j+1] = 1 << 21
			dp[i][j] = 0
		}
	}
	dp[0][0] = mat[0][0]
	for j := 1; j < len(mat); j++ {
		dp[0][j] = dp[0][j-1] + mat[0][j]
	}
	for i := 1; i < len(mat); i++ {
		dp[i][0] = dp[i-1][0] + mat[i][0]
	}
	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat); j++ {
			dp[i][j] = int64(math.Min(float64(dp[i][j-1]), float64(dp[i-1][j]))) + mat[i][j]
		}
	}
	//fmt.Println(dp[len(mat)-1][len(mat)-1])
	fmt.Println(dp)
}

var testVal = [][]int64{
	[]int64{131, 673, 234, 103, 18},
	[]int64{201, 96, 342, 965, 150},
	[]int64{630, 803, 746, 422, 111},
	[]int64{537, 699, 497, 121, 956},
	[]int64{805, 732, 524, 37, 331},
}

func main() {
	dat, err := ioutil.ReadFile("./p081_matrix.txt")
	if err != nil {
		panic(err)
	}
	datstring := string(dat)

	sByNL := strings.Split(strings.Trim(datstring, "\n"), "\n")
	mat := make([][]int64, len(sByNL))
	for i, line := range sByNL {
		sByComma := strings.Split(line, ",")
		mat[i] = make([]int64, len(sByComma))
		for j, valS := range sByComma {
			tmp, _ := strconv.Atoi(valS)
			mat[i][j] = int64(tmp)
		}
	}
	run(mat)
	//run(testVal)
}
