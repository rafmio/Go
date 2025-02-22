package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type LogEntry struct {
	SeqNum string
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
	// Параметры подключения
	connStr := "host=194.58.102.129 port=5432 user=raf password=qwq121 dbname=logtracker sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при открытии базы данных: %v\n", err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v\n", err)
	}
	fmt.Println("Успешно подключено к базе данных!")

	// Выполнение запроса
	rows, err := db.Query("SELECT * FROM lg_tab WHERE tmstmp BETWEEN $1 AND $2", "2024-08-19", "2024-08-20")
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v\n", err)
	}
	defer rows.Close()

	// Обработка результатов
	for rows.Next() {
		var entry LogEntry
		err = rows.Scan(&entry.SeqNum, &entry.TmStmp, &entry.SrcIP, &entry.Len, &entry.Ttl, &entry.Id, &entry.Spt, &entry.Dpt, &entry.Window) // Добавьте другие поля, если они есть в таблице lg_tab
		if err != nil {
			log.Fatalf("Ошибка сканирования строки: %v\n", err)
		}
		// fmt.Printf("Запись: %v\n", entry.TmStmp) // Вывод данных на экран
		fmt.Printf("%v %v %v %v %v %v %v %v %v\n", entry.SeqNum, entry.TmStmp, entry.SrcIP, entry.Len, entry.Ttl, entry.Id, entry.Spt, entry.Dpt, entry.Window)
	}

	// Проверка на ошибки после завершения итерации
	if err = rows.Err(); err != nil {
		log.Fatalf("Ошибка при обработке строк: %v\n", err)
	}
}
