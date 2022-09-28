package main
import "fmt"

func main() {
    str1 := "GeeksforGeeks"
    str2 := "GeeksForgeeks"
    str3 := "GeeksforGeeks"
    result1 := str1 == str2
    result2 := str1 == str3
    
    fmt.Println( result1)
    fmt.Println( result2)
    
    fmt.Println("The type of result1 is %T and " +
    "the type of result2 is %T", result1, result2)
}
