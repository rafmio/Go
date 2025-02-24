// https://прохоренок.рф/pdf/go/ch14-go-struktura-scanner-postrochnoe-chtenie-iz-fayla.html
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter text:")
	scanner := bufio.NewScanner(os.Stdin) // os.Stdin implements io.Reader?
	if scanner.Scan() {                   // bufio.Scan делит входной поток на токены
		txt := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println("You have entered:", txt)
	} else {
		fmt.Fprintln(os.Stderr, "Error occurred")
	}

	buf := strings.NewReader("Строка1\r\nСтрока2\nСтркока3")
	scanner = bufio.NewScanner(buf)
	for scanner.Scan() {
		fmt.Printf("%q\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
