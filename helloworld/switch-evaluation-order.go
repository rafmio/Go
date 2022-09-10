package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	
	fmt.Println("When's Friday?")
    switch time.Friday {
        case today + 0:
            fmt.Println("Today")
        case today + 1:
            fmt.Println("In two days")
        default:
            fmt.Println("Too far away")
    }
    
    fmt.Println("Value of 'today' variable is: %g", today)
    whattime := time.Now()
    fmt.Println("Time is: %g", whattime)
}
