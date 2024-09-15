package main

import (
	"log"
	"redis_tools/internal/cmd/delkeys"
)

func main() {
	err := delkeys.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
