package main

import (
	"fmt"
)

func nPartition(n int, total int, eles []int) int {
	canPartition := false
	partitionCount := n
	if total%n != 0 {
		return 0
	}
	for partitionCount > 0 {
		canPartition, eles = partition(total / n, eles)
		fmt.Println(canPartition, eles)
		if !canPartition {
			return 0
		}
		partitionCount--
	}
	return 1
}

// case example: 9 & 7 2 2 2 2 2 2 2 3, res: wrong
// case example: 6 & 2 1 4 4 4 6, res: wrong
// 4 [4,4,6,2,3,8,10,2,10,7], res: correct
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
		// fmt.Println(dp[i])
	}

	// rule out selected values
	selectedIdx := make(map[int]bool)
	tmpVal := value
	for j := 1; j <= len(eles); j++ {
		for tmpVal > 0 && j > 0 && dp[tmpVal][j] {
			if !dp[tmpVal][j-1] {
				selectedIdx[j-1] = true
				tmpVal -= eles[j-1]
			}
			j--
		}
	}
	// fmt.Println(selectedIdx)
	newEles := []int{}
	for i, ele := range eles {
		if !selectedIdx[i] {
			newEles = append(newEles, ele)
		}
	}
	return dp[value][len(eles)], newEles
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
	fmt.Println(nPartition(3, total, eles))
	// partition(total/3, eles)
}
