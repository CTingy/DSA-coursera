package main

import "fmt"

func threePartitionQuickSort(eles []int) {
	lenEles := len(eles)
	if lenEles <= 1 {
		return
	}
	pivot := eles[lenEles-1]
	i, j, k := 0, 0, lenEles-1
	for i < k {
		if eles[i] < pivot {
			eles[i], eles[j] = eles[j], eles[i]
			j++
			i++
		} else if eles[i] == pivot {
			eles[i], eles[k-1] = eles[k-1], eles[i]
			k--
		} else {
			i++
		}
	}

	lenPivot := lenEles - k
	for k < lenEles {
		eles[j], eles[k] = eles[k], eles[j]
		j++
		k++
	}
	threePartitionQuickSort(eles[:j-lenPivot]) // sort smaller slice
	threePartitionQuickSort(eles[j:])          // sort larger slice
}

func main() {
	var n int
	var eles []int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scanf("%d", &x)

		eles = append(eles, x)
	}
	threePartitionQuickSort(eles)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", eles[i])
	}
	fmt.Println()
}
