// Blank Identifier (underscore)
// https://www.geeksforgeeks.org/what-is-blank-identifierunderscore-in-golang/

package main
import "fmt"

func main() {
    // When we no need second returned value
    mul, _ := mul_div(105, 7)
    fmt.Println("105 x 7 = ", mul)
}

func mul_div(n1 int, n2 int) (int, int) {
    return n1 * n2, n1 / n2
}
