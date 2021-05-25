package main

import (
	"fmt"
)

func getTreeHeight(nodes []int, nodeMap map[int][]int) int {
	height := 1
	for len(nodes) > 0 {
		subNodes := []int{}
		for _, node := range nodes {
			subNodes = append(subNodes, nodeMap[node]...)
		}
		nodes = subNodes
		height++
	}
	return height
}

func main() {
	n, root := 0, 0
	nodeMap := make(map[int][]int)

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		node := 0
		fmt.Scanf("%d", &node)
		if node == -1 {
			root = i
		} else {
			nodeMap[node] = append(nodeMap[node], i)
		}
	}
	fmt.Println(getTreeHeight(nodeMap[root], nodeMap))
}
