package main

import (
	"log"
	"redis_tools/internal/cmd/keyttl"
)

func main() {
	err := keyttl.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
