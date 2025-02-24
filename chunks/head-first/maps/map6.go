package main

import "fmt"

func main() {
  var ok bool
  ranks := make(map[string]int)
  var rank int

  ranks["bronze"] = 3
  rank, ok = ranks["bronze"]
  fmt.Printf("rank: %d, ok: %v\n", rank, ok)

  ranks["gold"] = 1
  rank, ok = ranks["gold"]
  fmt.Printf("rank: %d, ok: %v\n", rank, ok)

  delete(ranks, "bronze")
  rank, ok = ranks["bronze"]
  fmt.Printf("rank: %d, ok: %v\n", rank, ok)

  isPrime := make(map[int]bool)
  var prime bool
  isPrime[5] = true
  prime, ok = isPrime[5]
  fmt.Printf("prime: %v, ok: %v\n", prime, ok)
  fmt.Println("isPrime[4] = ", isPrime[4])
  fmt.Println(isPrime)

  delete(isPrime, 5)
  prime, ok = isPrime[5]
  fmt.Printf("prime: %v, ok: %v\n", prime, ok)
}
