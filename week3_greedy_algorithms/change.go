package main

import "fmt"

func getChange(m int) int {
	var i, n int
	cur := m
	sortedSizes := [3]int{10, 5, 1}
	for {
		n += cur / sortedSizes[i]
		cur = cur % sortedSizes[i]
		if cur == 0 {
			break
		}
		i += 1
	}
	return n
} 

func main() {
	var input int
	fmt.Scanln(&input)
	fmt.Println(getChange(input))
}
