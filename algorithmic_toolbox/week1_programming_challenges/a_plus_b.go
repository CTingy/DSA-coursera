package main

import "fmt"

func APlusB(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	var firstDigit, secondDigit int
	fmt.Scanf("%d %d", &firstDigit, &secondDigit)
	APlusB(firstDigit, secondDigit)
}
