package main

import (
	"fmt"
	"sort"
)

func fractionalKnapsack(n int, totalW float64, valuesAndWeights [][]float64) float64 {
	var fraction float64
    var sortedFractions []float64
	var res float64
    
	// get fractions and sort
	items := make(map[float64][]float64)
	for i := 0; i < n; i++ {
		fraction = valuesAndWeights[i][0] / valuesAndWeights[i][1]
		items[fraction] = valuesAndWeights[i]
		sortedFractions = append(sortedFractions, fraction)
	}
	// sortedFractions = sort.Float64Slice(sortedFractions)
	sort.Sort(sort.Reverse(sort.Float64Slice(sortedFractions)))

	// calculate
	for i := 0; i < n; i++ {
		fraction = sortedFractions[i]
		valuesAndWeight := items[fraction]
        if totalW < valuesAndWeight[1] {
			return res + fraction * totalW
		}
		totalW -= valuesAndWeight[1]
		res += valuesAndWeight[0]
	}
	return -1.0
}

func main() {
	var n int
	var totalW float64
	fmt.Scanln(&n, &totalW)
	valuesAndWeights := make([][]float64, n)

	for i := 0; i < n; i++ {
		var v, w int
		fmt.Scanln(&v, &w)
		valuesAndWeights[i] = []float64{float64(v), float64(w)}
	}
	fmt.Printf("%.4f\n", fractionalKnapsack(n, totalW, valuesAndWeights))
}
