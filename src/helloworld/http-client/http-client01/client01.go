// https://networkbit.ch/golang-http-client/
package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func main() {
  resp, err := http.Get("http://www.example.com")
  if err != nil {
    log.Fatal("Error getting response: ", err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal("Error reading response: ", err)
  }
  fmt.Printf("%s\n", body)

  for key, value := range resp.Header {
    for k, v := range value {
      fmt.Printf("[%s] - [%v] : %s\n", key, k, v)
    }
  }
}
