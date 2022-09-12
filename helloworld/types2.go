package main
import "fmt"

func main() {
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i = i + i
    }
    
    fmt.Println("\n")
    
    for i := 0; i <= 5; i++ {
        fmt.Println(i)
    }
    
    fmt.Println("\n")
    
    text := []string{"be", "geek", "hello"}
    fmt.Println(text[2])
    
    for _, i := range text {
        fmt.Println(i)
        if i == "geek" {
            continue
        }
    }
}
