package main
import "fmt"

func main() {
    sss := []int{2, 4, 8, 16, 32, 64, 128, 256, 512}
    fmt.Println("sss slice: ", sss)
    
    sss = sss[2:7]
    fmt.Println("slice: ", sss) 
    fmt.Println("sss[0]: ", sss[0])
    
    sss = sss[3:]
    fmt.Println("sss: ", sss)
}


// Slice defaults 
// https://go.dev/tour/moretypes/10
