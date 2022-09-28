package main
import "fmt"


func main() {
    // Example 1 -----------------------------------------
    mychnl := make(chan string)
    
    go func() {
        mychnl <- "Kissy-Missy"
        mychnl <- "Huggy-Wuggy"
        mychnl <- "Tosi-Bosi"
        mychnl <- "Vinne-Pooh"
    
        close(mychnl)
    } ()
    
    for res := range mychnl {
        fmt.Println(res)
    }
    
    fmt.Println("----------------------------------------")
    // Example 2 -----------------------------------------
    mychan2 := make(chan string, 4)
    mychan2 <- "My fried of mystery"
    mychan2 <- "Sad but true"
    mychan2 <- "Nothing else metter"
    mychan2 <- "Master"
    
    fmt.Println("Length of the channel is :", len(mychan2))
    
    fmt.Println("----------------------------------------")
    // Example 3 -----------------------------------------
    mychan3 := make(chan string, 5)
    mychan3 <- "My fried of mystery"
    mychan3 <- "Sad but true"
    mychan3 <- "Nothing else metter"
    mychan3 <- "Master"
    
    fmt.Println("Length of the channel is: ", len(mychan3))
    fmt.Println("Capacity of the channel is: ", cap(mychan3))
}


// Channel 
// for loop in the channel
// https://www.geeksforgeeks.org/channel-in-golang/
