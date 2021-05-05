package main

import (
	"fmt"
	// "sort"
	// "time"
)

// func quickSort(eles []int) {
// 	lenEles := len(eles)
// 	if lenEles <= 1 {
// 		return
// 	}
// 	pivot := eles[lenEles-1]
// 	i, j := 0, 0
// 	for i < lenEles-1 {
// 		if eles[i] < pivot {
// 			eles[i], eles[j] = eles[j], eles[i]
// 			j++
// 		}
// 		i++
// 	}
// 	eles[lenEles-1], eles[j] = eles[j], eles[lenEles-1]
// 	quickSort(eles[0:j])  // sort smaller slice
// 	quickSort(eles[j+1:]) // sort larger slice
// }

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

func majorityElement(eles []int) int {
	windowLen := len(eles)/2 + 1
	i := 0
	// sort.Ints(eles)
	// quickSort(eles)
	threePartitionQuickSort(eles)
	for i+windowLen-1 < len(eles) {
		if eles[i] == eles[i+windowLen-1] {
			return 1
		}
		i++
	}
	return 0
}

func main() {
	var n, ele int
	var eles []int

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ele)
		eles = append(eles, ele)
	}
	fmt.Println(majorityElement(eles))
}

// stress test

// func main() {
// 	for {
// 		var a, b []int

// 		n := rand.Intn(100000)
// 		for i := 0; i < n; i++ {
// 			x := rand.Intn(1000000000)
// 			a = append(a, x)
// 			b = append(b, x)
// 		}
// 		quickSort(a)
// 		sort.Ints(b)

// 		for i := 0; i < n; i++ {
// 			if a[i] != b[i] {
// 				fmt.Println(a, b)
// 				panic("")
// 			}
// 		}
// 		time.Sleep(1 * time.Millisecond)
// 	}
// }

// check sort
// func main() {
// 	// a := []int{6, 4, 2, 3, 9, 8, 9, 4, 7, 6, 1}
// 	a := []int{5, 3, 2, 4, 8, 4}
// 	quickSort(a)
// 	fmt.Println(a)
// }
