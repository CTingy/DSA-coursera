package main

import (
	"fmt"
)

func knapsack(totalW, n int, items map[int]int) int {
	val := 0
	dp := make([][]int, n+1)

	// init row 0
	dp[0] = make([]int, totalW+1)

	itemCount := 1
	for item := range items {
		for j := 0; j <= totalW; j++ {
			if j == 0 {
				val = 0
			} else if j == item {
				// take the item directly
				val = j
			} else if j < item {
				val = dp[itemCount-1][j]
			} else {
				// pick out the max one
				if dp[itemCount-1][j] >= dp[itemCount-1][j-item]+item {
					val = dp[itemCount-1][j]
				} else {
					val = dp[itemCount-1][j-item] + item
				}
			}
			dp[itemCount] = append(dp[itemCount], val)
		}
		itemCount++
	}
	// fmt.Println(dp)
	return dp[n][totalW]
}

func main() {
	var totalW, n int
	fmt.Scanln(&totalW, &n)

	items := make(map[int]int)
	for i := 0; i < n; i++ {
		item := 0
		fmt.Scanf("%d", &item)
		items[item] = 1 // mark this item exists
	}
	fmt.Println(knapsack(totalW, n, items))
}
