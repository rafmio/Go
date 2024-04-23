package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // import PostgreSQL driver
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

var logEntry1 LogEntry = LogEntry{
	TmStmp: time.Date(2024, time.May, 8, 10, 11, 22, 0, time.UTC),
	SrcIP:  "213.194.150.12",
	Len:    "44",
	Ttl:    "240",
	Id:     "54321",
	Spt:    "44496",
	Dpt:    "2099",
}

var logEntry2 LogEntry = LogEntry{
	TmStmp: time.Date(2023, time.June, 11, 12, 13, 14, 0, time.UTC),
	SrcIP:  "210.220.30.140",
	Len:    "44",
	Ttl:    "249",
	Id:     "54321",
	Spt:    "55098",
	Dpt:    "50067",
	Window: "65535",
}

func main() {
	db, err := sql.Open("postgres", "user=raf dbname=logtracker password=qwq121 sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Подготовка SQL запросов для вставки
	stmt, err := db.Prepare("INSERT INTO log_entries (tmstmp, srcip, len, ttl, innerid, spt, dpt, wndw) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(logEntry1.TmStmp, logEntry1.SrcIP, logEntry1.Len, logEntry1.Ttl, logEntry1.Id, logEntry1.Spt, logEntry1.Dpt, logEntry1.Window)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec(logEntry2.TmStmp, logEntry2.SrcIP, logEntry2.Len, logEntry2.Ttl, logEntry2.Id, logEntry2.Spt, logEntry2.Dpt, logEntry2.Window)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Success!")
}
