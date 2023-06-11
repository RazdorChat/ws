package main

import (
	"log"
	"net"

	"github.com/gobwas/ws"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3050")
	if err != nil {
		log.Fatal(err)
	}
	upgrader := ws.Upgrader{
		OnHeader: func(key, value []byte) error {
			// TODO: get session from Cookie header
			return nil
		}}
	// TODO: run listener in goroutine

	log.Println("Listening on localhost:3050")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// TODO: process handshake result
		_, err = upgrader.Upgrade(conn)
		if err != nil {
			log.Println("Upgrade error:", err)
		}
		go listen(conn)
	}
}
