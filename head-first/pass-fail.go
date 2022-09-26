// pass-fail.go
// notifies whether the user has passed the exam

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    fmt.Print("Enter a grade: ")
    reader := bufio.NewReader(os.Stdin)
    input, err  := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(input)
}
