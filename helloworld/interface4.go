package main

import "fmt"

type greeter interface {
    greet(string) string
}

type russian struct{}
type american struct{}

func (r *russian) greet(name string) string {
    return fmt.Sprintf("Privet, %s\n", name)
}

func (a * american) greet(name string) string {
    return fmt.Sprintf("Hello, %s\n", name)
}

func sayHello(g greeter, name string) {
    fmt.Print(g.greet(name))
}

func main() {
    sayHello(&russian{}, "Petya")
    sayHello(&american{}, "Bill")
}

// Interfaces 
// https://youtu.be/tCN8ac6C1tA
