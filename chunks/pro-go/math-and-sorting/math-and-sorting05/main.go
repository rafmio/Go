// Sorting data
package main

import (
  "fmt"
  "sort"
  "math/rand" // for sorting
  "time"      // for sorting
)

func main() {
  ints := []int{9, 4, 2, -1, 10}
  Printfln("Ints: %v", ints)
  fmt.Println("Are ints sorted? :", sort.IntsAreSorted(ints))
  sort.Ints(ints)
  Printfln("Ints Sorted: %v", ints)
  fmt.Println()

  floats := []float64{279, 48.95, 19.50, 22.41, 14.58}
  Printfln("Floats: %v", floats)
  fmt.Println("Are floats sorted? :", sort.Float64sAreSorted(floats))
  sort.Float64s(floats)
  Printfln("Floats sorted: %v", floats)
  fmt.Println()

  strings := []string{"Kayak", "Lifejacket", "Stadium", "Stick", "Ball"}
  Printfln("Strings: %v", strings)
  fmt.Println("Are strings sorted? :", sort.StringsAreSorted(strings))
  if(!sort.StringsAreSorted(strings)) {
    sort.Strings(strings)
    Printfln("Strings Sorted: %v", strings)
  } else {
    Printfln("Strings Already Sorted: %v", strings)
  }
  fmt.Println()
  // ----------------------------------------------------
  rand.Seed(time.Now().UnixNano())
  rand.Shuffle(len(ints), func(first, second int) {
    ints[first], ints[second] = ints[second], ints[first]
  })
  fmt.Println("ints:", ints)

  sortedInts := make([]int, len(ints))
  copy(sortedInts, ints)
  sort.Ints(sortedInts)
  Printfln("Ints Sorted: %v", sortedInts)
  fmt.Println()

  // -----------------------------------------------------
  indexof4 := sort.SearchInts(sortedInts, 4)
  indexof3 := sort.SearchInts(sortedInts, 9)
  Printfln("Index of 4: %v", indexof4)
  Printfln("Index of 3: %v", indexof3)
  fmt.Println()

  Printfln("Index of 4: %v (present: %v)", indexof4, sortedInts[indexof4] == 4)
  Printfln("Index of 3: %v (present: %v)", indexof3, sortedInts[indexof3] == 3)
}
