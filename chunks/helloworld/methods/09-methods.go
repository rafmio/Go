package main

import (
	"fmt"
)

type Driver struct {
	Name string
	Age  int
	Car  string
}

func (d *Driver) printName() {
	fmt.Println("Driver's name is:", d.Name)
}

func (d *Driver) getAge() int {
	return d.Age
}

func (d *Driver) getCar() string {
	return d.Car
}

func main() {
	var driver1 *Driver = &Driver{"Kukushkin", 45, "DAF"}
	driver2 := &Driver{
		"Aliyakberov",
		40,
		"MAN",
	}

	driver1.printName()
	fmt.Println("The age of", driver1.Name, "is", driver1.getAge())
	fmt.Println("Car:", driver1.getCar())

	fmt.Println()

	driver2.printName()
	fmt.Println("The age of", driver2.Name, "is", driver2.getAge())
	fmt.Println("Car:", driver2.getCar())
}
