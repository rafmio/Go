package main

import (
	"fmt"
	"ltr/subdir01"
	"ltr/subdir02"
)

func main() {
	fmt.Println("main(): hello")
	subdir01.Sd01Func()
	subdir02.Sd02Func(subdir02.StrToPassToFunc)
}
