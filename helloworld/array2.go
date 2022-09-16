package main
import "fmt"

func main() {
    // Example 1 ----------------------------------------------
    arr := [4]string{"geek", "gfg", "Geek1234", "GeeksForGeeks"}
    fmt.Println("Elements of the array: ") 
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
    // Example 2 ----------------------------------------------
    arr2 := [3][3]string{ {"C#", "C", "Python"}, 
                          {"Java", "Scala", "Perl"} }

    fmt.Println("Elements of 2-D array:")
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            fmt.Println(arr2[x][y])
        }
    }
    // Example 3 ----------------------------------------------
    var arr3 [2][2]int
    arr3[0][0] = 100
    arr3[0][1] = 200
    arr3[1][0] = 300
    arr3[1][1] = 400
    
    fmt.Println("Elements of the 2-D array: ")
    for p := 0; p < 2; p++ { 
        for q := 0; q < 2; q++ {
            fmt.Println(arr3[p][q])
        }
    }
    // Example 4 ----------------------------------------------
    var arr4[2] int
    fmt.Println("Elements of the array: ", arr4)
    // Example 5 ----------------------------------------------
    arr5 := [3]int{9, 7, 6}
    arr6 := [...]int{9, 10, 323, 441, 4145, 10}
    arr7 := [3]int{120, 230, 440}
    
    fmt.Println("Length of the arr5 is: ", len(arr5))
    fmt.Println("Length of the arr6 is: ", len(arr6))
    fmt.Println("Length of the arr7 is: ", len(arr7))
    // Example 6 ----------------------------------------------
    arr8 := [...]string{"GFG", "gfg", "geeks", "GeekForGeeks", "GfG"}
    fmt.Println("The length of the arr8 is: ", len(arr8))
    // Example 7 ----------------------------------------------
    arr9 := [...]int{29, 43, 54, 54,
                     20, 51, 94, 11}
    for x := 0; x < len(arr9); x++ {
        fmt.Printf("%d\n", arr9[x])
    }
    // Example 8 ----------------------------------------------
    arr10 := [...]int{100, 200, 300, 400, 500, 900}
    fmt.Println("--------------------")
    fmt.Println("Original array (Before): ", arr10)
    
    arr11 := arr10
    fmt.Println("New array (Before):      ", arr11) 
    
    arr11[0] = 1000
    arr11[3] = 0 
    
    fmt.Println("New array (after):       ", arr11)
    fmt.Println("Original array (after):  ", arr10)
    
    fmt.Println()
    fmt.Println("--------------------")
    // Example 9 ----------------------------------------------
    arr12 := [3]int{9, 10, 12}
    arr13 := [...]int{9, 10, 12}
    arr14 := [3]int{9, 14, 16}
    
    fmt.Println(arr12 == arr13)
    fmt.Println(arr13 == arr14)
    fmt.Println(arr14 == arr12)

    fmt.Println()
    fmt.Println("--------------------")
    // Example 10 ----------------------------------------------
}
