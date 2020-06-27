package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logCfg config

type config struct {
	path     string
	fileName string
}

const (
	// FatalLevel FatalLevel
	FatalLevel = "fatal"
	PanicLevel = "panic"
	ErrorLevel = "error"
	WarnLevel  = "warn"
	InfoLevel  = "info"
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

/**
 * LogConf 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func LogConf() {
	now := time.Now()
	hook := &lumberjack.Logger{
		Filename:   fmt.Sprintf("log/%04d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()), //filePath
		MaxSize:    500,                                                                                                                     // megabytes
		MaxBackups: 10000,
		MaxAge:     100000, //days
		Compress:   false,  // disabled by default
	}
	defer hook.Close()
	/*zap 的 Config 非常的繁琐也非常强大，可以控制打印 log 的所有细节，因此对于我们开发者是友好的，有利于二次封装。
	  但是对于初学者则是噩梦。因此 zap 提供了一整套的易用配置，大部分的姿势都可以通过一句代码生成需要的配置。
	*/
	enConfig := zap.NewProductionEncoderConfig() //生成配置

	// 时间格式
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	level := zap.InfoLevel

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig), //编码器配置
		zapcore.AddSync(hook),               //打印到控制台和文件
		level,                               //日志等级
	)

	_ = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	//_log := log.New(hook, "", log.LstdFlags)
	//gb.Logger = logger.Sugar()
	//_log.Println("Start...")
}
