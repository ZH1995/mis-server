// 日志
package log

import (
	"mis-server/common/config"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var logToFile *logrus.Logger

// 日志文件名
var loggerFile string

func setLogFile(file string) {
	loggerFile = file
}

// 初始化
func init() {
	setLogFile(filepath.Join(config.Config.Log.Path, config.Config.Log.Name))
}

// 方法调用
func Log() *logrus.Logger {
	if config.Config.Log.Model == "file" {
		return logFile()
	} else {
		if log == nil {
			log = logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}
			log.SetLevel(logrus.DebugLevel)
		}
		return log
	}
}

// 日志方法
func logFile() *logrus.Logger {
	if logToFile == nil {
		logToFile = logrus.New()
		logToFile.SetLevel(logrus.DebugLevel)
		logWriter, _ := rotatelogs.New(
			loggerFile+"_%Y%m%d.log",
			rotatelogs.WithMaxAge(7*24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		writeMap := lfshook.WriterMap{
			logrus.DebugLevel: logWriter,
			logrus.InfoLevel:  logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.FatalLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		logToFile.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}))
	}
	return logToFile
}
