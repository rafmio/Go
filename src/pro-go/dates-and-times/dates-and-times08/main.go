package main

import (
  "fmt"
  "time"
)

func writeToChannel(channel chan <- string) {
  time := time.NewTimer(time.Minute * 10)

  go func() {
    time.Sleep(time.Second * 2)
    Printfln("Resetting timer")
    timer.Reset(time.Second)
  }()

  Printfln("Waiting for initial duration...")
  <- timer.C
  Printfln("Initial duration elapsed")

  names := []string {"Alice", "Bob", "Charlie", "Dora"}
  for _, name := range names {
    channel <- name
  }
  close(channel)
}
