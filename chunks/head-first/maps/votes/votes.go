package main

import (
    "fmt"
    "votesdata"
    "log"
)

func main() {
    lines, err := votesdata.GetStrings("votes.txt")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(lines)
}
