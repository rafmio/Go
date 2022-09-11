package main
import "fmt"

func main() {
    var sum int16 = 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
    
    summ := 1
    for summ < 256 {
        summ += summ
    }
    fmt.Println(summ)
}
