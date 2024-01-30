package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type MyStruct struct {
	aaa, bbb, ccc int
}

func (sss MyStruct) calc() int {
	return sss.aaa + sss.bbb + sss.ccc
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	fmt.Println("math.Sqrt2:", math.Sqrt2)

	sss := MyStruct{100, 200, 300}
	fmt.Printf("%d + %d + %d = %d\n", sss.aaa, sss.bbb, sss.ccc, sss.calc())
}

// Methods continued
// You can declare a method on non-struct types, too
// You can only declare a method with a receiver whose type is
// difined in the same package as the method. You cannot declare
// a method with a receiver whose type is defined in another package
// (which includes the build-in types such as int)

// https://go.dev/tour/methods/3
