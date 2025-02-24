package subdir03

import (
	"fmt"
	"ltr/subdir02"
)

func PrintExportedInstance(inst subdir02.ExportedSubdir02Struct) {
	fmt.Printf("name: %s, city: %s, price: %.2f, age: %d\n", inst.Name, inst.City, inst.Price, inst.Age)
}
