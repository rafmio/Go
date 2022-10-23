package main

import (
    "fmt"
    "modvar"
)

func main() {
    var a int = 10
    var b int = 20
    c := modvar.ExportFunc(a, b)
    fmt.Println("c = ", c)
    
    d := modvar.ExportedVar * c
    fmt.Println("d = ", d)
    
    var s modvar.Subscriber
    s.Rate = 4.99
    fmt.Println(s.Rate)
    
    var subsriber modvar.Subscriber
    subsriber.Name = "Aman Singh"
    subsriber.Rate = 4.99
    subsriber.Active = true
    fmt.Println(subsriber.Name, subsriber.Rate, subsriber.Active)
    
    subs2 := modvar.Subscriber{Name: "John Lennon", Rate: 4.99, 
        Active: true}
    fmt.Println("Name: ", subs2.Name)
    fmt.Println("Rate: ", subs2.Rate)
    fmt.Println("Active: ", subs2.Active)
    
    var employee modvar.Employee
    employee.Name = "Joy Carr"
    employee.Salary = 60000
    fmt.Println()
    fmt.Println(employee.Name)
    fmt.Println(employee.Salary)
    
    var address modvar.Address
    address.Street = "123 Oak St"
    address.City = "Omaha"
    address.State = "NE"
    address.PostalCode = "68111"
    fmt.Println(address)
    
}
