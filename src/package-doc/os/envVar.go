package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("USER"))
	fmt.Println(os.Getenv("ARBITRARYENVVAR"))
	fmt.Println(os.Getenv("MUSIC_INFO"))
}
