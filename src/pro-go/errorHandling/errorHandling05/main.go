package main

import "fmt"

type CategoryCountMessage struct {
	Category string
	Count int
	TerminalError interface{}
}

func processCategory(categories []string, outChan chan<- CategoryCountMessage){
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			outChan <- CategoryCountMessage {
				TerminalError: arg,
			}
			close(outChan)
		}
	}()

	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessage {
				Category: message.Category,
				Count: int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}

func main() {
	categories := []string {"Watersports", "Chess", "Running", "Warship Games"}

	channel := make(chan CategoryCountMessage)
	go processCategory(categories, channel)

	for message := range channel {
		if (message.TerminalError == nil) {
			fmt.Println(message.Category, "Total:", message.Count)
		} else {
			fmt.Println("A terminal error occured")
		}
	}
}

// The type of the value returned by the recover function is
// the empty interface ( interface[] ), which requires a type assertion
