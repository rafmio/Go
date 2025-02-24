package main

import (
	"fmt"
	"os"
	"time"
)

func Aname() {
	arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}

	for t1 := 0; t1 <= len(arr1)-1; t1++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%d - %s\n", os.Getpid(), arr1[t1])
	}
}

func Aid() {
	arr2 := [4]int{300, 400, 500, 600}

	for t2 := 0; t2 <= len(arr2)-1; t2++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d - %d\n", os.Getpid(), arr2[t2])
	}
}

func main() {
	fmt.Println("!... Main Go-routines Start...!")

	go Aname()

	go Aid()

	time.Sleep(4500 * time.Millisecond)
	fmt.Println("!...Main Go-routine end ...!")
}
