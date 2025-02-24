package main
import "fmt"

func main() {
    // Example 1 ---------------------------------------
    odd := [7]int{1, 3, 5, 7, 9, 11, 13}
    
    for i, item := range odd {
        fmt.Printf("odd[%d] = %d\n", i, item)
    }
    
    fmt.Println("--------------------------------------")
    // Example 2 ---------------------------------------
    var strng = "Hardwired to self destruct"
    
    for i, item := range strng {
        fmt.Printf("string[%d] = %d\n", i, item)
    }
    
    fmt.Printf("Type of strng is: %T\n", strng)

    fmt.Println("--------------------------------------")
    // Example 3 ---------------------------------------
    student_rank_map := map[string]int{ "Joshua": 3, "Nishka" : 2, "Zorro" :1 }
    
    for student := range student_rank_map {
        fmt.Println("Rank of", student, "is:\t", student_rank_map[student])
    }
}

// range
// https://www.geeksforgeeks.org/range-keyword-in-golang/
