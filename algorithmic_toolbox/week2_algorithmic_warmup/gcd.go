package main

import "fmt"

func gcd(s int, l int) int {
	for {
		r := l % s
		if r == 0 {
			return s
		}
		s, l = r, s
	}
}

func main() {
	var i, j int
	fmt.Scanf("%d %d", &i, &j)
	if i < j {
		fmt.Println(gcd(i, j))
	} else {
		fmt.Println(gcd(j, i))
	}
}
