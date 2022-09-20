package main
import  (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Printf("time.Now(): %s\n", t)
    
    switch {
        case t.Hour() < 12:
            fmt.Println("Good morning!")
        case t.Hour() < 17:
            fmt.Println("Good afternoon")
        default:
            fmt.Println("Good evening")
    }
    
    
    fmt.Printf("h.Hour(): %d\n", t.Hour())
}

// Switch with no condition
// https://go.dev/tour/flowcontrol/11
