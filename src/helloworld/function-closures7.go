package main

import (
    "fmt"
    "strings"
)

func wordList () func (string) string { 
    var sliceOfWords[] string
    
    return func (newWord string) string {
        sliceOfWords = append(sliceOfWords, newWord)
        return strings.Join(sliceOfWords, "--")
    }
}

func main() {
    var closure1 = wordList()
    fmt.Println("closure1 = ", closure1("apple"))
    fmt.Println("closure1 = ", closure1("ninja"))
    fmt.Println("closure1 = ", closure1("mad coder"))
    fmt.Println("closure1 = ", closure1("ba-ba-ba-nana"))
}

// Funcion closures 
// https://www.golearningsource.com/fundamentals/function-closures-in-golang/
