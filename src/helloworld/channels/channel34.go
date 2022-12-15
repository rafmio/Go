// https://habr.com/ru/post/490336/
// Fan-in, Fan-out
package main

import (
  "fmt"
  "sync"
)


func getInputChan() <-chan int {
  input := make(chan int, 100)

  numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

  go func() {
    for num := range numbers {
      input <- num
    }

    close(input)
  }()
  return input
}

func getSquareChan(input <-chan int) <-chan int {
  output := make(chan int, 100)

  go func() {
    for num := range input {
      output <- num * num
    }
    close(output)
  }()
  return output
}

// returns a merged channel of `outputsChan` channels
// this produce fan-in channel
func merge(outputsChan ...<-chan int) <-chan int {
  var wg sync.WaitGroup

  merged := make(chan int, 100)

  // increase counter to number of channels `len(outputsChan)`
  // as we will spawn number of goroutines to number of channels
  // reseived to merge
  fmt.Println("len(outputsChan)", len(outputsChan))
  wg.Add(len(outputsChan))

  // function than accept a channel (which sends square numbers) to
  // push numbers to merged channel
  output := func(sc <-chan int) {
    // run until channel (square numbers sender) closes
    fmt.Printf("sc: %#v, type of sc: %T\n", sc, sc)
    for sqr := range sc {
      merged <- sqr
    }

    wg.Done()
  }

  for _, optChan := range outputsChan {
    go output(optChan)
  }

  go func() {
    // wait until WaitGroup finishes
    wg.Wait()
    close(merged)
  } ()

  return merged
}

func main() {
  chanInputNums := getInputChan()

  chanOptSqr1 := getSquareChan(chanInputNums)
  chanOptSqr2 := getSquareChan(chanInputNums)

  chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)

  sqrSum := 0

  for num := range chanMergedSqr {
    sqrSum += num
  }

  fmt.Println("Sum of squares between 0-9 is", sqrSum)
}
