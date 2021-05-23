package main

import (
	"fmt"
	"strings"
	"strconv"
)

func binarySearch(start int, end int, target int, eles []int) int {
	if start == end {
		return -1
	}
	midInx := (start + end) / 2
	mid := eles[midInx]
	if mid == target {
		return midInx
	}
	if mid > target {
		end = midInx
	} else {
		start = midInx + 1
	}
	return binarySearch(start, end, target, eles)
}

func main() {
	var m, n, idx int
	var eles []int
	var res []string

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var ele int
		fmt.Scanf("%d", &ele)
		eles = append(eles, ele)
	}

	fmt.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		var target int
		fmt.Scanf("%d", &target)
		idx = binarySearch(0, len(eles), target, eles)
		res = append(res, strconv.Itoa(idx))
	}
	fmt.Println(strings.Join(res, " "))
}
