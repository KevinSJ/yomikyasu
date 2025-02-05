package main

import (
	"fmt"
	"log"
	"yomikyasu/internal/server"
)

func main() {

	server := server.NewServer()

	log.Printf("Server will listen on %v ", server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
