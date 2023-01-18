package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  // resp, err := http.Get("http://himpostavka.ru")
  resp, err := http.Get("https://гибдд.рф/check/fines")
  if (err != nil) {
    log.Fatal(err)
  } else {
    fmt.Println(resp.Status)
    fmt.Println(resp.StatusCode)
    for key, value := range resp.Header {
      for k, v := range value {
        fmt.Println(key, " : ", k, "---", v)
      }
    }
  }

}
