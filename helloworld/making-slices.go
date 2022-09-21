package main
import "fmt"

func printSlice(s string, x []int) { 
    fmt.Printf("%s len = %d, cap = %d %v\n", 
               s, len(x), cap(x), x)
}

func main() {
    aaa := make([]int, 5)
    aaa[4] = 1000
    printSlice("aaa: ", aaa)
    
    bbb := make([]int, 1, 5)
    bbb[0] = 2000
    printSlice("bbb: ", bbb)
    
    ccc := bbb[:5]
    printSlice("ccc: ", ccc)
    
    ccc[4] = 1_500
    printSlice("ccc: ", ccc)
    
    ddd := ccc[2:5]
    printSlice("ddd: ", ddd)
   
}


// Creating a slice with make
// https://go.dev/tour/moretypes/13
