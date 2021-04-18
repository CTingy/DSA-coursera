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
    var i, j, g int
	fmt.Scanf("%d %d", &i, &j)
	if i < j {
        g = gcd(i, j)
	} else {
        g = gcd(j, i)
	}
	fmt.Println(i * (j/g))
}
