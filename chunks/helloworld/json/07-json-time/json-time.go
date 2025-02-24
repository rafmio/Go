package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type OfficeDay struct {
	ID        int    `json:"id"`
	Day       string `json:"day"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

type OfficeHoursData struct {
	OfficeHours []*OfficeDay `json:"office_hours"`
}

func ReadJSONFile(filePath string) ([]*OfficeDay, error) {
	var officeHoursData OfficeHoursData

	log.Println("Try to open JSON-file...")
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error opening JSON-file:", err)
		return nil, err
	} else {
		log.Println("JSON-file has been opened successfully")
	}

	err = json.Unmarshal(data, &officeHoursData)
	if err != nil {
		log.Println("Error unmarshalling JSON-data:", err)
		return nil, err
	} else {
		log.Println("JSON-data has been successfully unmarshalled")
	}

	return officeHoursData.OfficeHours, nil
}

func (o *OfficeDay) ValidateTimeFormat() bool {
	fmt.Println("day to test:", o.Day)
	var validDays = []string{"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday"}

	var dayIsValid bool

	for _, validDay := range validDays {
		if strings.EqualFold(o.Day, validDay) {
			dayIsValid = true
			break
		}
	}
	if !dayIsValid {
		log.Println("Invalid day:", o.Day)
		return false
	}

	const layout = "15:04:05" // Format HH:MM

	openTime, err := time.Parse(layout, o.OpenTime)
	if err != nil {
		log.Println("Error parsing open time")
		return false
	}

	closeTime, err := time.Parse(layout, o.CloseTime)
	if err != nil {
		log.Println("Error parsing close time")
		return false
	}

	if openTime.After(closeTime) {
		log.Println("open time should be before close time")
		return false
	}

	return true
}

func main() {
	filePath := "office_hours.json"

	officeDays, err := ReadJSONFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, day := range officeDays {
		log.Printf("Day: %s, Open time: %s, Close time: %s\n", day.Day, day.OpenTime, day.CloseTime)
	}

	for _, day := range officeDays {
		if day.ValidateTimeFormat() {
			log.Println("All times in the correct format.")
		}
	}
}
