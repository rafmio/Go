package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs of ")
    switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Println("OS X")
        case "linux":
            fmt.Println("Linux")
        default:
            // freeBSD, openBSD, Plan9, Windows...
            fmt.Printf("%s \n", os)
    }
}
