// Switch Statement in Go
// Types of switch statements:
// 1. Expression Switch
// 2. Type Switch
// 
// Expression switch is similar in C, C++, Java. 
// Example 1:
    switch day := 4; day {
        case 1:
            fmt.Println("Monday")
        case 2:
            fmt.Println("Tuesday")
        case 3: 
            fmt.Println("Wednesday")
        case 4:
            fmt.Println("Thursday")
        case 5:
            fmt.Println("Friday")
        case 6:
            fmt.Println("Saturday")
        case 7:
            fmt.Println("Sunday")
        default:
            fmt.Println("Invalid")

// Example 2:
    var value int = 2
    switch {
        case value == 1:
            fmt.Println("One")
        case value == 2:
            fmt.Println("Two")
        case value == 3:
            fmt.Println("Three")
        default:
            fmt.Println("Invalid")
    }
    
// Example 3:
    var value string = "five"
    switch value {
        case "one":
            fmt.Println("One")
        case "two", "three":
            fmt.Println("Two and Three")
        case "four", "five", "six":
            fmt.Println("Four, Five, Six")
    }
}

// Type Switch
// Type switch is used when you want to compare types
var value interface{}
switch q := value.(type) {
    case bool: 
        fmt.Println("Value is of boolean type")
    case float64:
        fmt.Println("Value is of float64 type")
    case int:
        fmt.Println("Value is fo int type")
    default:
        fmt.Println("Value is of type %T\n")
}
