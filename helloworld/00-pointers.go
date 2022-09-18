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

// https://www.geeksforgeeks.org/pointers-in-golang/
