package middleware

import (
	"nyx_api/pkg/log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		// 在 Gin 上下文中设置 Logrus 日志实例
		log.Log.WithFields(map[string]interface{}{
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
