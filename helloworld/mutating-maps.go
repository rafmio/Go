package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("The value: ", m["Answer"])

  m["Answer"] = 48
  fmt.Println("The value: ", m["Answer"])

  m["Question"] = 150
  fmt.Println("The value: ", m["Question"])

  m["Kissy"] = 1
  m["Missy"] = 2

  v, ok := m["Answer"]
  fmt.Println("The value: ", v, "Present? ", ok)

  vv, oo := m["Question"]
  fmt.Println("The value: ", vv, "Present: ", oo)

  _, oo2 := m["Quoris"]
  fmt.Println("The value: 'Quoris'", "Present: ", oo2)

  m["Huggy"] = 3
  m["Wuggy"] = 4
  m["Tosi"]  = 5
  m["Bosi"]  = 6

  fmt.Println()
  fmt.Println("m: ", m)

  fmt.Println("m[Bosi]", m["Bosi"])

  for _, vl := range m {
    fmt.Println(vl)
  }

  _, val := m["Missy"]
  fmt.Println("Is val 'Missy' present? : ", val)
}


// Mutating maps
// https://go.dev/tour/moretypes/22
// Insert or update an element in map 'm':
// m[key] = elem
// Retrieve an element:
// elem = m[key]
// Delete an element:
// delete(m, key)
// Test that a key is present with a two-value assignment:
// elem, ok = m[key]
