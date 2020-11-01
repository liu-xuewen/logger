package logger

import (
	"context"
	"runtime"
)

// Recover Recover full stack
func Recover(ctx context.Context) {
	if r := recover(); r != nil {
		// 完整的堆栈信息
		fullStack := string(FullStack())
		Error(ctx, "recover_err", r, "recover_err", fullStack)
	}
}

// FullStack get full stack
func FullStack() []byte {
	buf := make([]byte, 2048)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}

		buf = make([]byte, 2*len(buf))
	}
}
