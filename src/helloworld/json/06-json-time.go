package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// Структура для хранения имени и даты рождения
type Birthdate struct {
	Name       string `json:"name"`
	DateString string `json:"birthdate"`
	DateTime   time.Time
}

func (b *Birthdate) ParseDate() {
	layout := "2006-01-02"
	t, err := time.Parse(layout, b.DateString)
	if err != nil {
		log.Fatalf("cannot parse the data: %v", err)
	}
	b.DateTime = t
}

func main() {
	// Чтение содержимого файла
	file, err := os.Open("input-birthdays.json") // Убедитесь, что файл находится в нужной директории
	if err != nil {
		log.Fatalf("cannot open the file: %v", err)
	}
	defer file.Close()

	var birthdates []Birthdate

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&birthdates)
	if err != nil {
		log.Println("error parsing (Decode()):", err)
	}

	for _, date := range birthdates {
		date.ParseDate()
		log.Printf("%s: %s, type of the date field: %T\n", date.Name, date.DateTime.Format("2006-01-02"), date.DateTime)
	}
}
