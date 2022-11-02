package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func main() {
  responseSize("https://exapmle.com/")
  responseSize("https://golang.org/")
  responseSize("https://golang.org/doc")
  responseSize("http://himpostavka.ru/")
  responseSize("https://factory5.ai/")
}

func responseSize(url string) {
  response, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }

  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(url, "\t", len(body))
}
