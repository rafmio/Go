package main

import (
	"bufio"
	"fmt"
	"os"
)

type Playground struct {
	height int
	width  int
}

func (p Playground) Square() int {
	return p.height * p.width
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var width, height int
	fmt.Fscan(reader, &height)
	fmt.Fscan(reader, &width)

	rows := make([]string, 0)
	reader = bufio.NewReader(os.Stdin)
	for i := 0; i < height; i++ {
		var line string
		fmt.Fscan(reader, &line)
		rows = append(rows, line)
	}

	for _, value := range rows {
		fmt.Println(value)
	}
}
