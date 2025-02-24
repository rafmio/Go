// https://habr.com/ru/post/490336/
package main

import "fmt"

func squares(cn chan int) {
	for i := 0; i <= 9; i++ {
		cn <- i * i
	}

	close(cn)
}

func main() {
	fmt.Println("main() started")
	cn := make(chan int)

	go squares(cn)

	for {
		val, ok := <-cn
		if ok == false {
			fmt.Println(val, ok, "<--loop is broken!")
			break
		} else {
			fmt.Println(val, ok)
		}
	}

	fmt.Println("main() stopped")
}
