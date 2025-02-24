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
  n, err := fmt.Scanln(&name, &category, &price)

  if (err == nil) {
    Printfln("Scanned %v values", n)
    Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
  } else {
    Printfln("Error: %v", err.Error())
  }
  fmt.Println()

  var name2, category2 string
  var price2 float64
  source := "Lifejacket Watersports 48.95"
  n, err = fmt.Sscan(source, &name2, &category2, &price2)
  if (err == nil) {
    Printfln("Scanned %v values", n)
    Printfln("Name: %v, Category: %v, Price: %.2f", name2, category2, price2)
  } else {
    Printfln("Error: %v", err.Error())
  }
  fmt.Println()

  var name3, category3 string
  var price3 float64

  source3 := "Product Lifejacket Watersports 48.95"
  template3 := "Product %s %s %f"
  n, err = fmt.Sscanf(source3, template3, &name3, &category3, &price3)

  if (err == nil) {
    Printfln("Scanned %v values", n)
    Printfln("Name: %v, Category: %v, Price: %.2f", name3, category3, price3)
  } else {
    Printfln("Error: %v", err.Error())
  }

}
