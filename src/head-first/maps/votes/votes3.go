package main

import (
	"fmt"
	"votesdata"
	"log"
)

func main() {
	lines, err := votesdata.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
