package main

import (
	"context"
	"fmt"
	"log"
)

type keystore string

const sessionID keystore = "SESSIONID"

func main() {
	ctx := context.WithValue(context.Background(), sessionID, "SESSIONID_ session id")
	sessionID := ctx.Value(sessionID)
	str, ok := sessionID.(string)
	if !ok {
		log.Fatal("sessionID is not a string")
	}

	fmt.Println("value:", str)
	fmt.Println(sessionID)
}
