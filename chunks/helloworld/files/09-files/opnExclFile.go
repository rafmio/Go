package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Try to open '/home/raf/Go/src/helloworld/files/08-files/iinput.txt'...")
	file, err := os.Open("/home/raf/Go/src/helloworld/files/08-files/iinput.txt")
	if err != nil {
		fmt.Println("opening file:", err.Error())
		return
	}
	defer file.Close()

	time.Sleep(time.Second * 15)
	data := []string{"str10", "str11", "str12"}

	for _, line := range data {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("writing to file:", err.Error())
			return
		}
	}

	fmt.Println("Try to read from file...")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("UFW: ", line)
	}

}
