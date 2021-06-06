package main

import (
	"fmt"
)

func shiftDown(numbers []int, i int, swaps [][]int) [][]int {
	if len(numbers) == 1 {
		return swaps
	}
	maxIdx := i
	rightIdx := 2*i + 1
	leftIdx := 2*i + 2
	if leftIdx < len(numbers) && numbers[leftIdx] < numbers[maxIdx] {
		maxIdx = leftIdx
	}
	if rightIdx < len(numbers) && numbers[rightIdx] < numbers[maxIdx] {
		maxIdx = rightIdx
	}
	if maxIdx != i {
		numbers[i], numbers[maxIdx] = numbers[maxIdx], numbers[i]
		// fmt.Println(i, maxIdx)
		swaps = shiftDown(numbers, maxIdx, append(swaps, []int{i, maxIdx}))
	}
	return swaps
}

func heapSort(n int, numbers []int) [][]int {
	swaps := [][]int{}
	for i := n/2 - 1; i >= 0; i-- {
		swaps = shiftDown(numbers, i, swaps)
	}
	return swaps
}

func main() {
	var n int
	var numbers []int

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		number := 0
		fmt.Scanf("%d", &number)
		numbers = append(numbers, number)
	}
	swaps := heapSort(n, numbers)
	fmt.Println(len(swaps))
	for _, swap := range swaps {
		fmt.Println(swap[0], swap[1])
	}
}
