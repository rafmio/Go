package main

import (
  "net/http"
  "fmt"
  "strconv"
)

func GetAndSetCookie(w http.ResponseWriter, r *http.Request) {
  counterVal := 1
  counterCookie, err := r.Cookie("counter")
  if (err == nil) {
    counterVal, _ = strconv.Atoi(counterCookie.Value)
    counterVal++
  }

  http.SetCookie(w, &http.Cookie {
    Name: "counter", Value: strconv.Itoa(counterVal),
  })

  if(len(r.Cookies()) > 0) {
    for _, c := range r.Cookies() {
      fmt.Fprintf(w, "Cookie Name: %v, Value: %v", c.Name, c.Value)
    }
  } else {
    fmt.Fprintln(w, "Request contains no cookies")
  }
}

func init() {
  http.HandleFunc("/cookies", GetAndSetCookie)
}
