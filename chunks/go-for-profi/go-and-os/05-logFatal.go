package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	// syslog.New() establishes a new connection to the system log daemon
	// returns *Writer, err
	sysLog, err := syslog.New(syslog.LOG_ALERT | syslog.LOG_MAIL, "Some program")
	if err != nil {
		log.Fatal(err)
	} else {
		// SetOutput() sets the output destination for the standard logger
		log.SetOutput(sysLog)
	}

	log.Fatal(sysLog)
	fmt.Println("Will you see this?")
}
