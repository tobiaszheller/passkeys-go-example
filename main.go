package main

import (
	"log"

	"github.com/tobiaszheller/passkeys-go-example/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
