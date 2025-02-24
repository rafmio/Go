// the code doesn't work!
package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type LogEntry struct {
	TmStmp time.Time
	SrcIP  string
	Len    string
	Ttl    string
	Id     string // will named 'inner id' in database
	Spt    string
	Dpt    string
	Window string // will named 'wndw' in database
}

func main() {
	db, err := sql.Open("postgres", "host=194.58.102.129 port=5432 dbname=logtracker user=raf password=qwq121 sslmode=disable")
	if err != nil {
		log.Println("Open DB:", err)
		log.Fatalf("Open DB: %v", err)
		return
	}
	defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Println("Ping DB:", err)
	// 	return
	// } else {
	// 	fmt.Println("Successfully connected!")
	// }

	rows, err := db.Query("SELECT * FROM logs WHERE tmstmp BETWEEN '2022-06-01' AND '2022-06-30'")
	if err != nil {
		log.Println("Query DB:", err)
		return
	}

	defer rows.Close()

	logEnties := make([]LogEntry, 0)

	for rows.Next() {
		// var logID, level, message string
		// err = rows.Scan(&logID, &level, &message)
		// if err != nil {
		// 	log.Println("Scan row:", err)
		// 	return
		// }
		// fmt.Printf("ID: %s, Level: %s, Message: %s\n", logID, level, message)
		var entry LogEntry
		err = rows.Scan(&entry.TmStmp, &entry.SrcIP, &entry.Len, &entry.Ttl, &entry.Id, &entry.Spt, &entry.Dpt, &entry.Window)
		if err != nil {
			log.Println("Scan row:", err)
			return
		}
		logEnties = append(logEnties, entry)
	}
}
