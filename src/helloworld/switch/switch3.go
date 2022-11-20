package main
import "fmt"

func main() {
    var value interface{}
        switch q := value.(type) {
            case bool: 
                fmt.Println("Value is of boolean type")
            case float64: 
                fmt.Println("Value is of float64 type")
            case int:
                fmt.Println("Value is of int type")
            default:
                fmt.Println("Value is of type: %T", q)
        }
}
