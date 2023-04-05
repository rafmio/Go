package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{
			{Mobile: true, Number: "4321-1234"},
			{Mobile: true, Number: "3564-gfss"},
			{Mobile: false, Number: "sfgg-4545"},
		},
	}
	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(rec))

	var unRec Record
	err1 := json.Unmarshal(rec, &unRec)
	if err != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(unRec)
}
