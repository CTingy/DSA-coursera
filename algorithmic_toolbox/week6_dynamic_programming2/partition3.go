package main

import (
	"fmt"
	"sort"
)

func nPartition(n int, total int, eles []int) bool {
	sort.Slice(eles, func(i, j int) bool {
		return eles[i] > eles[j]
	})
	canPartition := false
	partitionCount := n
	if total%n != 0 {
		return false
	}
	for partitionCount > 0 {
		canPartition, eles = partition(total / n, eles)
		// fmt.Println(canPartition, eles)
		if !canPartition {
			return false
		}
		partitionCount--
	}
	return true
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
		for tmpVal >= 0 && j >= 1 && dp[tmpVal][j] {
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
    res := nPartition(3, total, eles)
	// partition(total/3, eles)
	if res {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

// func main() {
//     s := []int{8, 2, 6, 3, 1, 4} 
//     sort.Slice(s, func(i, j int) bool {
//         return s[i] > s[j]
//     })
//     fmt.Println(s)  // [8 6 4 3 2 1]
// }
