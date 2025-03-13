package main

import "fmt"

func odd(channel chan int) {
  channel <- 1
  channel <- 3
  channel <- 5
}

func even(channel chan int) {
  channel <- 2
  channel <- 4
  channel <- 6
}

func main() {
  channelA := make(chan int)
  channelB := make(chan int)
  go odd(channelA)
  go even(channelB)
  fmt.Println(<- channelA)
  fmt.Println(<- channelB)
  fmt.Println(<- channelA)
  fmt.Println(<- channelB)
  fmt.Println(<- channelA)
  fmt.Println(<- channelB)
}
