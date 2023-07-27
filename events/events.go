package events

import (
	"context"

	"github.com/gobwas/ws/wsutil"
)

type Event struct {
	// Event handlers should always write a response before returning
	// w.Flush() is deferred to be run after the handler
	handler func(ctx context.Context, w *wsutil.Writer, r *wsutil.Reader)
}

func (e *Event) Handler(ctx context.Context, w *wsutil.Writer, r *wsutil.Reader) {
	defer w.Flush()
	e.handler(ctx, w, r)
}

// Map of event handlers
var Events = make(map[string]*Event)

func init() {
	Events["ping"] = &Event{
		handler: Ping}
}
