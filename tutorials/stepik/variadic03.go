package main

import "fmt"

func calcRes(results ...string) {
  var total int
  for _, v := range results {
    if v == "w" {
      total += 3
    } else if v == "d" {
      total += 1
    } else {
      total += 0
    }
  }
  fmt.Println(total)
}

func main() {
  results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

  var input string
  fmt.Scanln(&input)

  results = append(results, input)

  calcRes(results...)
}
