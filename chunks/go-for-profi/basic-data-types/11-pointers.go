package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	pid := os.Getpid()
	readMapsFile(pid)

	i := -10
	j := 25

	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	readMapsFile(pid)

	*pI = 123456
	*pI--
	fmt.Println("i:", i)

	readMapsFile(pid)

	getPointer(pJ)

	readMapsFile(pid)

	fmt.Println("j:", j)

	k := returnPointer(12)

	readMapsFile(pid)

	fmt.Println("*k:", *k)
	fmt.Println("k:", k)
}

func readMapsFile(pid int) {

	f, err := os.Open(fmt.Sprintf("/proc/%d/maps", pid))
	if err != nil {
		fmt.Println("Error opening maps file:", err)
		return
	}
	defer f.Close()

	var stackStr string
	var heapStr string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "stack") {
			// fmt.Println(line)
			stackStr = strings.Fields(line)[0]
		}
		if strings.Contains(line, "heap") {
			fmt.Println(line)
			heapStr = strings.Fields(line)[0]
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading maps file:", err)
		}
	}

	fmt.Println("[stack]:", stackStr)
	fmt.Println("[heap]:", heapStr)
}
