package main

import (
	"fmt"
)

func threePartition(total int, eles []int) int {
	canPartition := false
	partitionCount := 3
	if total%partitionCount != 0 {
		return 1
	}
	value := total / partitionCount
	for partitionCount > 0 {
		canPartition, eles = partition(value, eles)
		if !canPartition {
			return 1
		}
		partitionCount--
	}
	if len(eles) == 0 {
		return 0
	}
	return 1
}

// case example: 7 2 2 2 2 2 2 2 3, res: wrong
// case example: 2 1 4 4 4 6, res: wrong
// 2 partition ref: http://thisthread.blogspot.com/2018/02/2-partition-problem.html
func partition(value int, eles []int) (bool, []int) {
	// init DP 2D array
	dp := make([][]bool, value+1)
	for i := 0; i <= value; i++ {
		dp[i] = make([]bool, len(eles)+1)
	}
	for i := 1; i <= value; i++ {
		for j := 1; j <= len(eles); j++ {
			if dp[i][j-1] || i == eles[j-1] {
				dp[i][j] = true
			} else {
				valueDiff := i - eles[j-1]
				if valueDiff > 0 && dp[valueDiff][j-1] {
					dp[i][j] = true
				}
			}
		}
		fmt.Println(dp[i])
	}
	return dp[value+1][len(eles)+1], eles
}

func main() {
	var n, total int
	fmt.Scanln(&n)
	eles := []int{}

	for i := 0; i < n; i++ {
		var ele int
		fmt.Scanf("%d", &ele)
		total += ele
		eles = append(eles, ele)
	}
	//fmt.Println(threePartition(total, eles))
	partition(total/3, eles)
}
