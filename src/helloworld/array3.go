package main
import "fmt"

func main() {
    //--------------------------------------------------------
    my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "C"}
    my_arr2 := my_arr1
    
    fmt.Println("Array_1: ", my_arr1)
    fmt.Println("Array_2: ", my_arr2)
    
    my_arr1[0] = "C++"
    
    fmt.Println("\nArray_1: ", my_arr1)
    fmt.Println("Array_2: ", my_arr2)
    //--------------------------------------------------------
    my_arr3 := [6]int{12, 56, 654, 44, 10, 11}
    my_arr4 := &my_arr3
    
    fmt.Println("arr3: ",  my_arr3)
    fmt.Println("arr4: ", *my_arr4)
    
    my_arr3[5] = 100000
    
    fmt.Println("arr3: ",  my_arr3)
    fmt.Println("arr4: ", *my_arr4)
}
