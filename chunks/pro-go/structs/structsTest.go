package main

import "fmt"

type Human struct {
  name string
  age int
}

type Student struct {
  *Human
  id int
}

func (stud *Student) printStudDetails() {
  fmt.Println(stud)
  fmt.Println(stud.name)
  fmt.Println(stud.age)
  fmt.Println(stud.id)
}

func main() {
  human := &Human{ "Peter Green", 21 }
  student := &Student{ human, 1122 }

  student.printStudDetails()

  student.name = "Neil Young"

  student.printStudDetails()
}
