package main

import (
  "net/http"
  "io"
  "strings"
)

type StringHandler struct {
  message string
}

func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  Printfln("Request for %v", r.URL.Path)
  io.WriteString(w, sh.message)

  for key, value := range r.Header {
    io.WriteString(w, key)
    io.WriteString(w, " : ")
    for _, val := range value {
      io.WriteString(w, val)
    }
    io.WriteString(w, "\n")
  }
}

func HTTPRedirect(w http.ResponseWriter, r *http.Request) {
  host := strings.Split(r.Host, ":")[0]
  target := "https://" + host + ":5500" + r.URL.Path
  if len(r.URL.RawQuery) > 0 {
    target += "?" + r.URL.RawQuery
  }
  http.Redirect(w, r, target, http.StatusTemporaryRedirect)
}

func main() {
  http.Handle("/message", StringHandler{message: "Hello again! It's me"})
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

  fsHandler := http.FileServer(http.Dir("./static"))
  http.Handle("/files/", http.StripPrefix("/files", fsHandler))

  go func() {
    err := http.ListenAndServeTLS(":5500", "certificate.crt", "certificate.key", nil)
    if (err != nil) {
      Printfln("HTTPS Error: %v", err.Error())
    }
  }()

  err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPRedirect))
  if (err != nil) {
    Printfln("Error: %v", err.Error())
  }
}
