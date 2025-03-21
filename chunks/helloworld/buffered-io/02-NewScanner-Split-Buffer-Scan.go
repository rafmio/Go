// https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := "abcdefghijklmnopqrstuvw"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)
		return 0, nil, nil
	}
	scanner.Split(split)
	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}
