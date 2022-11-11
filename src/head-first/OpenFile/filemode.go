package main

import (
  "fmt"
  "os"
)


func main() {
  fmt.Println(os.FileMode(0700))
  fmt.Println(os.FileMode(0070))
  fmt.Println(os.FileMode(0007))
  fmt.Println(os.FileMode(0777))

  fmt.Println()

  octBase := 0000
  octStep := 0111
  for octVal := octBase; octVal <= 0777; octVal = octVal + octStep {
    fmt.Printf("%3o %09b %s\n", octVal, octVal, os.FileMode(octVal))
  }
}
// Преобразовал в цикл следующий код
// fmt.Printf("%09b %s\n", 0000, os.FileMode(0000))
// fmt.Printf("%09b %s\n", 0111, os.FileMode(0111))
// fmt.Printf("%09b %s\n", 0222, os.FileMode(0222))
// fmt.Printf("%09b %s\n", 0333, os.FileMode(0333))
// fmt.Printf("%09b %s\n", 0444, os.FileMode(0444))
// fmt.Printf("%09b %s\n", 0555, os.FileMode(0555))
// fmt.Printf("%09b %s\n", 0666, os.FileMode(0666))
// fmt.Printf("%09b %s\n", 0777, os.FileMode(0777))
