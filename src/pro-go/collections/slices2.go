package main

import "fmt"

func main() {
  var arr [5]int = [5]int{1, 2, 3, 4, 5}
  fmt.Println(arr)
  for index, value := range arr {
    fmt.Println(index, value)
  }
  sls := arr[0:5]
  fmt.Println(sls)
  arr[0] = 100
  fmt.Println(arr)
  fmt.Println(sls)

  sls[1] = 200
  fmt.Println(arr)
  fmt.Println(sls)
  sls = append(sls, 600, 700)
  fmt.Println(arr)
  fmt.Println(sls)

  arr[2] = 300
  fmt.Println(arr)
  fmt.Println(sls)

  fmt.Println()
  fmt.Printf("Type of arr: %T\n", arr)
  fmt.Printf("Type of sls: %T\n", sls)
}
