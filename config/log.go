package config

import (
	"MyGin/global"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogConfig struct {
	Level      string `mapstructure:"level"`      // 日志级别
	Filename   string `mapstructure:"filename"`   // 日志文件路径
	MaxSize    int    `mapstructure:"maxsize"`    // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int    `mapstructure:"maxbackups"` // 日志文件最多保存多少个备份
	MaxAge     int    `mapstructure:"maxage"`     // 文件最多保存多少天
	Compress   bool   `mapstructure:"compress"`   // 是否压缩
}

var LogConf LogConfig

func initLogger() {
	// 日志文件配置
	hook := &lumberjack.Logger{
		Filename:   LogConf.Filename,
		MaxSize:    LogConf.MaxSize,
		MaxBackups: LogConf.MaxBackups,
		MaxAge:     LogConf.MaxAge,
		Compress:   LogConf.Compress,
	}

	// 设置日志级别
	var level zapcore.Level
	switch LogConf.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(hook)),
		atomicLevel,
	)

	caller := zap.AddCaller()
	development := zap.Development()
	filed := zap.Fields(zap.String("serviceName", "MyGin"))
	global.Logger = zap.New(core, caller, development, filed)
}
