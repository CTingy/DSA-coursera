package main

import (
	"fmt"
)

func max(values []int) int {
	maxVal := 0
	for _, i := range values {
		if i > maxVal {
			maxVal = i
		}
	}
	return maxVal
}

func find(parents []int, i int) int {
	for i != parents[i] {
		// i = parents[i]
		// need to use path compression, otherwise: time limit exceeded
		i = find(parents, parents[i])
	}
	return i
}

func execQuery(tableNum int, tables []int, queries [][]int) []int {
	// init parent set
	var parents, results []int
	for i := 0; i < tableNum; i++ {
		parents = append(parents, i)
	}
	for _, query := range queries {
		destination, source := find(parents, query[0]), find(parents, query[1])
		if destination != source {
			tables[destination] += tables[source]
			tables[source] = 0
			parents[source] = destination
		}
		results = append(results, max(tables))
		// fmt.Println(tables)
	}
	return results
}

func main() {
	var tableNum, queryNum int
	var tables []int
	var queries [][]int

	fmt.Scanln(&tableNum, &queryNum)
	for i := 0; i < tableNum; i++ {
		table := 0
		fmt.Scanf("%d", &table)
		tables = append(tables, table)
	}
	for i := 0; i < queryNum; i++ {
		destination, source := 0, 0
		fmt.Scanf("%d %d\n", &destination, &source)
		queries = append(queries, []int{destination - 1, source - 1})
	}

	// execute
	results := execQuery(tableNum, tables, queries)
	for _, result := range results {
		fmt.Println(result)
	}
}
