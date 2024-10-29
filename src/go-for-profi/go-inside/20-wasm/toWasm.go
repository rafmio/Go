package main

import (
	"fmt"
)

func main() {
	fmt.Println("Creating WebAssembly code from Go!")
}
// $ GOOS=js GOARCH=wasm go build -o main.wasm toWasm.go
