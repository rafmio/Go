package modvar

type Subscriber struct {
    Name string
    Rate float64
    Active bool
    Address     // Anonymous struct fields
}

type Employee struct {
    Name string
    Salary float64
    Street string
    Address     // Anonymous struct fields
}

type Address struct {
    Street string
    City string
    State string
    PostalCode string
}
