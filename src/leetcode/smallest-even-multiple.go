//https://leetcode.com/problems/smallest-even-multiple/
package main

import "fmt"

func smallestEvenMultiple(n int) int {
  var result int
  if n % 2 == 0 {
    result = n
  } else {
    for i := n; 1 <= i && i <= 150; i++ {
      if (i % 2 == 0 && i % n == 0) {
        i = result
        break
      } else {
        continue
      }
    }
  }
  return result
}

func main() {
  var number int
  number = 16
  result := smallestEvenMultiple(number)
  fmt.Println(result)
}
