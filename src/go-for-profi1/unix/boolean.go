// https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go-ru
package main

import (
	"flag"
	"fmt"
	"os"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	useColor := flag.Bool("color", false, "display colorized output")
	flag.Parse()

	if *useColor {
		colorize(ColorBlue, "Hello, ColorBlue!")
	} else {
		colorize(ColorYellow, "Yeah, it's Yellow!")
	}

	fmt.Fprintf(os.Stdout, "Hello-mello, Tosy-Bosy, Kissy-Missy\n")
}
