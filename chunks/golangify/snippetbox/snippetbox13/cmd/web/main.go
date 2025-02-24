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

  logFile, err := os.OpenFile("info.log", os.O_RDWR | os.O_CREATE, 0666)
  if err != nil {
    log.Fatal(err)
  }
  defer logFile.Close()

  infoLog := log.New(logFile, "info\t", log.Ldate | log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)

  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  // Инизиализируем новую структуру http.Server
  srv := &http.Server {
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: mux,
  }


  infoLog.Printf("Launching server on %s", *addr)
  err = srv.ListenAndServe()
  errorLog.Fatal(err)
}
