// https://www.alexedwards.net/blog/serving-static-sites-with-go
package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", fs)

  log.Print("Listening on :5000...")
  err := http.ListenAndServe(":5000", nil)
  if (err != nil) {
    log.Fatal(err)
  }
}
