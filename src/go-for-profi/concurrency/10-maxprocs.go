// To programmatically find out the value of the GOMAXPROCS environment variable, you need
// to use the following code
package main

import (
	"fmt"
	"runtime"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", getGOMAXPROCS())

}
