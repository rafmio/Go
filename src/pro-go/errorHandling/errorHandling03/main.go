package main

import "fmt"

func main() {
	// recoveryFunc := func() {
	// 	if arg := recover(); arg != nil {
	// 		if err, ok := arg.(error); ok {
	// 			fmt.Println("Error:", err.Error())
	// 		} else if str, ok := arg.(string); ok {
	// 			fmt.Println("Message:", str)
	// 		} else {
	// 			fmt.Println("Panic recovered")
	// 		}
	// 	}
	// }
	// defer recoveryFunc()

	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message:", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	} ()

	categories := []string{"Watersports", "Chess", "Running"}

	channel := make(chan ChannelMessage, 10)

	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			// fmt.Println(message.Category, "(no such category)")
			panic(message.CategoryError)
		}
	}
}

// The type of the value returned by the recover function is
// the empty interface ( interface[] ), which requires a type assertion
