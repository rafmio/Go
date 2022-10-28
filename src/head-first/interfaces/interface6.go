package main

import "fmt"

type Truck string
func (t Truck) Accelerate() {
  fmt.Println("Speeding up")
}
func (t Truck) Brake() {
  fmt.Println("Stopping")
}
func (t Truck) Steer(direction string) {
  fmt.Println("Turning", direction)
}
func (t Truck) LoadCargo(cargo string) {
  fmt.Println("Loading", cargo)
}

type Vehicle interface {
  Accelerate()
  Brake()
  Steer(string)
}

func TryVehicle(vehicle Vehicle) {
  vehicle.Accelerate()
  vehicle.Steer("left")
  vehicle.Steer("right")
  vehicle.Brake()
  truck, ok := vehicle.(Truck)
  if ok {
    truck.LoadCargo("test cargo")
  }

}

func main() {
  TryVehicle(Truck("Fnord F180"))

  fmt.Println()
  var tstVar Truck = "Suziki"
  fmt.Println(tstVar)

  fmt.Println()
  var tstVar2 Vehicle = Truck("MAN")
  tstVar2.Accelerate()
  tstVar2.Brake()

  fmt.Println()
  tstVar3 := Truck("DAF")
  tstVar3.Accelerate()
  fmt.Printf("Type of tstVar3: %T\n", tstVar3)

  fmt.Println()
  TryVehicle(tstVar3)
}
