package server

import (
	"context"
	"log"
	"net"

	"github.com/RazdorChat/ws/core"
	"github.com/RazdorChat/ws/events"
	"github.com/RazdorChat/ws/packet"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// Listen to events on a ws conn
func listen(conn net.Conn) (err error) {
	defer conn.Close()
	var (
		r      = wsutil.NewReader(conn, ws.StateServerSide)
		w      = wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
		header ws.Header
	)

	for {
		// Read next frame
		header, err = r.NextFrame()
		// If the header can't be read, discard the unread message
		if err != nil {
			w.Flush()
			log.Println(err)
			continue
		}
		evt := new(string)
		length := int(header.Length)
		if n, err := packet.DecodeEvent(r, length, evt); err != nil {
			w.Flush()
			log.Println(err)
			continue
		} else {
			// Subtract decoded length
			length -= n
		}

		// Prepare ctx
		ctx := context.Background()
		ctx = core.SetLength(ctx, length)
		ctx = core.SetOpCode(ctx, header.OpCode)

		// If a handler exists for the event, run it
		if event := events.Events[*evt]; event != nil {
			event.Handler(ctx, w, r)
		}
	}
}
