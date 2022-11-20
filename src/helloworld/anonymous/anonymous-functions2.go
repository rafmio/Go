package main
import "fmt"

func GFG() func(i, j string) string {
    myf := func(i, j string) string {
        return i + j + "Paul McCartney"
    }
    return myf
}

func main() {
    value := GFG()
    fmt.Println(value ("Welcome ", "to "))
}
