package main

import "fmt"

type Driver struct {
  Name string
  Age int
  Car string
}

func (d *Driver) printName() {
  fmt.Println("Driver's name is: ", d.Name)
}

func (d *Driver) getAge() int {
  return d.Age
}

func (d *Driver) getCar() string {
  return d.Car
}

func main() {
  var driver1 *Driver = &Driver{"Kukushkin", 37, "DAF"}
  driver2 := &Driver{
    "Vasilyev",
    35,
    "MAN",
  }

  driver1.printName()
  fmt.Println("Dirver's age is: ", driver1.getAge())
  fmt.Println("Car:", driver1.getCar())

  driver2.printName()
  fmt.Println("Driver's age is: ", driver2.getAge())
  fmt.Println("Car: ", driver2.getCar())
}
