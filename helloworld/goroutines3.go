package main
import (
    "fmt"
    "time"
)

func display(str string) {
    for w := 0; w < 6; w++ {
        time.Sleep(1 * time.Second)
        fmt.Println("String inside for loop: ", str)
    }
}

func main() {
    go display("Hello inside Goroutine")
    
    display("Hola inside normal fucntion") 
}

// https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/
