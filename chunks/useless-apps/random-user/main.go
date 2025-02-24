// generated using Claude 3.5
package main

import (
	"fmt"
	"randomuser/cmd"
	"randomuser/dbops"
)

func main() {
	randomUserResponse, err := cmd.GetResponse()
	if err != nil {
		fmt.Println("Error fetching random user data:", err)
		return
	} else {
		fmt.Println("random user parsed successfully")
	}

	// DEBUG print
	// fmt.Println(randomUserResponse)

	db, err := dbops.EstablishDBConnection()
	if err != nil {
		fmt.Println("Error establishing database connection:", err)
		return
	}

	err = dbops.InsertEntryToDB(db, randomUserResponse)
	if err != nil {
		fmt.Println("Error inserting data into the database:", err)
		return
	}

}
