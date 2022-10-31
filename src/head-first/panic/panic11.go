package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "os"
)

func scanDirectory(path string) {
    fmt.Println("path variable: ", path)
    files, err := ioutil.ReadDir(path)
    if err != nil {
        panic(err)
    }
    
    for _, file := range files {
        filePath := filepath.Join(path, file.Name())
        if file.IsDir() {
            scanDirectory(filePath)
        } else {
            fmt.Println(filePath)
        }
    }
}

func main() {
    flpth := os.Args[1]
    scanDirectory(flpth)
}
