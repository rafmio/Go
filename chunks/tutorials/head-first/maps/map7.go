package main

import "fmt"

func main() {
  grades := map[string]float64{"Alma": 74.2,
                               "Rohit": 86.5,
                               "Carl": 59.7,
                             }
   for name, grade := range grades {
     fmt.Printf("%s's\tgrade is: %0.2f\n", name, grade)
     }

  fmt.Println("Class roster: ")
  for name := range grades {
    fmt.Println(name)
  }

  fmt.Println()
  fmt.Println("Grades:")
  for _, grade := range grades {
    fmt.Println(grade)
  }
}

// Цикл for...ragne обрабатывает карты (maps) в случайном порядке
// Порядок вывода fmt.Println(grade) и пр. каждый раз будет разным
// Карта является неупорядоченной коллекцией ключей и значений
