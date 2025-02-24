package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    fileInfo, err := os.Stat("00-head-first-book.txt")
    
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(fileInfo.Size())
}
