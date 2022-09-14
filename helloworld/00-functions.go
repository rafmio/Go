// Functions
// By default Go use call by value method to pass arguments in function.
// Go supports 2 ways to pass arguments to function:
// - Call by value
// - Call by reference

// Variadic Functions
// The function that is called with the varyin number of arguments is known as Variadic function. User is allowed to pass zero or more arguments in the Variadic function. 
// 
// In the declaration of the variadic function, the type of the last parametr is preceded by an ellipsis (...). It indicates that the function can be called at any number of parameters of this type.
// 
// Inside the function ...type behaves like a slice. 
// You can also pass an existing slice in a variadic function
// 
// The variadic function are generally used for functions that perform string formatting
// 
// You can pass multiple slices in the variadic function
// 
// Anonymous functions
// An anonymous function is a function which doesn't contain any name.
// It is useful when you want to create an inline function. 
// An anonymous functions is also known as function literal. 
func () {
    fmt.Println("Hello, buddy, how it's going?")
} ()

// You are allowed to assign an anonymous function to a variable. When you assign a function to a variable, then the type of the variableis of function type and you can call variable like a function.

value := func() {
    fmt.Println("Wellcome to anonymous function")
}

value()

// (without parenthesis)

// We can also pass arguments in the anonymous function:
func (ele string) {
    fmt.Println(ele)
} ("Joe Satriani")

// Pass an anonymous function as an argument into other function
func GFG(i func(p, q string) string {
    fmt.Println(i ("Mikle", "Gordan")
}

func main() {
    value := func(p, q string) string {
        return p + q + "Chicago Bulls"
    }
    GFG(value)
}

// https://www.geeksforgeeks.org/functions-in-go-language/
