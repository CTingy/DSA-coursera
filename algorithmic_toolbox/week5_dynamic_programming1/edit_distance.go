package main

import (
	"fmt"
)

func min(nums [3]int) int {
	minValue := nums[0]
	for i := 1; i < 3; i++ {
		if nums[i] < minValue {
			minValue = nums[i]
		}
	}
	return minValue
}

func editDistance(str1, str2 string) int {
	// init
	matrix := make([][]int, len(str2)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(str1)+1)
	}

	var match, deletion, insertion, mismatch int
	for i := 0; i <= len(str2); i++ {
		for j := 0; j <= len(str1); j++ {
			if i == 0 {
				matrix[0][j] = j
				continue
			} else if j == 0 {
				matrix[i][0] = i
				continue
			}
			// start comparing
			ch1, ch2 := str1[j-1], str2[i-1]
			match = matrix[i-1][j-1]
			mismatch = matrix[i-1][j-1] + 1
			deletion = matrix[i][j-1] + 1
			insertion = matrix[i-1][j] + 1

			if ch1 == ch2 {
				matrix[i][j] = min([3]int{insertion, deletion, match})
			} else {
				matrix[i][j] = min([3]int{insertion, deletion, mismatch})
			}
		}
	}
	//fmt.Println(matrix)
	return matrix[len(str2)][len(str1)]
}

func main() {
	var str1, str2 string
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)
	fmt.Println(editDistance(str1, str2))
}
