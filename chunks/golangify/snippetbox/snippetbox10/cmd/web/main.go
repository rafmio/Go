package main

import (
  "log"
  "net/http"
  "path/filepath"
)

type neuteredFileSystem struct {
  fs http.FileSystem
}

func main() {
  mux := http.NewServeMux() // Новый маршутизатор

  fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
  mux.Handle("/static", http.NotFoundHandler())
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}

// Open() вызывается каждый раз, когда http.FileServer получает запрос
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
  f, err := nfs.fs.Open(path) // http.Open() inmplements http.FileSystem interface
  if err != nil {             // using os.Open()
    return nil, err
  }

  s, err := f.Stat()
  if s.IsDir() {          // Если это папка, то с пом. Stat() проверим наличие index.html внутри
    index := filepath.Join(path, "index.html")
    if _, err := nfs.fs.Open(index); err != nil {
      closeErr := f.Close() // Закрываем чтобы избежать утечки файлового дескриптора
      if closeErr != nil {
        return nil, closeErr
      }
      return nil, err
    }
  }
  return f, nil
}
