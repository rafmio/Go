package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func intSeq() func() int {
	i := 0

	exploreVar(&i)

	return func() int {
		i++
		if i == 3 {
			return i * 100
		}
		return i
	}
}

func main() {
	readAndSaveFile()

	nextInt := intSeq()
	exploreFuncVar(&nextInt)

	fmt.Println("first call: ", nextInt())
	exploreFuncVar(&nextInt)

	fmt.Println("second call: ", nextInt())
	exploreFuncVar(&nextInt)

	fmt.Println("third call: ", nextInt())
	exploreFuncVar(&nextInt)

	newInts := intSeq()
	exploreFuncVar(&newInts)
	fmt.Println("fourth call:", newInts())
}

func exploreVar(ptr *int) {
	fmt.Printf("variable's : %p\n", &ptr)
}

func exploreFuncVar(f *func() int) {
	fmt.Printf("func's address: %p\n", f)
}

func readAndSaveFile() {
	pid := os.Getpid()
	path := fmt.Sprintf("/proc/%d/maps", pid)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file: %v", err)
	}

	err = os.WriteFile("15-maps.txt", []byte(content), 0644)
	if err != nil {
		log.Fatal("Error writing file: %v", err)
	} else {
		fmt.Println("The file has been wrote")
	}

}
