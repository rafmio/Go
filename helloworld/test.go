package main
import "fmt"

func main() {
    strg := []string{"Huggy", "Waggy", "Kissy", "Missy"}
    fmt.Println(strg)
    
    strg = strg[0:2]
    fmt.Println(strg)
    
    strg = strg[2:4]
    fmt.Println(strg)
    
    fmt.Println(len(strg), cap(strg))
}
