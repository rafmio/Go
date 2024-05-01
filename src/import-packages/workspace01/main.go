package main

import (
	"workspace01/pack1"

	"workspace01/pack1/subpack1"
)

func main() {
	pack1.SayHello()
	subpack1.SayCustomHello()
	pack2.sayGoodBye()
}
