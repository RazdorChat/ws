package main

import (
	"log"
	"net"

	"github.com/RazdorChat/ws/server"
)

func main() {
	l, err := net.Listen("tcp", "localhost:3050")
	if err != nil {
		log.Fatal(err)
	}
	server.Start(l)
}
