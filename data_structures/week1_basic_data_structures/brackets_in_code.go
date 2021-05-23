package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: create an interface pop for popInt
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

func popInt(slice []int) (int, []int) {
	lastItem := 0
	popedSlice := []int{}
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
	popItem := ""
	bracketsMap["("] = ")"
	bracketsMap["["] = "]"
	bracketsMap["{"] = "}"

	backBracketMap := make(map[string]bool)
	backBracketMap[")"] = true
	backBracketMap["]"] = true
	backBracketMap["}"] = true

	positions := []int{}

	for i, s := range str {
		s := string(s) // original s is rune type
		backBracket, exists := bracketsMap[s]
		if exists {
			stack = append(stack, backBracket)
			positions = append(positions, i+1)
			continue
		}
		if backBracketMap[s] {
			popItem, stack = pop(stack)
			if popItem != s {
				return i + 1
			}
			_, positions = popInt(positions)
		}
	}
	if len(stack) == 0 {
		return 0
	}
	return positions[len(positions)-1]
}

func main() {
	// var input string
	// fmt.Scanf("%s", &input)

	reader := bufio.NewReaderSize(os.Stdin, 65536)
	text, _ := reader.ReadString('\n')

	position := getUnmatchedPosition(text)
	if position == 0 {
		fmt.Println("Success")
	} else {
		fmt.Println(position)
	}
}
