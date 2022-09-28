Constants
// https://www.geeksforgeeks.org/constants-go-language/
const PI = 3.14
const PI float32 = 3.14

List of constants:
- numeric constants (integer constants, floating constants, complex constants)
- string literals
- boolean constants

Integer constant:
An integer literal can also have a suffix that is a combination of U and L, for unsigned and long, respectively

String literals
Go supports two types of string literals: 
" " - double-quote style
' ' - back-quote
type _string struct { 
    element *byte   // underlying bytes
    len      int    // number of bytes
}
