package main

import (
	"log"
	"redis_tools/internal/cmd/keysize"
)

func main() {
	err := keysize.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
