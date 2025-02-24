package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])

	// syslog.New() establishes a new connection to the system log daemon
	// returns *Writer, err
	sysLog, err := syslog.New(syslog.LOG_INFO | syslog.LOG_LOCAL7, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		// SetOutput() sets the output destination for the standard logger
		log.SetOutput(sysLog)
	}

	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	pid := os.Getpid()
	log.Println("Current PID:", pid)

	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
}
