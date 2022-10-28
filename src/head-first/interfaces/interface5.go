package main

import "fmt"

type Robot string
func (r Robot) MakeSound() {
  fmt.Println("Beep Boop")
}
func(r Robot) Walk() {
  fmt.Println("Powering ligs")
}

type NoiseMaker interface {
  MakeSound()
}

func main() {
  var noiseMaker NoiseMaker = Robot("Botco Ambler")
  noiseMaker.MakeSound()
  var robot Robot = noiseMaker.(Robot)
  fmt.Println("Print robot var", robot)
  robot.Walk()
  fmt.Printf("Type of robot: %T\n", robot)

}
