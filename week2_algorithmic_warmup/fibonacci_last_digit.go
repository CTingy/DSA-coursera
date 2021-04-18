package main

import "fmt"

func Fibonacci(n int) int {
    var s []int
	s = append(s, 0)
	s = append(s, 1)
	for i := 2; i <= n; i++ {
		s = append(s, (s[i-1]+s[i-2])%10)
	}
	return s[len(s)-1]
}

func main() {
	var i int
	fmt.Scanf("%d", &i)
    fmt.Println(Fibonacci(i))
}
