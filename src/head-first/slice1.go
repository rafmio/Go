package main

import "fmt"

func main() {
  notes := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
  fmt.Println(notes[3], notes[6], notes[0])
  primes := []int{
    2,
    3,
    5,
  }

  fmt.Println(primes[0], primes[1], primes[2])

  fmt.Println()
  fmt.Println("notes: ", notes)
  fmt.Println("primes: ", primes)
}
