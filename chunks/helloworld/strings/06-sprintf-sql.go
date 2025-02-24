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
func buildExistsQuery(req entry) (string, error) {
	if len(req.params) != len(req.columns) {
		return "", fmt.Errorf("params and columns must have the same length")
	}

	var conditions []string

	// conditions for WHERE
	for i, column := range req.columns {
		conditions = append(conditions, fmt.Sprintf("%s = $%d", column, i+1))
	}

	query := fmt.Sprintf(
		"SELECT EXISTS (SELECT 1 FROM %s WHERE %s)",
		req.tableName,
		strings.Join(conditions, " AND "),
	)

	return query, nil
}

func main() {
	newQuery := entry{
		tableName: "users",
		columns:   []string{"id", "name", "email", "phone"},
		params:    []string{"1", "Ivan Pomidorov", "ivanterrorist@example.com", "555-32-31"},
	}

	oneMoreQuery := entry{
		tableName: "users",
		columns:   []string{"name", "email", "phone"},
		params:    []string{"Jane Doe", "janedoe@example.com", "555-32-32"},
	}
	generateSQLstring(newQuery)
	generateSQLstring(oneMoreQuery)

	existQueryOne, err := buildExistsQuery(newQuery)
	if err != nil {
		fmt.Println("Error building exists query:", err)
		return
	} else {
		fmt.Println(existQueryOne)
	}

	existsQuery, err := buildExistsQuery(oneMoreQuery)
	if err != nil {
		fmt.Println("Error building exists query:", err)
		return
	} else {
		fmt.Println(existsQuery)
	}
}
