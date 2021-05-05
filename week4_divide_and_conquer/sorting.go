package main

import "fmt"

func threePartitionQuickSort(eles []int) {
	lenEles := len(eles)
	if lenEles <= 1 {
		return
	}
	pivot := eles[0]
	i, j, k := 1, 0, 0
	for i < lenEles {
		if eles[i] < pivot {
			eles[i], eles[j+1] = eles[j+1], eles[i]
			j++
		} else if eles[i] == pivot {
			eles[i], eles[k+1] = eles[k+1], eles[i]
		}
		i++
	}
	for {
		eles[k], eles[j] = eles[j], eles[k]
		if k == 0 {
			break
		}
		k--
		j--
	}
	threePartitionQuickSort(eles[0:j])    // sort smaller slice
	threePartitionQuickSort(eles[j+k+1:]) // sort larger slice
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
