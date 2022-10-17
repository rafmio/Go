package main

import (
<<<<<<< HEAD
    "fmt"
    "log"
    "datafile"
)
=======
  "fmt"
  "datafile"
  "log"
)

func main() {
  numbers, err := datafile.GetFloats("/home/raf/go/Go/src/head-first/average/data.txt")
  if err != nil {
    log.Fatal(err)
  }

  var sum float64 = 0
  for _, number := range numbers {
    sum += number
  }

  sampleCount := float64(len(numbers))
  fmt.Printf("Average: %0.2f\n", sum / sampleCount)

  fmt.Printf("The type of numbers is : %T\n", numbers)
}
>>>>>>> c9916e36585592a71321373748e872f6582d4557
