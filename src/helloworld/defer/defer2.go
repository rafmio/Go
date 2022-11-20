package main
import "fmt"

func add(a1, a2 int) int {
    res := a1 + a2 
    fmt.Println("Result: ", res)
    return 0
}

func main() {
    fmt.Println("Start")
    defer fmt.Println("End")
    defer add(34, 10)
    defer add(23, 100)
} 



// https://www.geeksforgeeks.org/defer-keyword-in-golang/
