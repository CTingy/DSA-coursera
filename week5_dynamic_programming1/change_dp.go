package main

import (
	"fmt"
	"math"
)

// The minimum number of coins with denominations 1, 3, 4 that changes money.
func changeDp(n int) int {
	denomination := 0
    denominations := []int{1, 3, 4}

	dp := []int{0, 1, 2, 1, 1}
	for i := 5; i <= n; i++ {
		minNum := math.MaxInt32
		for j := range denominations {
			denomination = denominations[j]
			if dp[i-denomination] < minNum {
				minNum = dp[i-denomination]
			}
		}
		dp = append(dp, minNum + 1)
	}
	return dp[n]
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(changeDp(n))
}
