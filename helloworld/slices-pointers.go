package main
import "fmt"

func main() {
    names := [6] string{
        "Kissy",
        "Missy",
        "Huggy",
        "Wuggy",
        "Tosy",
        "Bosy",
    }
    
    fmt.Println(names)
    
    huggy_wuggy := names[1:3]
    tosy_bosy   := names[4:6]
    
    fmt.Println(huggy_wuggy, tosy_bosy)
    
    tosy_bosy[1] = "Kuzmin"
    fmt.Println(huggy_wuggy, tosy_bosy)
    fmt.Println(names)
}

// Slices are like references to arrays
// https://go.dev/tour/moretypes/8
// A slice goes not store any data, it just describes a 
// section of an underlying array
// Changing the elements of its underlying array
// Other slices that share the same underlying array will see 
// those changes
