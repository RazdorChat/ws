package test

import "github.com/gobwas/ws"

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
