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
    
}
