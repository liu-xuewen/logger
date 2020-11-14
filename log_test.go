package logger

import (
	"context"
	"errors"
	"testing"
)

func TestErrorLog(t *testing.T) {
	err := errors.New("test error")
	ctx:=context.Background()
	Error(ctx,"test err", "err", err)
	err = nil
	Error(ctx,"test nil", "err", err)
}
