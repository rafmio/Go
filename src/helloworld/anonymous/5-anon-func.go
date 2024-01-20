// https://kovardin.ru/articles/go/zamykaniya/
package main

import "fmt"

func counter(start int) (func() int, func()) {
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	return ctr, incr
}

func main() {
	ctr1, incr1 := counter(100)
	ctr2, incr2 := counter(100)
	fmt.Println("counter1 = ", ctr1())
	fmt.Println("counter2 = ", ctr2())

	incr1()
	fmt.Println("counter1 = ", ctr1())
	fmt.Println("counter2 = ", ctr2())
	fmt.Println()

	incr2()
	incr2()

	fmt.Println("counter1 = ", ctr1())
	fmt.Println("counter2 = ", ctr2())
	fmt.Println()

	ctr3, incr3 := counter(300)
	fmt.Println("ctr1: ", ctr1())
	fmt.Println("ctr2: ", ctr2())
	fmt.Println("ctr3: ", ctr3())

	incr3()
	fmt.Println("Afger incr3():")
	fmt.Println("ctr1: ", ctr1())
	fmt.Println("ctr2: ", ctr2())
	fmt.Println("ctr3: ", ctr3())
}
