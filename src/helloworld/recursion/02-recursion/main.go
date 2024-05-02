package main

import (
	"fmt"
)

func f(v int) {
	if v == 0 {
		fmt.Println("Zero")
		return
	} else {
		fmt.Printf("%d ", v)
		f(v - 1)
	}
}

func main() {
	f(5)
}

// базовый случай (base case) - условие выхода из рекурсии
// рекурсивное отношение - вызов самой рекурсии
