package logger

import "github.com/liu-xuewen/go-lib/ctxlib"

var cfg *config

func init() {
	cfg = &config{
		defaultKeyValMap: make(map[string]interface{}),
	}
}

type config struct {
	defaultCtxKeys   []ctxlib.String
	defaultKeyValMap map[string]interface{}
}

// SetDefaultCtxKey set logger ctx key
func SetDefaultCtxKey(ctxKeys ...ctxlib.String) {
	cfg.defaultCtxKeys = append(cfg.defaultCtxKeys, ctxKeys...)
}

// SetDefaultKeyVal set default key value
func SetDefaultKeyVal(m map[string]interface{}) {
	for k, v := range m {
		cfg.defaultKeyValMap[k] = v
	}
}
