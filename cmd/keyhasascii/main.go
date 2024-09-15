package main

import (
	"log"
	"redis_tools/internal/cmd/keyhasascii"
)

func main() {
	err := keyhasascii.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
