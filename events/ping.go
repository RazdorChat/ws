package events

import (
	"context"

	"github.com/RazdorChat/ws/core"
	"github.com/RazdorChat/ws/packet"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func Ping(ctx context.Context, w *wsutil.Writer, r *wsutil.Reader) {
	opcode := core.GetOpCode(ctx)
	if opcode == ws.OpPing {
		w.ResetOp(ws.OpPong)
		packet.EncodeEvent(w, "pong")
		w.Flush()
		w.ResetOp(ws.OpText)
	} else {
		packet.EncodeEvent(w, "pong")
	}
}
