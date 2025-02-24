// ERROR! Wrong syntax

package main

import "fmt"

func functions() []func() {
    arr := []int{1, 2, 3, 4}
    result := make([]func(), 0)
    
    for i := range arr {
        result = append(result, func() {
            fmt.Printf("index - %d, value - %d\n", i, arr[i])
        })
    }
    
    for i := range arr {
        result = append(result, func(index, val int) { fmt.Printf("index - %d, value - %d\n", index, val)
        })
    }
    return result
}

func main() {
    fns := functions()
    for f := range fns {
        fns[f]()
    }
}

// https://kovardin.ru/articles/go/zamykaniya/
