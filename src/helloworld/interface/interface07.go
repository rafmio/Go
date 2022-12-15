// https://habr.com/ru/company/vk/blog/463063/
package main

import (
  "fmt"
  "strconv"
  "log"
)

type Book struct {
  Title string
  Author string
}

func (b Book) String() string {
  return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

type Count int

func (c Count) String() string {
  return strconv.Itoa(int(c))
}

func WriteLog(s fmt.Stringer) {
  log.Println(s.String())
}

func main() {
  book := Book{"Alice in Wonderland", "Lewis Carrol"}
  WriteLog(book)

  count := Count(3)
  WriteLog(count)
}
