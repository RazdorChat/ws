package server

import (
	"io"
	"log"
	"net"

	"github.com/RazdorChat/ws/packet"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func listen(conn net.Conn) (err error) {
	defer conn.Close()
	var (
		r      = wsutil.NewReader(conn, ws.StateServerSide)
		w      = wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
		header ws.Header
	)

	for {
		r.Discard()
		w.Flush()
		w.ResetOp(ws.OpText)
		// Read frame header
		header, err = r.NextFrame()
		// If the header can't be read, discard the unread message
		if err != nil {
			log.Println(err)
			continue
		}
		switch header.OpCode {
		case ws.OpClose:
			return io.EOF
		case ws.OpPing:
			w.ResetOp(ws.OpPong)
			packet.EncodeEvent(w, "pong")
			continue
		}

		payload, err := packet.Decode(r, header)
		if err != nil {
			log.Println(err)
			continue
		}

		// Process events
		switch payload.Event {
		case "ping":
			packet.EncodeEvent(w, "pong")
		}
	}
}
