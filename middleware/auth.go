package middleware

import (
	"net/http"
	"nyx_api/pkg/app"
	"nyx_api/pkg/e"
	"nyx_api/pkg/log"
	"nyx_api/pkg/redis"
	"nyx_api/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 白名单路径：无需验证 Authorization 的接口
		whitelist := map[string]bool{
			"/api/v1/user/login": true,
		}

		// 获取请求路径
		path := c.Request.URL.Path

		// 如果在白名单中，直接跳过验证
		if _, ok := whitelist[path]; ok {
			c.Next()
			return
		}

		appG := app.Gin{C: c}
		// 验证 Authorization 头
		authHeader := c.GetHeader("Authorization")
		if strings.TrimSpace(authHeader) == "" {
			appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_TOKEN, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		claims, err := util.ParseToken(authHeader)
		if err != nil {
			log.Log.Errorf("Token parse error: %v", err)
			appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_TOKEN, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		redisToken, err := redis.RDB.Get(redis.CTX, "user:"+claims.Username).Result()
		if err != nil {
			appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_TOKEN, gin.H{
				"error": "Failed to retrieve token from Redis",
			})
			c.Abort()
			return
		}

		if redisToken != authHeader {
			appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_TOKEN, gin.H{
				"error": "Token mismatch",
			})
			c.Abort()
			return
		}

		c.Set("uuid", claims.Uuid)
		c.Set("username", claims.Username)
		c.Set("phone", claims.Phone)

		c.Next()
	}
}
