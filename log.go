package logger

import (
	"log"
	"os"
	"path/filepath"
)

var logCfg config

type config struct {
	path     string
	fileName string
}

const (
	// FatalLevel FatalLevel
	FatalLevel = "fatal"
)

func init() {
	logCfg = config{
		path: os.TempDir(),
	}
}

// InitLogCfg InitLogCfg
func InitLogCfg(path string, skip bool) {
	if path != "" {
		if !CheckPathExist(path) {
			if err := os.MkdirAll(path, 0755); err != nil {
				log.Fatalf("mkdir log path:%s err:%s \n", path, err)
				return
			}
		}
		logCfg.path = path
		logCfg.fileName = filepath.Join(logCfg.path, filepath.Base(os.Args[0]+".log"))
		log.Println("logCfg.fileName:", logCfg.fileName)
	}
}
