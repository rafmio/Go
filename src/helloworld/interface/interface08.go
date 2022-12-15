// https://golangify.com/interface
package main

import (
	"fmt"
	"strings"
)

type martian struct {
	Skin string
}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

func main() {
	mar := martian{Skin: "Green"}
	var laserAccuracy laser = 10

	fmt.Println(mar.talk())
	fmt.Println(laserAccuracy.talk())
}
