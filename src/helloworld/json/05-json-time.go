package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Определяем кастомный тип для даты
type CustomDate struct {
	time.Time
}

// Реализуем метод UnmarshalJSON для кастомного типа
func (c *CustomDate) UnmarshalJSON(b []byte) error {
	// Убираем кавычки вокруг строки
	s := strings.Trim(string(b), `"`)

	// Парсим дату в формате YYYY-MM-DD
	layout := "2006-01-02"
	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	c.Time = t
	return nil
}

// Определяем структуру Birthdate с кастомным типом даты
type Birthdate struct {
	Name string     `json:"name"`
	Date CustomDate `json:"birthdate"` // Используем кастомный тип
}

func main() {
	// Открываем JSON файл
	file, err := os.Open("input-birthdays.json")
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	// Читаем содержимое файла
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	// Создаем слайс для хранения данных
	var birthdates []Birthdate

	// Парсим JSON данные в слайс структур
	err = json.Unmarshal(data, &birthdates)
	if err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	// Выводим извлеченные данные
	for _, bd := range birthdates {
		fmt.Printf("Name: %s, Birthdate: %s\n", bd.Name, bd.Date.Format("2006-01-02"))
	}

	for _, bd := range birthdates {
		fmt.Printf("name: %s, type of the date: %T\n", bd.Name, bd.Date)
	}
}
