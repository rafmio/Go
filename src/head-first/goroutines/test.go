package main

import "fmt"

func main() {
  channel := make(chan string)
  channel = go initChan(channel)
  printChan(channel)
}


func initChan(s chan string) {
  s <- "Tosi-Bosi"
}

func printChan(s chan string) {
  fmt.Println(s)
}
