package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
  fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
  fmt.Println("Honk!")
}

type NoiseMaker interface {
  MakeSound()
}

func play(n NoiseMaker) {
  n.MakeSound()
}

func main() {
  var toy NoiseMaker

  toy = Whistle("Tyoco Canary")
  fmt.Printf("toy value: %s, toy type: %T\n", toy, toy)
  toy.MakeSound()
  play(Whistle("Toyco Canary"))

  toy = Horn("Toyco Blaster")
  fmt.Println()
  fmt.Printf("toy value: %s, toy type: %T\n", toy, toy)
  toy.MakeSound()
  play(Horn("Toyco Blaster"))
}
