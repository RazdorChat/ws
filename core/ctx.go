package core

import (
	"context"

	"github.com/gobwas/ws"
)

type key string

func SetLength(ctx context.Context, length int) context.Context {
	return context.WithValue(ctx, key("length"), length)
}
func GetLength(ctx context.Context) (length int) {
	return ctx.Value(key("length")).(int)
}

func SetOpCode(ctx context.Context, opcode ws.OpCode) context.Context {
	return context.WithValue(ctx, key("opcode"), opcode)
}
func GetOpCode(ctx context.Context) (opcode ws.OpCode) {
	return ctx.Value(key("opcode")).(ws.OpCode)
}
