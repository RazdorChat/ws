package server

import (
	"log"
	"net"

	"github.com/gobwas/ws"
)

var upgrader = ws.Upgrader{
	OnHeader: func(key, value []byte) error {
		// TODO: get session from Cookie header
		return nil
	}}

func Start(l net.Listener) {
	log.Printf("Listening on %s\n", l.Addr().String())
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// TODO: process handshake result
		_, err = upgrader.Upgrade(conn)
		if err != nil {
			log.Println("WS upgrade error:", err)
			conn.Close()
			continue
		}
		go listen(conn)
	}
}
