package main

import (
  "encoding/json"
  "fmt"
  "os"
  // "log"
)

type Account struct {
  Wallet string
  Password string
}

var accouts []Account

func main() {
  acc1 := Account{"asdfl;kjasdfl;jasdf93fl;ldsaasf", "qwerty"}
  json_data, err := json.Marshal(acc1)
  if err != nil {
    fmt.Println("marshal data: ", err.Error())
  }

  file, err := os.OpenFile("password.json", os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0644)
  if err != nil {
    fmt.Println("opening file: ", err.Error())
  }

  defer file.Close()

  _, err = file.WriteString(string(json_data))
  if err != nil {
    fmt.Println("writing to file: ", err.Error())
  }

  fmt.Println()

  file2, err := os.ReadFile("password.json")
  if err != nil {
    fmt.Println("reading file: ", err.Error())
  }

  err = json.Unmarshal(file2, &accouts)
  if err != nil {
    fmt.Println("unmarshal file: ", err.Error())
  }

  for i, v := range accouts {
    fmt.Println(i, ":", v)
  }

}
