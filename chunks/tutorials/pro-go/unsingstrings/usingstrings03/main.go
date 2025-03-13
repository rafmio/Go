package main

import "fmt"

func Printfln(template string, values ...interface{}) {
  fmt.Printf(template + "\n", values...)
}

func main() {
  var name string
  var category string
  var price float64

  fmt.Print("Enter text to scan: ")
  n, err := fmt.Scan(&name, &category, &price)

  if (err == nil) {
    Printfln("Scanned %v values", n)
    Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
  } else {
    Printfln("Error: %v", err.Error())
  }
  fmt.Println()

  vals := make([]string, 3)
  ivals := make([]interface{}, 3)
  for i := 0; i < len(vals); i++ {
    ivals[i] = &vals[i]
  }
  fmt.Print("Enter text to scan (3 elements): ")
  fmt.Scan(ivals...)
  Printfln("Name: %v", vals)

}
