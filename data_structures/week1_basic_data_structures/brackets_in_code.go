package main

import (
	"fmt"
)

func pop(slice []string) (string, []string) {
	lastItem := ""
	popedSlice := []string{}
	if len(slice) > 0 {
		lastItem = slice[len(slice)-1]
		if len(slice) >= 2 {
			popedSlice = append(popedSlice, slice[:len(slice)-1]...)
		}
	}
	return lastItem, popedSlice
}

func getUnmatchedPosition(str string) int {
	bracketsMap := make(map[string]string)
	stack := []string{}
	bracketsMap["("] = ")"
	bracketsMap["["] = "]"
	bracketsMap["{"] = "}"

	backBracketMap := make(map[string]bool)
	backBracketMap[")"] = true
	backBracketMap["]"] = true
	backBracketMap["}"] = true

	unmatchedPosition := 0

	for i, s := range str {
		s := string(s) // original s is rune type
		backBracket, exists := bracketsMap[s]
		if exists {
			stack = append(stack, backBracket)
			continue
		}
		if backBracketMap[s] {
			popItem, popedSlice := pop(stack)
			if popItem != s {
				unmatchedPosition = i + 1
				break
			}
			stack = popedSlice
		}
	}
	return unmatchedPosition
}

func main() {
	var input string
	fmt.Scanln(&input)

	position := getUnmatchedPosition(input)
	if position == 0 {
		fmt.Println("Success")
	} else {
		fmt.Println(position)
	}
}
