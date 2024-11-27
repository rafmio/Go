package dbops

import (
	// cmd "command-line-arguments/home/raf/Projects/Go/src/useful-scripts/random-user/cmd/randomUserResponseStruct.go"

	"database/sql"
	"fmt"
	"randomuser/cmd"

	_ "github.com/lib/pq" // $ go get .
)

func InsertEntryToDB(db *sql.DB, users cmd.RandomUserResponse) error {
	queryString := "INSERT INTO users (gender, first_name, last_name, email, phone, cell, dob, registered) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	for _, user := range users.Results {
		_, err := db.Exec(queryString,
			user.Gender,
			user.Name.First,
			user.Name.Last,
			user.Email,
			user.Phone,
			user.Cell,
			user.DOB.Date,
			user.Registered.Date,
		)
		if err != nil {
			fmt.Println("Error inserting data:", err)
		}
	}

	return nil
}
