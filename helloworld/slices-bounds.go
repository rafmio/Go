package main
import "fmt"

func main() {
    s := []int{10, 30, 50, 70, 110, 130}
    
    s = s[1:4]
    fmt.Println(s)
    
    s = s[:2]
    fmt.Println(s)
    
    s = s[1:]
    fmt.Println(s)
    
}
