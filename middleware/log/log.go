package log

import (
	"nyx_api/pkg/setting"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		// 在 Gin 上下文中设置 Logrus 日志实例
		Log.WithFields(map[string]interface{}{
			"status":  statusCode,
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"latency": duration,
			"client":  c.ClientIP(),
		}).Info("HTTP Request")

		// 继续处理请求
		c.Next()
	}
}
