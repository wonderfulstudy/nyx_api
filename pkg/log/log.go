package log

import (
	"nyx_api/pkg/setting"
	"os"

	"github.com/sirupsen/logrus"
	logurs "github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	// 设置日志输出级别
	Log.SetLevel(logurs.DebugLevel)
	// 配置 Logrus，记录到文件
	file, err := os.OpenFile(setting.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
