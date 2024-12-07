package dbops

import (
	"getcode/models"
	"log"
)

func SongsSearchDB(queryParams *models.QueryParams) (*models.SongDetail, error) {
	log.Println("the SongsSearchDB() has been called")

	// try to connect to DB

	dbCfg, err := NewDBConfig("../config/dbconf.env")
	if err != nil {
		log.Println("Error creating DBConfig:", err)
		return nil, err
	}
	err = dbCfg.EstablishDbConnection()
	if err != nil {
		log.Println("Error establishing DB connection:", err)
		return nil, err
	}
	defer dbCfg.DB.Close()

}
