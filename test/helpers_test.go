package test

import (
	"testing"

	"github.com/gobwas/ws"
)

func wsOpCodeName(op ws.OpCode) string {
	switch op {
	case ws.OpText:
		return "OpText"
	case ws.OpBinary:
		return "OpBinary"
	case ws.OpClose:
		return "OpClose"
	case ws.OpPing:
		return "OpPing"
	case ws.OpPong:
		return "OpPong"
	}
	return "unknown ws.OpCode"
}

func checkOpCode(t *testing.T, actual ws.OpCode, expected ws.OpCode) {
	if actual != expected {
		t.Errorf("Expected OpCode %s, received %s", wsOpCodeName(expected), wsOpCodeName(actual))
	}
}

func checkEvent(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("Expected event %s, received %s", expected, actual)
	}
}

// Panic if e != nil
func must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}
