package main

import (
	"app/apiserver"
	"log"
)

func main() {
	server := apiserver.NewServer()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
