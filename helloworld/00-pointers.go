Pointers
'*' - dereferencing operator used to declare pointer variable and access the value stored in the address
'&' - address operator used to returns the address of a variable or to  access the address of a variable to a pointer

Declaring pointer:
var pointer_name *DataType
var pointer_to_int *int
var pointer_to_flt *float32
var s *string

The default value or zero-value of a pointer is always nil. 
Declaration and initializaion of the pointers can be done into a single line:
var s *int = &a

You can pass the address of the struct to the pointer which represents the pointer to struct concept. There is no need to use dereferencing explicitly as it will give the same result.

type Employee struct {
    name    string
    empid   int
}

func main() {
    emp := Employee{"Justin Jonhson", 19078}
    pts := &emp 
    
    fmt.Println(pts)
    fmt.Println(pts.name)
    fmt.Println((*pts).name)
    fmt.Println(pts.empid)
    fmt.Println((*pts).empid)
}
// ---------------------------------------------
Double Pointers - pointers to pointers

// Will give the same result
// &{Justin Jonhson 19078}
// Justin Jonhson
// Justin Jonhson
// 19078
// 19078

// https://www.geeksforgeeks.org/pointers-in-golang/
