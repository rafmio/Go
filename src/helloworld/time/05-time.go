// see 'go doc time.Layout'
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	tfmt := t.Format("2006-01-02T15:04")
	fmt.Println(tfmt)
}
