// Write logs to file
package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGo.log"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags)
	iLog.Println("Hello-mello, Tosy-Bosy, Huggy-Wuggy, Kissy-Missy")
	iLog.Printf("This is a printf-formatted message: %s", "Hello, Go!")
}
