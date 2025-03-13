package main

import "fmt"

func main() {
  counters := map[string]int{"a" : 3, "b" : 0, "d" : 4}
  var value int
  var ok bool
  value, ok = counters["a"]
  fmt.Println(value, ok)
  value, ok = counters["b"]
  fmt.Println(value, ok)
  value, ok = counters["c"]
  fmt.Println(value, ok)

  _, ok = counters["b"]
  fmt.Println(ok)
  _, ok = counters["c"]
  fmt.Println(ok)
  _, ok = counters["d"]
  fmt.Println(ok)
}

// Проверяет пристутствие в карте значения
