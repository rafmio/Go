package main

import (
	"fmt"
	"subdir01"
	"subdir02"
)

func main() {
	fmt.Println("main(): hello")
	subdir01.Sd01Func()
	subdir02.Sd02Func(subdir02.StrToPassToFunc)
}
