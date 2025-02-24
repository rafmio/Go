package main
import "fmt"

func main() {
  var s []int
  printSlice(s)

  s = append(s, 1)
  printSlice(s)

  s = append(s, 2, 3, 4, 100)
  printSlice(s)

  var sls []int
  printSlice(sls)

  sls = append(s, 200, 300, 400, 600)
  printSlice(s)
  printSlice(sls)
}

func printSlice(s []int) {
  fmt.Printf("len = %d, cap = %d %v \n", len(s), cap(s), s)
}

// Append to a slice
// https://go.dev/tour/moretypes/15
