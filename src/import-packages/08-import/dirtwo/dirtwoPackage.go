package main

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

import (
	"fmt"

	"../dironepackage"
)

func main() {
	// Printing the data from 'dironePackage'
	dironepackage.PrintData()

	// Printing a custom message
	fmt.Println("This is from dirtwoPackage")
}
