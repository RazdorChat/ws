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
		if header.OpCode != ws.OpPong {
			t.Errorf("OpPing expects OpPong response, received %s", wsOpCodeName(header.OpCode))
		}
		var evt string
		should(packet.DecodeEvent(r, int(header.Length), &evt))

		if evt != "pong" {
			t.Errorf("Pong event expected, received %s", evt)
		}
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
		if header.OpCode != ws.OpText {
			t.Errorf("OpText expects OpText response, received %s", wsOpCodeName(header.OpCode))
		}
		var evt string
		should(packet.DecodeEvent(r, int(header.Length), &evt))

		if evt != "pong" {
			t.Errorf("Pong event expected, received %s", evt)
		}
	})
}
