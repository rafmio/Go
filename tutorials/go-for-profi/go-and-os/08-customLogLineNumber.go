/*
How to specify in the log file the line number of the source file that executed
the instruction that made the log entry.

	#номер строки файла
*/
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
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
