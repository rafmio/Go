// https://pkg.go.dev/encoding/json#Marshal
package main

import (
  "encoding/json"
  "fmt"
  "os"
)

func main() {
  type ColorGroup struct {
    ID int
    Name string
    Colors []string
  }

  group := ColorGroup {
    ID: 1,
    Name: "Reds",
    Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
  }

  b, err := json.Marshal(group)
  fmt.Printf("Type of b: %T\n", b)
  fmt.Println()
  
  if err != nil {
    fmt.Println("error:", err)
  }
  os.Stdout.Write(b)
  fmt.Println()

}
