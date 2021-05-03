package main

import "fmt"

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

func main() {
	a := []int{6, 4, 2, 3, 9, 8, 9, 4, 7, 6, 1}
	quickSort(a)
	fmt.Println(a)
}
