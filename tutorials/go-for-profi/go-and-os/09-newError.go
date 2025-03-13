package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

var LOGFILE = "/tmp/mGo.log"

func main() {
	file, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer file.Close()

	iLog := log.New(file, "customLog ", log.LstdFlags)
	// log.Lshortfile - file name element and line number: d.go:23
	iLog.SetFlags(log.LstdFlags | log.Lshortfile) // log.Lshortfile - file name element and line number: d.go:23

	a := rand.Intn(10)
	b := rand.Intn(10)

	err = returnError(a, b)
	if err == nil {
		iLog.Println("returnError() ended normally!")
		fmt.Println("returnError() ended normally!")
	} else {
		iLog.Printf("Error in returnError(): %s\n", err)
		fmt.Printf("Error in returnError(): %s\n", err)
	}

}
