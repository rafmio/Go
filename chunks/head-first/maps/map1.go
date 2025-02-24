package main

import "fmt"

func main() {
  var ranks map[string]int
  ranks = make(map[string]int)
  ranks["gold"] = 1
  ranks["silver"] = 2
  ranks["bronze"] = 3

  fmt.Println(ranks["gold"])
  fmt.Println(ranks["silver"])

  ranks2 := make(map[string]string)
  ranks2["Kissy"] = "Missy"
  ranks2["Huggy"] = "Wuggy"

  fmt.Println(ranks2["Kissy"], ranks2["Huggy"])

  elements := make(map[string]string)
  elements["Hi"] = "Hydrogen"
  elements["Li"] = "Lithium"
  fmt.Println(elements["Hi"])
  fmt.Println(elements["Li"])

  ranks3 := map[string]int{"bronze" : 3, "silver" : 2, "gold" : 1}
  fmt.Println(ranks3["gold"], ranks3["silver"])

  creatures := map[string]string{"Huggy" : "Wuggy", "Tosi" : "Bosi"}
  fmt.Println(creatures)
  fmt.Println(creatures["Huggy"])
  fmt.Println(creatures["Tosi"])
}
