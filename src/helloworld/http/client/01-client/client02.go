// https://networkbit.ch/golang-http-client/
package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"
)

func main() {
  userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 YaBrowser/22.11.3.838 Yowser/2.5 Safari/537.36"
  req, err := http.NewRequest("GET", "http://example.com", nil)
  if err != nil {
    log.Fatal("Error reading request: ", err)
  }
  req.Header.Set("Cache-Control", "no-cache")
  req.Header.Set("User-Agent", userAgent)
  client := &http.Client{Timeout: time.Second * 10}

  resp, err := client.Do(req)
  if err != nil {
    log.Fatal("Error reading response: ", err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal("Error reading body: ", err)
  }

  fmt.Printf("%s\n", body)
}
