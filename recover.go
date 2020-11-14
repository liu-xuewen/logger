package logger

import (
	"context"
	"runtime"
)

// Recover recover full stack
func Recover(ctx context.Context) {
	if err := recover(); err != nil {
		// 完整的堆栈信息
		Error(ctx, "panic_recover", "err", err, "full_stack", string(FullStack()))
	}
}

// RecoverRet recover full stack
func RecoverRet(ctx context.Context) interface{} {
	if r := recover(); r != nil {
		// 完整的堆栈信息
		Error(ctx, "panic_recover", "recover_err", r, "full_stack", string(FullStack()))
		return r
	}
	return nil
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
