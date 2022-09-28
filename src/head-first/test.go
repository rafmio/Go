package main

import (
	"bufio"
	"os"
	"fmt"
	// "fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Println(input, err, reader)
}
