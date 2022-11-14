package main

import (
  "fmt"
  "math"
  "strconv"
)

func main() {
  price, tax := 275.00, 27.40
  sum := price + tax
  difference := price - tax
  product := price * tax
  quotient := price / tax

  fmt.Println(sum)
  fmt.Println(difference)
  fmt.Println(product)
  fmt.Println(quotient)
  fmt.Println()

  var intVal = math.MaxInt64
  var floatVal = math.MaxFloat64
  fmt.Println(intVal * 2)
  fmt.Println(floatVal * 2)
  fmt.Println(math.IsInf((floatVal * 2), 0))
  fmt.Println()

  posResult := 3 % 2
  negResult := -3 % 2
  absResult := math.Abs(float64(negResult))
  fmt.Println(posResult)
  fmt.Println(negResult)
  fmt.Println(absResult)
  fmt.Println()

  value := 10.2
  value++
  fmt.Println(value)
  value += 2
  fmt.Println(value)
  value -= 2
  fmt.Println(value)
  value--
  fmt.Println(value)
  fmt.Println()

  greeting := "Hello"
  language := "Go"
  combinedString := greeting + ", " + language
  fmt.Println(combinedString)
  fmt.Println()

  first := 100
  const second = 200.00
  equal := first == second
  notEqual := first != second
  lessThan := first < second
  lessThanOrEqual := first <= second
  greaterThan := first > second
  greaterThanOrEqual := first >= second
  fmt.Println(equal)
  fmt.Println(notEqual)
  fmt.Println(lessThan)
  fmt.Println(lessThanOrEqual)
  fmt.Println(greaterThan)
  fmt.Println(greaterThanOrEqual)
  fmt.Println()

  first2 := 100
  second2 := &first2
  third2 := &first2
  alpha2 := 100
  beta2 := &alpha2
  fmt.Println(second2 == third2)
  fmt.Println(second2 == beta2)
  fmt.Println(*second2 == *third2)
  fmt.Println(*second2 == *beta2)
  fmt.Println()

  maxMph := 50
  passengerCapacity := 4
  airbags := true
  familyCar := passengerCapacity > 2 && airbags
  sportsCar := maxMph > 100 || passengerCapacity == 2
  canCategorize := !familyCar && !sportsCar
  fmt.Println(familyCar)
  fmt.Println(sportsCar)
  fmt.Println(canCategorize)
  fmt.Println()

  kayak := 275
  soccerBall := 19.50
  total := float64(kayak) + soccerBall
  fmt.Println(total)
  fmt.Println()

  kayak2 := 275
  soccerBall2 := 19.50
  total2 := kayak2 + int(soccerBall2)
  fmt.Println(total2)
  fmt.Println(int8(total2))
  fmt.Println()

  kayak3 := 275
  soccerBall3 := 19.50
  total3 := kayak3 + int(math.Round(soccerBall3))
  fmt.Println(total3)
  fmt.Println()

  val1 := "true"
  val2 := "False"
  val3 := "not TRUE" // invalid syntax
  bool1, b1err := strconv.ParseBool(val1)
  bool2, b2err := strconv.ParseBool(val2)
  bool3, b3err := strconv.ParseBool(val3)
  if b3err == nil {
    fmt.Println("Parsed value:", bool3)
  } else {
    fmt.Println("Cannot parse", val3)
  }
  fmt.Println("Bool 1", bool1, b1err)
  fmt.Println("Bool 2", bool2, b2err)
  fmt.Println("Bool 3", bool3, b3err)
  fmt.Printf("Types of bool1: %T, bool2: %T, bool3: %T\n", bool1, bool2, bool3)
  fmt.Println()

  val4 := "0"
  if bool4, b4err := strconv.ParseBool(val4); b4err == nil {
    fmt.Println("Parsed value:", bool4)
  } else {
    fmt.Println("Cannot parse", val4)
  }

}
