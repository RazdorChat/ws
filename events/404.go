package events

import (
	"context"
	"io"

	"github.com/gobwas/ws/wsutil"
)

func UnknownEvent(_ context.Context, w *wsutil.Writer, r *wsutil.Reader) {
	io.WriteString(w, "error: unknown event")
}
