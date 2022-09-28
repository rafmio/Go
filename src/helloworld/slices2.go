package main
import "fmt"

func main() {
    // Example 1 -----------------------------------------------
    arr1 := [7]string{"This", "is", "the", "tutorial", "of", "Go",  "language"}
    fmt.Println("Array: ", arr1)
    
    myslice := arr1[1:6]
    fmt.Println("Slice: ", myslice)
    fmt.Printf("Length of the slice: %d\n", len(myslice))
    fmt.Printf("Capacity of the slice: %d\n", cap(myslice))
    
    fmt.Println("---------------------------------------------")
    // Example 2 -----------------------------------------------
    var my_slice_1 = []string{"Kissy", "Missy", "Huggy", "Waggy"}
    fmt.Println("my_slice_1: ", my_slice_1)
    
    my_slice_2 := []int{12, 43, 242, 483, 2, 72, 614, 14}
    fmt.Println("my_slice_2: ", my_slice_2)
    
    fmt.Println("---------------------------------------------")
    // Example 3 -----------------------------------------------
    arr2 := [4]string{"Kissy", "Missy", "Huggy", "Waggy"}
    var my_slice_3 = arr2[1:2]
    my_slice_4 := arr2[0: ]
    my_slice_5 := arr2[ :2]
    my_slice_6 := arr2[ : ]
    
    fmt.Println("arr2: ", arr2)
    fmt.Println("my_slice_3: ", my_slice_3)
    fmt.Println("my_slice_4: ", my_slice_4)
    fmt.Println("my_slice_5: ", my_slice_5)
    fmt.Println("my_slice_6: ", my_slice_6)
 
    fmt.Println("---------------------------------------------")
    // Example 3 -----------------------------------------------
    oRignAl_slice := []int{34, 75, 23, 5, 34, 234, 73}
    var my_slice_7 = oRignAl_slice[1:5]
    my_slice_8  := oRignAl_slice[0: ]
    my_slice_9  := oRignAl_slice[ :6]
    my_slice_10 := oRignAl_slice[ : ]
    my_slice_11 := oRignAl_slice[2:4]
    
    fmt.Println("oRignAl_slice: ", oRignAl_slice)
    fmt.Println("my_slice_7:  ", my_slice_7)
    fmt.Println("my_slice_8:  ", my_slice_8)
    fmt.Println("my_slice_9:  ", my_slice_9)
    fmt.Println("my_slice_10: ", my_slice_10)
    fmt.Println("my_slice_11: ", my_slice_11)
 
     fmt.Println("---------------------------------------------")
    // Example 4 -----------------------------------------------
     var my_slice_12 = make([]int, 4, 7)
     fmt.Printf("my_slice_12 = %v, length = %d, capacity = %d\n", my_slice_12, len(my_slice_12), cap(my_slice_12))
     
     var my_slice_13 = make([]int, 7)
     fmt.Printf("my_slice_13 = %v, length = %d, capacity = %d\n", my_slice_13, len(my_slice_13), cap(my_slice_13))

    fmt.Println("---------------------------------------------")
    // Example 4 -----------------------------------------------
}
