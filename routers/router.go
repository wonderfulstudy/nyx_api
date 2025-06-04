package routers

import (
	"bytes"
	"io"
	"nyx_api/middleware/aes"
	"nyx_api/pkg/setting"
	api "nyx_api/routers/api"
	v1 "nyx_api/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// 跨域中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		// c.Next()
	}
}

// 解密中间件
func decryptMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过不需要解密的请求（如 OPTIONS 或非 POST/PUT 请求）
		if c.Request.Method == "OPTIONS" || c.Request.Method != "POST" {
			c.Next()
			return
		}

		// 跳过非 JSON 请求
		if c.ContentType() != "application/json" {
			c.Next()
			return
		}

		// 读取原始请求体
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to read request body"})
			return
		}
		defer c.Request.Body.Close()

		decryptedData, err := aes.AesDecryptCBCBase64(string(body))
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to decrypt data", "details": err.Error()})
			return
		}

		// 重新设置请求体
		c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(decryptedData)))

		// 继续执行后续处理
		c.Next()
	}
}
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(corsMiddleware())
	r.Use(decryptMiddleware())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	{
		// 获取用户信息
		apiv1.GET("/user", v1.GetUserBy)
		apiv1.POST("/user/login", v1.UserLogin)
	}

	return r
}
