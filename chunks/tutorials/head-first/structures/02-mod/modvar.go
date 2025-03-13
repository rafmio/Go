package modvar

var ExportedVar int = 10

func ExportFunc (a, b int) int {
    return a + b
}

type Subscriber struct {
    Name string
    Rate float64
    Active bool
    HomeAddress Address
}

type Employee struct {
    Name string
    Salary float64
    Street string
    HomeAddress Address
}

type Address struct {
    Street string
    City string
    State string
    PostalCode string
}
