package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxPairwiseProduct(numbers []int) int {
	maxIdx := 0
	max := numbers[maxIdx]
	for i := 1; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
		if numbers[i] > max {
			maxIdx = i
			max = numbers[i]
		}
	}
	fmt.Println(maxIdx, max)
	secMax := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > secMax && i != maxIdx {
			secMax = numbers[i]
		}
	}
	fmt.Println(secMax)
	return max * secMax
}

func main() {
	var numLen int
	fmt.Scanf("%d", &numLen)

	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	texts := strings.Split(input, " ")
	for _, text := range texts {
		number, _ := strconv.Atoi(text)
		numbers = append(numbers, number)
	}
	fmt.Println(maxPairwiseProduct(numbers))
}
