package logger

import (
	"context"
	"os"
	"path/filepath"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zLog *zap.Logger

func Init(fileName string) {
	if fileName == "" {
		fileName = filepath.Join("logs/", filepath.Base(os.Args[0]+".log"))
	}

	hook := &lumberjack.Logger{
		Filename:   fileName, //filePath
		MaxSize:    512,      // megabytes
		MaxBackups: 10000,
		MaxAge:     100000, //days
		Compress:   false,  // disabled by default
	}
	defer hook.Close() // todo

	enConfig := NewProductionEncoderConfig() //生成配置
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(enConfig), //编码器配置
		zapcore.AddSync(hook),            //打印到控制台和文件
		zap.InfoLevel,                    //日志等级
	)

	zLog = zap.New(core)
}

// NewProductionEncoderConfig returns an opinionated EncoderConfig for
// production environments.
func NewProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      zapcore.OmitKey,
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// Info Info
func Info(ctx context.Context, msg string, args ...interface{}) {
	zLog.Info(msg, parseArgs(ctx, args)...)
}

// Warn Warn
func Warn(ctx context.Context, msg string, args ...interface{}) {
	zLog.Warn(msg, parseArgs(ctx, args)...)
}

// Error Error
func Error(ctx context.Context, msg string, args ...interface{}) {
	zLog.Error(msg, parseArgs(ctx, args)...)
}

func parseArgs(ctx context.Context, args []interface{}) (zf []zap.Field) {
	for _, key := range cfg.defaultCtxKeys {
		zf = append(zf, zap.Any(string(key), ctx.Value(key)))
	}

	for key, val := range cfg.defaultKeyValMap {
		zf = append(zf, zap.Any(key, val))
	}

	var ok bool
	str := ""
	for i, v := range args {
		if i%2 == 0 {
			str, ok = v.(string)
			if !ok {
				zf = append(zf, zap.Any("arg_"+strconv.Itoa(i), v))
			}
		} else {

			zf = append(zf, zap.Any(str, v))
		}
	}
	return
}
