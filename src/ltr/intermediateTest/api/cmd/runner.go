package cmd

import (
	"fmt"
	"readpath/handlers"

	pho "readpath/pkg/printhelloone"
	pht "readpath/pkg/printhellotwo"
)

func Runner() error {
	// result, err := handlers.OpenAndReadFile("../config/conf.txt") // doesn't work
	result, err := handlers.OpenAndReadFile("./config/conf.txt")
	if err != nil {
		return err
	}

	for _, val := range result {
		fmt.Println(val)
	}

	pho.PrintHelloOne()
	pht.PrintHelloTwo()

	setEnv()
	getEnv()

	return nil
}
