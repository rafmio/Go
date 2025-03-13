package main

import "fmt"

func main() {
  data := []string{"a", "c", "e", "a", "e"}
  counts := make(map[string]int)

  fmt.Println(data)
  fmt.Println(counts)

  for _, item := range data {
    counts[item]++
  }

  fmt.Println()
  fmt.Println(data)
  fmt.Println(counts)

  letters := []string{"a", "b", "c", "d", "e", "x"}
  for _, letter := range letters {
    count, ok := counts[letter]
    if !ok {
      fmt.Printf("%s: not found\n", letter)
    } else {
      fmt.Printf("%s: %d\n", letter, count)
    }
  }

  fmt.Println()
  sls := []int{0, 0, 0}
  fmt.Println(sls)
  sls[0]++
  sls[0]++
  sls[0]++
  sls[1]++
  fmt.Println(sls)
}
