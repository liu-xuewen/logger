package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zLog *zap.Logger

func init() {
	hook := &lumberjack.Logger{
		Filename:   filepath.Join("logs/", filepath.Base(os.Args[0]+".log")), //filePath
		MaxSize:    512,                                                      // megabytes
		MaxBackups: 10000,
		MaxAge:     100000, //days
		Compress:   false,  // disabled by default
	}
	defer hook.Close() // todo

	enConfig := zap.NewProductionEncoderConfig() //生成配置

	enConfig.EncodeCaller = zapcore.FullCallerEncoder
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(enConfig), //编码器配置
		zapcore.AddSync(hook),            //打印到控制台和文件
		zap.InfoLevel,                    //日志等级
	)

	zLog = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// Info Info
func Info(ctx context.Context, msg string, args ...interface{}) {
	check(args)
	zLog.Info(msg, parseArgs(ctx, args)...)
}

// Warn Warn
func Warn(ctx context.Context, msg string, args ...interface{}) {
	check(args)
	zLog.Warn(msg, parseArgs(ctx, args)...)
}

// Error Error
func Error(ctx context.Context, msg string, args ...interface{}) {
	check(args)
	zLog.Error(msg, parseArgs(ctx, args)...)
}

func check(args []interface{}) {
	if len(args)%2 == 1 {
		panic(fmt.Sprintf("check:%v", args))
	}
}

func parseArgs(ctx context.Context, args []interface{}) (zf []zap.Field) {
	var ok bool
	str := ""
	for i, v := range args {
		if i%2 == 0 {
			str, ok = v.(string)
			if !ok {
				panic(args)
			}
		} else {

			zf = append(zf, zap.Any(str, v))
		}
	}
	return
}
