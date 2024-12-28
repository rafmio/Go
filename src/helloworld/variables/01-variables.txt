Variables
// https://www.geeksforgeeks.org/go-variables/

Declaring:
- Using var keyword
- Using short variable declaration

Using var keyword:
var variable_name type = expression

var myVariable = 20
var myVariable int = 20

var myVariable = "Cogito ergo sum"
var myVariable string = "Cogito ergo sum"

var myVariable = 64.77
var myVariable float64 = 64.77

In the above syntax, either type or =expression can be omitted, but not both. If the =expression is omitted, then the variable value is determined by its type's default value. The default value is usually 0.

There is no such concept of an unitialized variable in Go
var 

If you use type, then you are allowed to declare multiple variables of the same type in the single declaration:

var myVar1, myVar2, myVar3 int = 2, 464, 57

If you remove type, then you are allowed to declare multiple variables of a different type on the single declaration:

var myVar1, myVar2, myVar3 = 2, "GFG", 67.56

Using short variable declaration:
variable_name := expression
:= - is a declaration
 = - is a assignment 
 
 Using short variable declaration you are allowed to declare multiple variables in the single declaration:
 
 myVar1, myVar2, myVar3 := 800, 435, 56
 
 In a short variable declaration, you are allowed to initialize a set of variables by the calling function that returns multiple values:
 i, j := os.Open(name) // os.Open function return a file in i and
                       // error in j variable
 A short variable declaration acts only when for those variables that already declared in the same lexical block
