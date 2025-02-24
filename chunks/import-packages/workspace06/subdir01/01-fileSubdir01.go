package subdir01

import "fmt"

var unexportedName string = "Rick"

func PrintExportedHello() {
	fmt.Println("from subdir01 (subdir01 package): Hello", unexportedName)
}
