package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"
)

func responseSize(url string) {
  fmt.Println("Getting", url)
  response, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(url, "\t\t", len(body))
}

func main() {
  go responseSize("https://example.com/")
  go responseSize("https://golang.org/")
  go responseSize("https://golang.org/doc")
  go responseSize("http://himpostavka.ru/")
  go responseSize("https://www.google.com/")
  time.Sleep(5 * time.Second)
}
