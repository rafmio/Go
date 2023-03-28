package main

import (
	"fmt"
)

func main() {
	type XYZ struct {
		X int
		Y int
		Z int
	}
	var s1 XYZ
	fmt.Println(s1.X, s1.Y, s1.Z)

	p1 := XYZ{23, 12, -2}
	p2 := XYZ{Z: 12, Y: 13}
	fmt.Println(p1)
	fmt.Println(p2)

	pSlise := [4]XYZ{}
	pSlise[2] = p1
	pSlise[0] = p2
	fmt.Println(pSlise)
	p2 = XYZ{1, 2, 3}
	fmt.Println(pSlise)

}
