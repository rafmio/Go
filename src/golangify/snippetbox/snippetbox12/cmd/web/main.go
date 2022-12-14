package main

import (
  "flag"
  "log"
  "net/http"
  "os"
)

func main() {
  addr := flag.String("addr", ":4000", "Web-server's net address")
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
  errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate | log.Ltime)

  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  infoLog.Printf("Launching the server on %s", *addr)
  err := http.ListenAndServe(*addr, mux)
  errorLog.Fatal(err)

}
