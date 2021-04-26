package main

import "fmt"

func fractionalKnapsack(d int, m int, stops []int) int {
	count := 0
	lastRefillStop := 0
	stops = append(stops, d)
	for i := 0; i < len(stops); i++ {
		if stops[i]-lastRefillStop > m {
			return -1
		}
		// no need to check this when the last stop before the end
		if i == len(stops)-1 {
			break
		}
		if stops[i+1]-lastRefillStop > m {
			count += 1
			lastRefillStop = stops[i]
		}
	}
	return count
}

func main() {
	var d, m, n, stop int
	var stops []int

	fmt.Scanln(&d)
	fmt.Scanln(&m)
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &stop)
		stops = append(stops, stop)
	}
	fmt.Println(fractionalKnapsack(d, m, stops))
}
