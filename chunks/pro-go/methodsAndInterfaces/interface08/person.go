package main

type Person struct {
  name, city string
}

func (p *Person) getPerson () {
  fmt.Println("Person:", p.name, "City:", p.city)
}
