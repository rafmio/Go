// https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Pet struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
}

func main() {
	dogs := []Pet{
		{
			Name:   "Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  "German Shepherd/Pitbull",
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "14 years, 3 months",
			Breed:  "German Shepherd Collie",
		},
		{
			Name:   "Bruce Wayne",
			Sex:    "Male",
			Intact: false,
			Age:    "3 years, 7 months",
			Breed:  "Chihuahua",
		},
	}

	var tmpFile = "16-pets.tmpl"
	tmpl, err := template.New(tmpFile).ParseFiles(tmpFile)
	if err != nil {
		fmt.Println("Error creating new template:", err.Error())
	}

	err = tmpl.Execute(os.Stdout, dogs)
	if err != nil {
		fmt.Println("Error executing template:", err.Error())
	}
}
