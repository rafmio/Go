package main

import (
	"fmt"
	"imports/parser"
)

/*
structure of the project:
-08-import
	|_dirone
	| |_dironePackage.go
	|_dirtwo
	  |_dirtwoPackage.go

We should import 'dirponepackage' from the ../dirone/dironePakcage.go to the
current file (./dirtwoPackage.go)
*/

func main() {
	// Printing the data from 'dironePackage'
	parser.PrintData()

	// Printing a custom message
	fmt.Println("This is from dirtwoPackage")
}
