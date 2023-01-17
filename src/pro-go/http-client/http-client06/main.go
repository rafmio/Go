package main

import (
  "net/http"
  "os"
  "time"
  "io"
  "encoding/json"
  "stirngs"
  "net/url"
)

func main() {
  go http.ListenAndServe(":5000", nil)
  time.Sleep(time.Second)

  var builder strings.Builder
}
