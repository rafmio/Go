package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DBConfig struct {
	Name        string `json:"Name"`        // server's name for internal using in code, mapping etc ('cute_ganymede')
	DisplayName string `json:"DisplayName"` // the same name as 'Name', only for display ('Cute Ganymede')
	DriverName  string `json:"DriverName"`  // e.g. "postgres"
	Host        string `json:"Host"`        // "194.58.102.129", "localhost", etc
	Port        string `json:"Port"`        // port number, e.g. "5432", "8543", etc
	DBName      string `json:"DBName"`      // name of DB inside of 'PostgreSQL'
	User        string `json:"User"`        // username "raf", "postgres", etc
	Password    string `json:"Password"`    // password
	SslMode     string `json:"SslMode"`     // SSL mode, etc "disable", "require", "verify-full", etc"
	Dsn         string // data source name
	DB          *sql.DB
	Err         error
}

func NewDBConfig(dbConfigFilePath, srcName string) (*DBConfig, error) {

	// check if dbConfigFilePath is empty
	if dbConfigFilePath == "" {
		log.Println("Database config file path is empty")
		return nil, fmt.Errorf("Database config file path is empty")
	}

	// reading file with configuration for DB connection
	file, err := os.ReadFile(dbConfigFilePath)
	if err != nil {
		log.Println("Opening config file:", err)
		return nil, err
	}

	// unmarshalling JSON data to struct
	dbConfigs := make(map[string]DBConfig) // variable for storing unmarshalled data
	err = json.Unmarshal(file, &dbConfigs)
	if err != nil {
		log.Println("Unmarshalling JSON:", err)
		return nil, err
	}

	dbCfg := new(DBConfig) // new DBConfig instance
	// dbCfg = &dbConfigs[srcName] // assigning data from map to struct for given source name
	tmpCfg := dbConfigs[srcName] // assigning data from map to struct for given source name
	*dbCfg = tmpCfg              // assigning data from map to struct for given source name
	dbCfg.Name = srcName

	// DEGUG print
	fmt.Println("Iside NewDBConfig() function:")
	fmt.Printf("dbCfg.Name: %s\n", dbCfg.Name)
	fmt.Printf("dbCfg.DisplayName: %s\n", dbCfg.DisplayName)
	fmt.Printf("dbCfg.DriverName: %s\n", dbCfg.DriverName)
	fmt.Printf("dbCfg.Host: %s\n", dbCfg.Host)
	fmt.Printf("dbCfg.Port: %s\n", dbCfg.Port)
	fmt.Printf("dbCfg.DBName: %s\n", dbCfg.DBName)
	fmt.Printf("dbCfg.User: %s\n", dbCfg.User)
	fmt.Printf("dbCfg.Password: %s\n", dbCfg.Password)
	fmt.Printf("dbCfg.SslMode: %s\n", dbCfg.SslMode)
	fmt.Println("------------------")

	return dbCfg, nil
}

func main() {
	path := "/home/raf/Projects/log-tracker/api/config/db-config.json"
	srcName := "test_server"

	dbCfg, err := NewDBConfig(path, srcName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Before changing some fields:\n")
	fmt.Println("dbCfg.Host:", dbCfg.Host)
	fmt.Println("dbCfg.Port:", dbCfg.Port)

	fmt.Println(dbCfg.DisplayName)
	fmt.Printf("Type of dbCfg.DB: %T\n", dbCfg)

	// try to change some fields
	dbCfg.Host = "112.200.12.129"
	dbCfg.Port = "8432"

	fmt.Printf("After changing some fields:\n")
	fmt.Println("dbCfg.Host:", dbCfg.Host)
	fmt.Println("dbCfg.Port:", dbCfg.Port)
}
