package main

import (
	"strings"
	"fmt"
)

func main() {
	var inputString string
	fmt.Scanln(&inputString)
	s := strings.Split(inputString, " ")
	fmt.Println(s)
}
