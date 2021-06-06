package main

import (
	"fmt"
)

func shiftDown(numbers []int, i int) {
	if len(numbers) == 1 {
		return
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
		fmt.Println(i, maxIdx)
		shiftDown(numbers, maxIdx)
	}
	return
}

func heapSort(n int, numbers []int) {
	for i := n/2 - 1; i >= 0; i-- {
		shiftDown(numbers, i)
	}
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
	heapSort(n, numbers)
}
