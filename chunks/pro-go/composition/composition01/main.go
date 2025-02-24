package main

import (
  "fmt"
  "composition01/store"
)

func main() {
  kayak := store.NewProduct("Kayak", "Watersports", 275.00)
  lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}
  submarine := &store.Product{Name: "Nuclear submarine", Category: "Watersports"}

  for _, p := range []*store.Product{ kayak, lifejacket, submarine } {
    fmt.Println("Name:", p.Name, "Category:", p.Category,
    "Price:", p.Price(0.2))
  }
  fmt.Println()

  boats := []*store.Boat {
    store.NewBoat("Kayakus", 275, 1, false),
    store.NewBoat("Canoe", 400, 3, false),
    store.NewBoat("Tender", 650.25, 2, true),
  }

  for _, b := range boats {
    fmt.Println("Conventional:", b.Product.Name, "| Direct:", b.Name)
  }
  fmt.Println()

  for _, b := range boats {
    fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
  }
  fmt.Println()

// Creating a chain of nested types
  rentals := []*store.RentalBoat{
    store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
    store.NewRentalBoat("Yacht", 50_000, 5, true, true, "Bob", "Alice"),
    store.NewRentalBoat("Super Yacht", 100_000, 15, true, true, "Dora", "Charlie"),
  }

  for i, r := range rentals {
    fmt.Println(i, "-:-", "RentalBoat:", r.Name, "Rental Price:", r.Price(0.2))
  }
  fmt.Println()

  for i, r := range rentals {
    fmt.Println(i, "-:-", "Category:" ,r.Category, "Crew:", r.IncludeCrew,
    "Motorized:", r.Motorized)
  }
  fmt.Println()

  for i, r := range rentals {
    fmt.Println(i, "-:-", "Rental Boat:", r.Name, "Rental Price:", r.Price(0.2),
  "Captain:", r.Captain)
  }
  fmt.Println()

// Understanding when promotion cannot be performed
  product := store.NewProduct("Schooner", "Watersports", 300.00)
  deal := store.NewSpecialDeal("Weekend Special", product, 50)

  Name, price, Price := deal.GetDetails()

  fmt.Println("Name:", Name)
  fmt.Println("Price field:", price)
  fmt.Println("Price method:", Price) // price field of *Product, not *SpecialDeal

}
