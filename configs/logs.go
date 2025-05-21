package configs

import (
	"os"

	"github.com/sirupsen/logrus"
)

// 全局 Logger
var Logger = logrus.New()

func InitLogger() {
	// 设置日志输出格式
	Logger.SetFormatter(&logrus.JSONFormatter{})
	// 设置日志输出到标准输出（也可以是文件）
	Logger.SetOutput(os.Stdout)
	// 设置日志级别（Debug, Info, Warn, Error）
	Logger.SetLevel(logrus.DebugLevel)
}
