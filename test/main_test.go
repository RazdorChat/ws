package test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/RazdorChat/ws/server"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// Websoket connection r/w
var r *wsutil.Reader
var w *wsutil.Writer

func TestMain(m *testing.M) {
	// Exit with test result code
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// Start server
	const port = 3050
	go func() {
		l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
		if err != nil {
			log.Fatal(err)
		}
		server.Start(l)
	}()
	time.Sleep(10 * time.Millisecond) // Add 10ms delay for server to start

	// Connect to websocket
	conn, _, _, err := ws.Dial(context.Background(), fmt.Sprintf("ws://localhost:%d", port))
	if err != nil {
		fmt.Printf("Error connecting to websocket: %s\n", err)
		exitCode = 1
		return
	}

	r = wsutil.NewReader(conn, ws.StateClientSide)
	w = wsutil.NewWriter(conn, ws.StateClientSide, ws.OpText)

	m.Run()
}
