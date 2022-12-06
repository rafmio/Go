package main

import (
  "time"
  "fmt"
)

func writeToChannel(channel chan <- string) {
  names := []string{"Alice", "Bob", "Charlie", "Dora", "Jessie"}

  for _, name := range names {
    channel <- name
    time.Sleep(time.Second * 1)
  }
  close(channel)
}

func writeToChannelDefer(channel chan<- string) {
  names := []string{"Kissy-Missy", "Huggy-Wuggy", "Tosy-Bosy"}
  for _, name := range names {
    channel <- name
  }
  close(channel)
}

func main() {
  nameChannel := make(chan string)

  go writeToChannel(nameChannel)

  for name := range nameChannel {
    Printfln("Read name: %v", name)
  }

  fmt.Println("Stopped-------------------------------------")
  fmt.Println()

  nameChannel2 := make(chan string)

  time.AfterFunc(time.Second * 3, func() {
    writeToChannelDefer(nameChannel2)
  })
  for name2 := range nameChannel2 {
    Printfln("Read name: %v", name2)
  }


}
