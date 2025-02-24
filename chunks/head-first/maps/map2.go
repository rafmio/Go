package main

import "fmt"

func main() {
  jewelry := make(map[string]float64)
  jewelry["necklace"] = 89.99
  jewelry["earrings"] = 79.99
  clothing := map[string]float64{"pants" : 59.99, "shirt" : 39.99}

  fmt.Println("Earrings :", jewelry["earrings"])
  fmt.Println("Necklace :", jewelry["necklace"])
  fmt.Println("Shirt :", clothing["shirt"])
  fmt.Println("Pants: ", clothing["pants"])

  numbers := make(map[string]int)
  numbers["I've been assigned"] = 12
  fmt.Printf("%#v\n", numbers["I've been assigned"])
  fmt.Printf("%#v\n", numbers["I haven't beed assigned"])

  counters := make(map[string]int)
  fmt.Println(counters["a"], counters["b"], counters["c"])
  counters["a"]++
  counters["a"]++
  counters["c"]++
  fmt.Println(counters["a"], counters["b"], counters["c"])
}
