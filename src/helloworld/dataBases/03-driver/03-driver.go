package main

import (
	"database/sql"
	"fmt"

	// import the PostgreSQL driver for datebase/sql
	_ "github.com/lib/pq" // $ go get .
)

func main() {
	drivers := sql.Drivers()

	fmt.Println(drivers)

}
