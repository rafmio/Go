package main
import "fmt" 

func main() {
    varvar := make([]int, 10)
    varvar[0] = 1 << uint(0)
    varvar[1] = 1 << uint(1)
    fmt.Println(varvar)
    
    varvar[4] = 1 << uint(4)
    fmt.Println(varvar)
}
