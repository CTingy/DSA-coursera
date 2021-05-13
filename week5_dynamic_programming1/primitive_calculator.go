package main

import (
	"fmt"
	"math"
)

func calculate(n int) (int, []int) {
	if n == 0 {
		return 0, []int{}
	}

	dp := []int{0, 0}
	dpIdx := []int{0, 0}
	for i := 2; i <= n; i++ {
		minValue := math.MaxInt32
		previousIdx := 0
		if i%2 != 0 || i%3 != 0 {
			if minValue > dp[i-1]+1 {
				previousIdx = i - 1
				minValue = dp[previousIdx] + 1
			}
		}
		if i%3 == 0 {
			if minValue > dp[i/3]+1 {
				previousIdx = i / 3
				minValue = dp[previousIdx] + 1
			}
		}
		if i%2 == 0 {
			if minValue > dp[i/2]+1 {
				previousIdx = i / 2
				minValue = dp[previousIdx] + 1
			}
		}
		dp = append(dp, minValue)
		dpIdx = append(dpIdx, previousIdx)
		// fmt.Println(i, dp, dpIdx)
	}
	seqs := []int{n}
	idx := dpIdx[n]
	for idx > 1 {
		seqs = append(seqs, idx)
		idx = dpIdx[idx]
	}
	seqs = append(seqs, 1)
	return dp[n], seqs
}

func main() {
	var n int
	fmt.Scanln(&n)

	times, sequences := calculate(n)
	fmt.Println(times)
	for i := len(sequences) - 1; i >= 0; i-- {
		fmt.Printf("%d ", sequences[i])
	}
	fmt.Println()
}
