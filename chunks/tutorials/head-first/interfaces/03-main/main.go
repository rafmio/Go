package main

import "gadget"

type Player interface {
  Play(string)
  Stop()
}

func playList(device Player, songs []string) {
  for _, song := range songs {
    device.Play(song)
  }
  device.Stop()
}

func main() {
  mixtape := []string{"Enter Sandman", "One", "Nothing Else Metters"}
  var player Player = gadget.TapePlayer{}
  playList(player, mixtape)
  player = gadget.TapeRecorder{}
  playList(player, mixtape)
}
