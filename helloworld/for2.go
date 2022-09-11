package main
import "fmt"

func main() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    
    fmt.Println(sum)
    
    summ  := 1
    for ; summ < 1000; {
        summ += summ
    }
    fmt.Println(summ)
    
}

// Init and post statements are optional
