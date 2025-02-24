package main

import (
	"database/sql"
	"fmt"
)

func main() {
	drivers := sql.Drivers()
	for i, val := range drivers {
		fmt.Println(i, val)
	}
	fmt.Println(drivers)
}
