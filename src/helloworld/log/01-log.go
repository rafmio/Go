package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)
	log.Println("Application started")

	text, err := os.ReadFile("/home/raf/log-tracker/ufw-1-line.log")
	if err != nil {
		log.Println("ERROR: os.ReadFile():", err.Error())
	} else {
		log.Println("file was read")
	}

	fmt.Fprintf(os.Stdout, "%s", text)
}
