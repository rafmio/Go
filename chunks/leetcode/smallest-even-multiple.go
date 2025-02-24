//https://leetcode.com/problems/smallest-even-multiple/
package main

import "fmt"

func smallestEvenMultiple(n int) int {
  var result int = 0
  for i := (n + 1); i <= 150; i++ {
    if ((i % 2) == 0 && (i % n) == 0) {
      result = i
      break
    }
  }
  return result
}

func main() {
  input := 8
  result := smallestEvenMultiple(input)
  fmt.Println(result)

}
