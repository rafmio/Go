package dbops

import "database/sql"

func EstablishDBConnection() (*sql.DB, error) {
	dataSourceName := "user=randomuser password=mypassword dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
