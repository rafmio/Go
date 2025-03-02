
package main

import (
  "fmt"
  "html"
  "net/http"
  "strconv"
  "strings"
  "time"

  "example/romanNumbers"
)


func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    urlPathElements := strings.Split(r.URL.Path, "/")

    // if request is GET with correct syntax
    if urlPathElements[1] == "roman_number" {
      number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
      if number == 0 || number > 10 {
        // if resource is not in the list, send Not Found status
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("404 - Not Found"))
      } else {

        w.WriteHeader(http.StatusOK)
        for name, values := range r.Header {
          for _, value := range values {
            fmt.Fprint(w, "%s : %s", name, value)
          }
        }
        fmt.Fprintf(w, "%q",
        html.EscapeString(romanNumbers.Numerals[number]))
        fmt.Fprintf(w, "\n")
      }
    } else {
      // for all other requests, tell that Client sent a bad request
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte("400 - Bad request"))
    }
  })

  // Create a server and run it on 8000 port
  s := &http.Server{
    Addr: ":8000",
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  s.ListenAndServe()
}
