package main

func d1() {
	for i := 3; i > 0; i-- {
		defer print(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			print(i, " ")
		}()
	}
	println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	println()
	d3()
	println()
}
