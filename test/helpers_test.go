package test

import (
	"log"

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

// Panic if e != nil
func must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}

// Log e if e != nil
func should[T any](v T, e error) T {
	if e != nil {
		log.Println(e)
	}
	return v
}
