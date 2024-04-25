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
	SrcIP:  "10.191.114.13",
	Len:    "40",
	Ttl:    "266",
	Id:     "54123",
	Spt:    "8080",
	Dpt:    "90",
	Window: "65535",
}

var logEntry2 LogEntry = LogEntry{
	TmStmp: time.Date(2023, time.June, 11, 12, 13, 14, 0, time.UTC),
	SrcIP:  "110.120.32.32",
	Len:    "60",
	Ttl:    "250",
	Id:     "54321",
	Spt:    "54340",
	Dpt:    "40040",
	Window: "65535",
}

var logEntry3 LogEntry = LogEntry{
	TmStmp: time.Date(2023, time.June, 11, 12, 13, 14, 0, time.UTC),
	SrcIP:  "58.132.112.102",
	Len:    "42",
	Ttl:    "240",
	Id:     "54443",
	Spt:    "443",
	Dpt:    "44343",
	Window: "11111",
}

func main() {
	db, err := sql.Open("postgres", "user=raf dbname=logtracker password=qwq121 sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	logEntries := []LogEntry{logEntry1, logEntry2, logEntry3}

	for _, el := range logEntries {
		isExists := CheckRecordExists(db, el)
		if isExists {
			fmt.Println("record exists")
		} else {
			fmt.Println("record doesn't exist")

			// Подготовка SQL запросов для вставки
			stmt, err := db.Prepare("INSERT INTO lg_tab_1 (tmstmp, srcip, len, ttl, innerid, spt, dpt, wndw) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
			if err != nil {
				log.Fatal(err.Error())
			}
			defer stmt.Close()

			_, err = stmt.Exec(el.TmStmp, el.SrcIP, el.Len, el.Ttl, el.Id, el.Spt, el.Dpt, el.Window)
			if err != nil {
				log.Fatal(err.Error())
			}

		}
	}

	fmt.Println("Success!")
}

func CheckRecordExists(db *sql.DB, logEntry LogEntry) bool {
	// Подготовка SQL запроса для проверки существования записи
	stmt, err := db.Prepare("SELECT * FROM lg_tab_1 WHERE tmstmp = $1 AND srcip = $2 AND len = $3 AND ttl = $4 AND innerid = $5 AND spt = $6 AND dpt = $7 AND wndw = $8")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	// Выполнение запроса
	rows, err := stmt.Query(logEntry.TmStmp, logEntry.SrcIP, logEntry.Len, logEntry.Ttl, logEntry.Id, logEntry.Spt, logEntry.Dpt, logEntry.Window)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	rowsSls, err := rows.Columns()
	for i, row := range rowsSls {
		fmt.Println(i, row)
	}

	return rows.Next()
}
