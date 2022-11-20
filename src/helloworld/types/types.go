package main
import "fmt"

func main() {
    textHello := "Hello"
    textBe    := "BeGeek"
    
    i := 10
    text := textHello + " " + fmt.Sprint(i)
    fmt.Println(text)
    fmt.Println(fmt.Sprintf("%s %d %s", textHello, i, textBe))
    
    var h float64
    h = 10.3
    next := float64(i) + h
    fmt.Println(next)
    
    var input float64
    fmt.Scanf("%f", &input)
    output := input + h
    fmt.Println(output)
    
    const cnstvar string = "const string"
    fmt.Println(cnstvar)
}

// Types
// User input
