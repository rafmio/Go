package main

import (
	"fmt"
)

const (
	read = 1 << iota
	write
	exec

	admin = read | write | exec
)

func main() {
	fmt.Printf("read = %v\nwrite = %v\nexec = %v\nadmin = %v\n",
		read, write, exec, admin,
	)
}
