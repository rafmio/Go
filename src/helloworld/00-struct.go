Structures
https://www.geeksforgeeks.org/structures-in-golang/

Syntaxx for declaring a structure:

var a Address

The above code creates a variable of a type Address which is by default set to zero. For struct, zero means all the fields are set to their corresponding zero value. 

Accessing by pointers
    emp8 := &Employee{"Sam", "Anderson", 55, 6000}
    
    fmt.Println("First Name: ", (*emp8).firstName)
    fmt.Println("Age: ", (*emp8).age)
    
    fmt.Println("First Name: ", emp8.firstName)
    fmt.Println("Age: ", emp8.age)
    
