package cmd

import (
	"readpath/handlers"
	"fmt"
)

func Runner() {
	result, err := handlers.OpenAndReadFile("../config/conf.txt")
	if err != nil {
		return err
	}

	fmt.Println(result)
}
