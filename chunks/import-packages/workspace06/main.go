package main

import (
	"fmt"
	"ltr/subdir01"
	"ltr/subdir02"
	"ltr/subdir03"
)

func main() {
	fmt.Println("hello from main()")

	inst := newInrootdirStruct("Bobby", "Tokio", 55.45)
	fmt.Printf("name: %s, city: %s, price: %.2f\n", inst.name, inst.city, inst.price)
	fmt.Printf("type of 'inst': %T\n", inst)
	sayHelloFromRootDir()

	subdir01.PrintExportedHello()

	inst2 := new(subdir02.ExportedSubdir02Struct)
	inst2.Name = "Holly"
	inst2.City = "Berlin"
	inst2.Price = 100.55
	inst2.Age = 34

	fmt.Println(inst2)
	subdir03.PrintExportedInstance(*inst2)
}
