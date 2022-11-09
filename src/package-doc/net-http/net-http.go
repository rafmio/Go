package main

import (
  "fmt"
  "net/http"
  "io"
  "os"
  "log"
)

func main() {
  resp, err := http.Get("http://example.com/")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  fmt.Println(resp)
  fmt.Printf("Type of resp: %T\n", resp)
  body, err := io.ReadAll(resp.Body)
  fmt.Println()
  fmt.Println(string(body))

  err = os.WriteFile("examleCom.html", body, 0666)
  if err != nil {
    log.Fatal(err)
  }
}
