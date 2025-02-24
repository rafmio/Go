package main

import (
    "encoding/json"
    "fmt"
    "os"
    "log"
)

type User struct {
    Id         int
    Name       string
    Occupation string
}

func main() {
    u1 := User{1, "John Doe", "gardener"}
    json_data, err := json.Marshal(u1)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(json_data))

    users := []User{
        {Id: 2, Name: "Roger Roe", Occupation: "driver"},
        {Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
        {Id: 4, Name: "David Brown", Occupation: "programmer"},
    }

    json_data2, err := json.Marshal(users)
    if err != nil {
        log.Fatal(err)
    }

    err = os.WriteFile("users.json", json_data2, 0644)
    if err !=  nil {
      fmt.Println("writing to file: ", err.Error())
    }

    fmt.Println(string(json_data2))
    fmt.Println()

    file, err := os.ReadFile("users.json")
    if err != nil {
      fmt.Println("reading file: ", err.Error())
    }

    var unmarshUsers []User
    err = json.Unmarshal(file, &unmarshUsers)
    if err != nil {
      fmt.Println("unmarshalling file: ", err.Error())
    }

    for i, v := range unmarshUsers {
      fmt.Println(i, " : ", v)
      if unmarshUsers[i].Name == "Roger Roe" {
        fmt.Println("GOTCHA!!!")
      }
    }

}
