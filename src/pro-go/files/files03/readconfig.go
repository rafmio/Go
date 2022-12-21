package main

import (
	"os"
	"encoding/json"
)

type ConfigData struct {
	UserName string
	AdditionalProducts []Product
}

var Config ConfigData

func LoadConfig() (err error) {
	file, err := os.Open("config.json")
	if (err == nil) {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config)
	}
	return
}

func init() {
	err := LoadConfig()
	if (err != nil) {
		Printfln("Error Loading Config: %v", err.Error())
	} else {
		Printfln("Username: %v", Config.UserName)
		Products = append(Products, Config.AdditionalProducts...)
	}
}
