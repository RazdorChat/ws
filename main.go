package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"regexp"
	"strings"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Payload struct {
	Event string `json:"event"`
	Body  []byte `json:"body,omitempty"`
}

var eventRegex *regexp.Regexp = regexp.MustCompile("")

func pollConnection(conn net.Conn) (err error) {
	defer conn.Close()
	var (
		r      = wsutil.NewReader(conn, ws.StateServerSide)
		w      = wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
		header ws.Header
	)
	const payloadDelim = '\n'
	for {
		r.Discard()
		// Read packet header
		header, err = r.NextFrame()
		// If the header can't be read, discard the unread message
		if err != nil {
			continue
		}
		if header.OpCode == ws.OpClose {
			return io.EOF
		}

		// TODO: set max payload size
		// Parse event type
		br := bufio.NewReader(r)
		payloadHeader, err := br.ReadString(payloadDelim)
		if err != nil {
			continue // Malformed/unterminated payload header
		}
		payloadHeaderParts := strings.SplitN(payloadHeader, ":", 1)
		if len(payloadHeaderParts) < 2 {
			r.Discard()
			continue
		}
		event := strings.Trim(payloadHeaderParts[1], " ")

		// Process events
		switch event {
		case "ping":
			io.WriteString(w, "event: pong\n\n")
		}
	}
}

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
		go pollConnection(conn)
	}
}
