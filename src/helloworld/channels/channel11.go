// https://habr.com/ru/post/490336/
package main

import "fmt"

func squares(c chan int) {
  for i := 0; i <= 9; i++ {
    c <- i * i
  }

  close(c)
}

func main() {
  fmt.Println("main() started")
  c := make(chan int)

  go squares(c)

  for {
    val, ok := <- c
    if ok == false {
      fmt.Println(val, ok, "<--loop is broken!")
      break
    } else {
      fmt.Println(val, ok)
    }
  }

  fmt.Println("main() stopped")
}
