package logger

import "github.com/liu-xuewen/go-lib/ctxlib"

var defaultCtxKeys []ctxlib.String

// SetDefaultCtxKey set logger ctx key
func SetDefaultCtxKey(ctxKeys ...ctxlib.String) {
	defaultCtxKeys = append(defaultCtxKeys, ctxKeys...)
}
