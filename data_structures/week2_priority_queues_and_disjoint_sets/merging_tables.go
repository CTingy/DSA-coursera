package main

import (
	"fmt"
)

func unionByFoundIdx(parents, ranks []int, i, j int) ([]int, []int) {
	if ranks[i] > ranks[j] {
		parents[j] = i
	} else {
		parents[i] = j
		if ranks[i] == ranks[j] {
			ranks[j] += 1
		}
	}
	return parents, ranks
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
	// init parent set and find max table value
	var parents, results, ranks []int
	maxVal := 0
	for i := 0; i < tableNum; i++ {
		parents = append(parents, i)
		ranks = append(ranks, 0)
		if tables[i] > maxVal {
			maxVal = tables[i]
		}
	}

	for _, query := range queries {
		destination, source := find(parents, query[0]), find(parents, query[1])
		if destination != source {
			// make nodes on the same path have the same value
			totalValues := tables[destination] + tables[source]
			tables[source], tables[destination] = totalValues, totalValues
			parents, ranks = unionByFoundIdx(parents, ranks, destination, source)
			if totalValues > maxVal {
				maxVal = totalValues
			}
		}
		results = append(results, maxVal)
	    // fmt.Println(tables, ranks, parents)
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
