package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGo.log"

func one(aLog *log.Logger) {
	aLog.Println("--FUNCTION one ------")
	defer aLog.Println("--FUNCTION one (defer) ------")

	for i := 0; i < 5; i++ {
		aLog.Println(i)
	}
}

func two(aLog *log.Logger) {
	aLog.Println("---- FUNCTION two")
	defer aLog.Println("FUNCTION two (defer) ------")

	for i := 10; i > 0; i-- {
		aLog.Println(i)
	}
}

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0655)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "again-logDefer ", log.LstdFlags)
	iLog.Println("Hello-mello there!")
	iLog.Println("Another log entry")

	one(iLog)
	two(iLog)
}
