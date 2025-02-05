package main

import (
	"fmt"
	"strings"
)

type entry struct {
	tableName string
	columns   []string
	params    []string
}

func generateSQLstring(q entry) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", q.tableName,
		strings.Join(q.columns, ", "),
		strings.Join(generatePlaceholders(len(q.params)), ", "))

	fmt.Println(query)
}

func generatePlaceholders(count int) []string {
	var placeholders []string
	for i := 1; i <= count; i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
	}
	return placeholders
}

// func generateExistsQuery() {
// 	for i, column := range
// }

func main() {
	newQuery := entry{
		tableName: "users",
		columns:   []string{"id", "name", "email", "phone"},
		params:    []string{"1", "John Doe", "johndoe@example.com", "555-32-31"},
	}

	oneMoreQuery := entry{
		tableName: "users",
		columns:   []string{"name", "email", "phone"},
		params:    []string{"Jane Doe", "janedoe@example.com", "555-32-32"},
	}
	generateSQLstring(newQuery)
	generateSQLstring(oneMoreQuery)
}
