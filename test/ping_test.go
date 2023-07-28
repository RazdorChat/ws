package test

import (
	"testing"

	"github.com/RazdorChat/ws/packet"
	"github.com/gobwas/ws"
)

func TestPing(t *testing.T) {
	t.Run("OpCode Ping", func(t *testing.T) {
		r.Discard()
		w.ResetOp(ws.OpPing)
		_, err := packet.EncodeEvent(w, "ping")
		if err != nil {
			t.Error(err)
		}
		w.Flush()

		header, err := r.NextFrame()
		if err != nil {
			t.Fatal(err)
		}
		checkOpCode(t, header.OpCode, ws.OpPong)
		var evt string
		if _, err = packet.DecodeEvent(r, int(header.Length), &evt); err != nil {
			t.Error(err)
		}
		checkEvent(t, evt, "pong")
	})
	t.Run("Text Ping", func(t *testing.T) {
		r.Discard()
		w.ResetOp(ws.OpText)
		_, err := packet.EncodeEvent(w, "ping")
		if err != nil {
			t.Error(err)
		}
		w.Flush()

		header, err := r.NextFrame()
		if err != nil {
			t.Fatal(err)
		}
		checkOpCode(t, header.OpCode, ws.OpText)
		var evt string
		if _, err = packet.DecodeEvent(r, int(header.Length), &evt); err != nil {
			t.Error(err)
		}
		checkEvent(t, evt, "pong")
	})
}
