package main

import (
	"fmt"
	// "math/rand"
	"sort"
	// "time"
)

func quickSort(eles []int) {
	lenEles := len(eles)
	if lenEles <= 1 {
		return
	}
	pivot := eles[0]
	i, j := 1, 0
	for i < lenEles {
		if eles[i] < pivot {
			eles[i], eles[j+1] = eles[j+1], eles[i]
			j++
		}
		i++
	}
	eles[0], eles[j] = eles[j], eles[0]
	quickSort(eles[0:j])  // sort smaller slice
	quickSort(eles[j+1:]) // sort larger slice
}

func majorityElement(eles []int) int {
	windowLen := len(eles)/2 + 1
	i := 0
	// quickSort(eles)
	sort.Ints(eles)
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
// 	a := []int{6, 4, 2, 3, 9, 8, 9, 4, 7, 6, 1}
// 	quickSort(a)
// 	fmt.Println(a)
// }
