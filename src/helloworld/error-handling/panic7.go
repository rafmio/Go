package main

import "fmt"

type Shark struct {
    Name string
}

func (s *Shark) SayHello() {
    fmt.Println("Hi! My name is", s.Name)
}

func main() {
    s := &Shark{"Sammy"}
    s = nil         // After this Go will generate panic
    s.SayHello()
}

// panic: runtime error: invalid memory address or nil pointer 
// dereference
