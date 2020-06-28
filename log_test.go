package logger

import (
	"errors"
	"testing"
)

func TestErrorLog(t *testing.T) {
	err := errors.New("test error")
	Error("test err", "err", err)
	err = nil
	Error("test nil", "err", err)
}
