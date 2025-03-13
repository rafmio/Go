// Doesn't work!
package main

// #include <stdio.h>
// void callC() {
// 	printf("Hello from C!\n");
// }

import "C"
import "fmt"

func main() {
	println("A Go statement!")
	C.callC()
	println("Another Go statement!")
	fmt.Println("Finish")
}
