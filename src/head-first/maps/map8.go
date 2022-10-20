package main

import (
  "fmt"
  "sort"
)

func main() {
  grades := map[string]float64{"Daniel": 70.1,
                               "Alma": 74.2,
                               "Rohit": 86.5,
                               "Carl": 59.7,
                             }
  var names []string
  for name := range grades {
    names = append(names, name)
  }

  fmt.Println(names)
  sort.Strings(names)
  fmt.Println(names)
  
  for _, name := range names {
    fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
  }
}
