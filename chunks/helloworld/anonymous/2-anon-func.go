// https://kovardin.ru/articles/go/zamykaniya/

package main

import (
	"fmt"
)

func printMessage(msg string) {
	fmt.Println(msg)
}

func getPrintMessage() func(string) {
	return func(message string) {
		fmt.Println(message)
	}
}

func main() {
	printMessage("Hello-mello function!")

	func(message string) {
		fmt.Println(message)
	}("Hello anonymous function!")

	printfunc := getPrintMessage()
	printfunc("Hello anonymous function using caller")
}
