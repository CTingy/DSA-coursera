package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	x = 263
	p = 1000000007
)

type query struct {
	action  string
	content string
}

type hashChain struct {
	size   int
	chains [][]string
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func H(s string, m int) int {
	res, power := 0, 0
	for idx, char := range s {
		power = int(math.Pow(x, float64(idx)))
		res = (res + int(char)*power) % p
	}
	return res % m
}

func (hc hashChain) add(key int, value string) {
	if hc.find(value) {
		return
	}
	hc.chains[key] = append([]string{value}, hc.chains[key]...)
}

func (hc hashChain) check(key int) {
	values := hc.chains[key]
	if len(values) > 0 {
		var s string
		for _, value := range values {
			s += fmt.Sprintf("%s ", value)
		}
		fmt.Println(s)
	} else {
		fmt.Println("")
	}
}

func (hc hashChain) find(content string) bool {
	for _, chain := range hc.chains {
		for _, value := range chain {
			if value == content {
				return true
			}
		}
	}
	return false
}

func (hc hashChain) del(content string) {
	for key, chain := range hc.chains {
		for idx, value := range chain {
			if value == content {
				hc.chains[key] = RemoveIndex(chain, idx)
			}
		}
	}
}

func execQuery(hc hashChain, q query) {
	if q.action == "add" {
		hc.add(H(q.content, hc.size), q.content)
	} else if q.action == "check" {
		key, _ := strconv.Atoi(q.content)
		hc.check(key)
	} else if q.action == "find" {
		if hc.find(q.content) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	} else if q.action == "del" {
		hc.del(q.content)
	}
}

func main() {
	var m, num int
	fmt.Scanln(&m)
	fmt.Scanln(&num)

	queries := make([]query, num)
	for i := 0; i < num; i++ {
		var s1, s2 string
		fmt.Scanf("%s %s", &s1, &s2)
		queries[i] = query{action: s1, content: s2}
	}

	hc := hashChain{size: m, chains: make([][]string, m)}
	for _, q := range queries {
		execQuery(hc, q)
	}
	// fmt.Println(hc.chains)
}
