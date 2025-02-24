package main

import (
  "time"
)

func writeToChannel(nameChannel chan <- string) {
  names := []string{"Alice", "Bob", "Charlie", "Dora", "Valera", "Kuzma"}

  tickChannel := time.Tick(time.Second)
  index := 0

  for {
    <- tickChannel
    nameChannel <- names[index]
    index++
    if (index == len(names)) {
      index = 0
    }
  }
}

func main() {
  nameChannel := make(chan string)

  go writeToChannel(nameChannel)

  for name := range nameChannel {
    Printfln("Read name: %v", name)
  }
}
