package main
import "fmt"

func myfun(mychnl chan string) {
    for v := 0; v < 4; v++ {
        mychnl <- "GFG"
    }
    close(mychnl)
}

func main() {
    c := make(chan string)
    
    go myfun(c)
    
    for {
        res, ok := <- c
        if ok == false {
            fmt.Println("Channel is Closed ", ok)
            break
        }
        fmt.Println("Channel is Open ", res, ok)
    }
}

// Channel 
// Closing a channel
// https://www.geeksforgeeks.org/channel-in-golang/
