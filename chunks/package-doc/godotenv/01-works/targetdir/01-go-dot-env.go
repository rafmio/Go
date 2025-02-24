package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from dbconf.env file
	err := godotenv.Load("../config/dbconf.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Assign environment variables to Go variables
	musicInfoHost := os.Getenv("MUSIC_INFO_HOST")
	musicInfoPort := os.Getenv("MUSIC_INFO_PORT")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresTableName := os.Getenv("POSTGRES_TABLE_NAME")
	postgresSslMode := os.Getenv("POSTGRES_SSL_MODE")
	postgresDriverName := os.Getenv("POSTGRES_DRIVER_NAME")

	// Print the values of the variables
	fmt.Println("Music Info Host:", musicInfoHost)
	fmt.Println("Music Info Port:", musicInfoPort)
	fmt.Println("Postgres Host:", postgresHost)
	fmt.Println("Postgres Port:", postgresPort)
	fmt.Println("Postgres User:", postgresUser)
	fmt.Println("Postgres Password:", postgresPassword)
	fmt.Println("Postgres DB:", postgresDB)
	fmt.Println("Postgres Table Name:", postgresTableName)
	fmt.Println("Postgres SSL Mode:", postgresSslMode)
	fmt.Println("Postgres Driver Name:", postgresDriverName)
}
