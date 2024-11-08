package main

import (
	"fmt"
	"readpath/cmd"
)

func main() {
	err := cmd.Runner()
	if err != nil {
			fmt.Println("ERROR:", err)
	}

}
