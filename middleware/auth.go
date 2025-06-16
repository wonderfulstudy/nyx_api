package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 白名单路径：无需验证 Authorization 的接口
		whitelist := map[string]bool{
			"/user/login": true,
			"/auth":       true,
		}

		// 获取请求路径
		path := c.Request.URL.Path

		// 如果在白名单中，直接跳过验证
		if _, ok := whitelist[path]; ok {
			c.Next()
			return
		}

		// 验证 Authorization 头
		authHeader := c.GetHeader("Authorization")
		if strings.TrimSpace(authHeader) == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		// TODO：可在此校验 token 合法性

		// 继续执行后续 handler
		c.Next()
	}
}
