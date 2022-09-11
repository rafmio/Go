package main
import "fmt"

func main() {
    primes := []int{2, 44, 55, 66, 77, 88, 190}
    var sls []int = primes[1:4]
    fmt.Println(sls)
    
    var numbers []int = []int{1, 2, 3, 4, 5, 100}
    var numsls[]int = numbers[3:6]
    fmt.Println(numbers)
    fmt.Println(numsls)
    
}

// An array has a fixed size. A slice, on the other hand, is 
// dynamically-sized, flexible view into the elements of an array
