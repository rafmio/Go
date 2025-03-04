package main

import (
	"fmt"
	"os"
)

func main() {
	envVarName := "HOME"

	value, exists := os.LookupEnv(envVarName)

	if exists {
		fmt.Printf("Environment variable %s is set to: %s\n", envVarName, value)
	} else {
		fmt.Printf("Environment variable %s is not set.\n", envVarName)
	}
}
