package main
import "fmt"

const Pi = 3.14159

func main() {
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
}

// Constants
// https://go.dev/tour/basics/15
