package main

import (
	"fmt"
)

type query struct {
	action string
	number string
	name   string
}

func execQuery(q query, phoneBook map[string]string) {
	if q.action == "add" {
		phoneBook[q.number] = q.name
	} else if q.action == "find" {
		name, exists := phoneBook[q.number]
		if exists {
			fmt.Println(name)
		} else {
			fmt.Println("not found")
		}
	} else if _, exists := phoneBook[q.number]; q.action == "del" && exists {
		delete(phoneBook, q.number)
	}
}

func main() {
	var number int
	fmt.Scanln(&number)

	queries := []query{}
	for i := 0; i < number; i++ {
		var action, number, name string
		fmt.Scanf("%s %s %s", &action, &number, &name)
		queries = append(queries, query{action, number, name})
	}
	phoneBook := make(map[string]string)
	for _, query := range queries {
		execQuery(query, phoneBook)
	}
}
