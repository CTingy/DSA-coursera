package main

import (
	"fmt"
)

func knapsack(totalW, n int, items []int) int {
	val := 0
	dp := make([][]int, n+1)

	// init row 0
	dp[0] = make([]int, totalW+1)

	for i, item := range items {
		for j := 0; j <= totalW; j++ {
			if j == 0 {
				val = 0
			} else if j == item {
				// take the item directly
				val = j
			} else if j < item {
				val = dp[i][j]
			} else {
				// pick out the max one
				if dp[i][j] >= dp[i][j-item]+item {
					val = dp[i][j]
				} else {
					val = dp[i][j-item] + item
				}
			}
			dp[i+1] = append(dp[i+1], val)
		}
	}
	// fmt.Println(dp)
	return dp[n][totalW]
}

func main() {
	var totalW, n int
	fmt.Scanln(&totalW, &n)

	items := []int{}
	for i := 0; i < n; i++ {
		item := 0
		fmt.Scanf("%d", &item)
		items = append(items, item)
	}
	fmt.Println(knapsack(totalW, n, items))
}
