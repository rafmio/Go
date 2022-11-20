package main
import "fmt"

type Student struct {
    name    string
    branch  string
    year    int
}

type Teacher struct {
    name    string
    subject string
    exp     int
    details Student
}

func main() {
    result := Teacher{
        name:       "Suman",
        subject:    "Java",
        exp:        5,
        details:    Student{"Bongo", "CSE", 2}, 
    }
    
    fmt.Println("Details of the Teacher")
    fmt.Println("Teacher's name: ", result.name)
    fmt.Println("Subject:        ", result.subject)
    fmt.Println("Experience:     ", result.exp)
    
    fmt.Println()
    
    fmt.Println("Details of Student")
    fmt.Println("Name:   ", result.details.name)
    fmt.Println("Branch: ", result.details.branch)
    fmt.Println("Year:   ", result.details.year) 
}

// Nested structures
// https://www.geeksforgeeks.org/nested-structure-in-golang/
