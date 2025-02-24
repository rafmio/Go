package main

import (
  "fmt"
  "io"
)

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error) {
  count := 0
  for i := 0; i < len(ph); i++ {
    if (ph[i] >= '0' && ph[i] <= '9') {
      p[count] = ph[i]
      count++
    }
  }
  return count, io.EOF  // io.EOF - маркер окончания чтения
}

func main() {
  phone1 := phoneReader("+1(234) 567 9010")
  phone2 := phoneReader("+2-345-678-42-21--")

  buffer := make([]byte, len(phone1))
  rd1, _ := phone1.Read(buffer)
  fmt.Println("readed:", rd1)
  fmt.Println(string(buffer))
  fmt.Println("len(buffer):", len(buffer), "cap(buffer):", cap(buffer))
  fmt.Println()

  buffer = make([]byte, len(phone2))
  rd2, _ := phone2.Read(buffer)
  fmt.Println("readed:", rd2)
  fmt.Println(string(buffer))
  fmt.Println("len(buffer):", len(buffer), "cap(buffer):", cap(buffer))
  for i, v := range buffer {
    fmt.Println(i, ":", v)
  }

}
