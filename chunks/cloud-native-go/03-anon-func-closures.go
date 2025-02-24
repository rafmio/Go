package main

import "fmt"

func incremetor() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	increment := incremetor()
	fmt.Println("increment():", increment())
	fmt.Println("increment():", increment())
	fmt.Println("imcrement():", increment())

	newIncrement := incremetor()
	fmt.Println("newIncrement():", newIncrement())
}

// Замыкание - это вложенная функция, сохраняющая доступ к переменным
// родительской функции даже после завершения родительской функции
