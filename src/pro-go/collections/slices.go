package main

import "fmt"

func main() {
  names := make([]string, 3)
  names[0] = "Kayak"
  names[1] = "Lifejacket"
  names[2] = "Paddle"
  fmt.Println(names)
  for index, value := range names {
    fmt.Println(index, ":", value)
  }

  names2 := []string{"Kissy", "Missy", "Huggy", "Wuggy"}
  fmt.Println(names2)
  for _, value := range names2 {
    fmt.Println(value)
  }
  fmt.Println()
  names2 = append(names2, "Tosy", "Bosy")
  fmt.Println(names2)
  for i, val := range names2 {
    fmt.Println(i, ":", val)
  }
  fmt.Println()

  appendedNames := append(names2, "Cheeackey", "Peakey")
  for index, value := range appendedNames{
    fmt.Println(index, ":", value)
  }
  fmt.Println("names2", names2)
  fmt.Println("appendedNames", appendedNames)
  fmt.Println("len(appendedNames):", len(appendedNames))
  fmt.Println("cap(appendedNames):", cap(appendedNames))
  fmt.Println()

  appendedNames = append(appendedNames, "Hat", "Glovers")
  appendedNames[0] = "Canoe"
  fmt.Println("names2", names2)
  fmt.Println("appendedNames", appendedNames)
  fmt.Println()

  fmt.Printf("names: %p\n", &names)
  fmt.Printf("names2: %p\n", &names2)
  fmt.Printf("appendedNames: %p\n", &appendedNames)
  fmt.Println()


}

// func make(t Type, size ...IntegerType) Type
