package main

import (
	"fmt"
)

privateKey, err := ecdsa.GenerateKey(elliptic.p256(), rand.Reader)
if err != nil {
	log.Fatal("Failed to generate private key: %v", err)
}