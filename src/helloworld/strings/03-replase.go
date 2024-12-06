package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Supermassive Black Hole"
	output := strings.Replace(input, " ", "", -1)
	fmt.Println(output) // Output: SupermassiveBlackHole

	input = "Muse"
	output = strings.Replace(input, " ", "", -1)
	fmt.Println(output) // Output: Muse

	input = "The Dark    Side of the Moon"
	output = strings.Replace(input, " ", "", -1)
	fmt.Println(output) // Output: TheDarkSideoftheMoon
}
