package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func saveToJSON(file *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(file)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{
			Telephone{Mobile: true, Number: "1234-4321"},
			Telephone{Mobile: true, Number: "6864-9121"},
			Telephone{Mobile: false, Number: "adrc-43g4"},
			Telephone{Mobile: true, Number: "1122-4499"},
		},
	}

	filename := "writeToThisFile.json"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("opening file:", err.Error())
		os.Exit(1)
	}
	defer file.Close()
	saveToJSON(file, myRecord)
}
