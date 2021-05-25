package main

import (
	"fmt"
)

func getTreeHeight(nodes []int) int {
	var node, height, count int

	for i := 0; i < len(nodes); i++ {
		node = nodes[i]
		count = 0
		for node > -1 {
			node = nodes[node]
			count++
		}
		if count > height {
			height = count
		}
	}
	return 1 + height
}

func main() {
	var n int
	var nodes []int

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		node := 0
		fmt.Scanf("%d", &node)
		nodes = append(nodes, node)
	}
	fmt.Println(getTreeHeight(nodes))
}
