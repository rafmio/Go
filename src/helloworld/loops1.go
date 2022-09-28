package main
import "fmt"

func main() {
    chnl := make(chan int)
    go func() {
        chnl <- 100
        chnl <- 1000
        chnl <- 10000
        chnl <- 100000
        close(chnl)
    } ()
    for i := range chnl {
        fmt.Println(i)
    }
}
