package main

import "fmt"

type responseBody struct {
	ColumnNames []string
}

func (rb *responseBody) fillColumnNames(fileds []string) {
	for _, cn := range fileds {
		rb.ColumnNames = append(rb.ColumnNames, cn)
		fmt.Printf("Added column name: %s\n", cn)
	}
}

func (rb *responseBody) callFillColumnNames() {
	fields := []string{"id", "name", "age", "gender", "city"}

	rb.fillColumnNames(fields)
}

func main() {
	rb := new(responseBody)
	rb.callFillColumnNames()
	fmt.Println(rb)
}
