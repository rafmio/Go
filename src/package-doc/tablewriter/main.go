// https://pkg.go.dev/github.com/xlab/tablewriter#section-readme
package main

import (
	"fmt"

	"github.com/xlab/tablewriter"
)

func main() {
	table := tablewriter.CreateTable()

	table.AddHeaders("Name", "Age")
	table.AddRow("John", "30")
	table.AddRow("Sam", 18)
	table.AddRow("Jane", "25")
	table.AddRow("Julie", 20.14)

	fmt.Println(table.Render())
}
