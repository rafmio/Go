package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
    var minute int = now.Minute()
    
	fmt.Println(year)
	fmt.Println(now)
    fmt.Println(minute)

	fmt.Println("Type of now is:  ", reflect.TypeOf(now))
    fmt.Println("Type of year is: ", reflect.TypeOf(year))
    
   
    
}
