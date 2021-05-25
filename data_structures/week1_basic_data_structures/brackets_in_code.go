package main

import (
	"bufio"
	"fmt"
	"os"
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

func ReadLine(reader *bufio.Reader) (string, error) {
    result := ""
    for {
        line, isPrefix, err := reader.ReadLine()
        if err != nil {
            return result, err
        }
        result += string(line)
        if !isPrefix {
            return result, nil
        }
    }
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 65536)
	text, err := ReadLine(reader)
	if err != nil {
		fmt.Println(err)
	}

	position := getUnmatchedPosition(text)
	if position == 0 {
		fmt.Println("Success")
	} else {
		fmt.Println(position)
	}
}
