package main

import (
	"fmt"
	"strings"
)

func main() {
	sls := []string{"cute_ganymede", "black_oxygenium", "test_server"}

	for _, srvName := range sls {
		if strings.Contains(srvName, "test") {
			continue
		} else {
			fmt.Println(srvName)
		}
	}
}
