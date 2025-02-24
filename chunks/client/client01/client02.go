package main

import (
  "net/http"
  "log"
  "io/ioutil"
  "os"
)

func main() {
  MakeRequest()
}

func MakeRequest() {
  resp, err := http.Get("https://гибдд.рф/check/fines")
  if (err != nil) {
    log.Fatalln(err)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if (err != nil) {
    log.Fatalln(err)
  }
  // log.Println(string(body))
  file, err := os.Create("gibdd.html")
  if (err != nil) {
    log.Fatalln("Unable to create file:", err)
  }
  defer file.Close()
  file.Write(body)
}
