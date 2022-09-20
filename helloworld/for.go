package main
import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }

    fmt.Println(sum)

    var variable1 int = 5
    for i := 0; i < 7; i++ {
          variable1 = i * variable1 + variable1
    }

    fmt.Println("variable1 = ", variable1)
}



// For
// https://go.dev/tour/flowcontrol/1
