package main

import (
  "fmt"
  "sumfunc"
)

func main() {
  var a int = 10
  b := 20
  var c int 
  c = sumfunc.Summ(a, b)
  fmt.Println(c)
  
  dd := sumfunc.Summ(150, 150) + sumfunc.ExportedVar
  fmt.Println(dd)
}
