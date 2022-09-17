package main
import "fmt"
import "sort"       // For Example 8

func main() {
    // Example 1----------------------------------------------
    myslice := []string{
        "This", "is", "the", "tutorial", "of", "Go", "language"}
    
    for e := 0; e < len(myslice); e++ {
        fmt.Printf("%s ", myslice[e])
    }
    
    fmt.Println("--------------------------------------------")
    // Example 2----------------------------------------------
    myslice2 := []string{"Kissy", "Missy", "Huggy", "Wuggy"}
    for index, ele := range myslice2 {
        fmt.Printf("Index = %d and element = %s\n", index + 3, ele)
    }
    
    fmt.Println("--------------------------------------------")
    // Example 3----------------------------------------------    
    myslice3 := []string{"Selectel", "Rostelecom", "Beeline", "MTC"}
    for _, ele := range myslice3 {
        fmt.Println("Element = ", ele)
    }
    
    
    fmt.Println("--------------------------------------------")
    // Example 4----------------------------------------------    
    var myslice4 []string
    fmt.Printf("len(myslice4): %d, cap(myslice4) = %d\n", 
               len(myslice4), cap(myslice4))

    
    fmt.Println("--------------------------------------------")
    // Example 5----------------------------------------------    
    arr1 := [6]int{555, 666, 777, 888, 999, 444}
    slc1 := arr1[0:4]
    
    fmt.Println("Original arr1: ", arr1)
    fmt.Println("Original slc1: ", slc1)
    
    slc1[0] = 100
    slc1[1] = 1000
    slc1[2] = 1000
    
    fmt.Println("Modified arr1: ", arr1)
    fmt.Println("Modified slc1: ", slc1)
    
    fmt.Println("--------------------------------------------")
    // Example 6----------------------------------------------    
    slc2 := []int{543, 654, 765}
    var slc3 [] int
    fmt.Println("slc2 == nil: ", slc2 == nil)
    fmt.Println("slc3 == nil: ", slc3 == nil)
    
    fmt.Println("--------------------------------------------")
    // Example 7----------------------------------------------    
    slc4 := [][]int{
        {12, 34}, 
        {56, 77},
        {29, 43},
        {43, 43},
    }
    
    fmt.Println("slc4: ", slc4)
    
    slc5 := [][]string{
        []string{"Geeks", "for"},
        []string{"Geeks", "GFG"},
        []string{"gfg", "geek"},
    }
    
    fmt.Println("slc5: ", slc5)
    
    fmt.Println("--------------------------------------------")
    // Example 8----------------------------------------------  
    slc6 := []string{"Python", "Java", "C#", "Go", "Ruby"}
    slc7 := []int{45, 65, 76, 34, 86, 4, 67, 34, 33, 98}
    
    fmt.Println("Before sorting: ")
    fmt.Println("slc6: ", slc6)
    fmt.Println("slc7: ", slc7)
    
    sort.Strings(slc6)
    sort.Ints(slc7)
    
    fmt.Println("After sorting: ")
    fmt.Println("slc6: ", slc6)
    fmt.Println("slc7: ", slc7)
    
}
    
    

// Slices
// Iterating slices
// https://www.geeksforgeeks.org/slices-in-golang/
