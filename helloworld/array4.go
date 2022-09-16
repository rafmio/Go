package main
import "fmt"

func myfun(a [6]int, size int) int {
    var k, val, r int
    
    for k = 0; k < size; k++ {
        val += a[k]
    }
    
    r = val / size
    return r
}

func main() {
    var arr1 = [6]int{34, 65, 2032, 103, 11, 53}
    var res int
    res = myfun(arr1, len(arr1))
    fmt.Printf("Final result is: %d\n", res)
}

// Passing an array to function
// https://www.geeksforgeeks.org/how-to-pass-an-array-to-a-function-in-golang/
