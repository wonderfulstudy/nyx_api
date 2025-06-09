package routers

import (
	"bytes"
	"fmt"
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
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, x-token")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, Location")
		// 允许携带凭证（如 Cookie）
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

// 解密中间件
func decryptMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过不需要解密的请求（如 OPTIONS 或非 POST/PUT 请求）
		fmt.Println("请求url", c.Request.URL.Path)
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
		fmt.Println("解密数据：", decryptedData)
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
		apiv1.POST("/user/update", v1.UpdateUser)
		apiv1.POST("/user/create", v1.CreateUser)
		apiv1.POST("/user/delete", v1.DeleteUser)
		apiv1.POST("/user/login", v1.UserLogin)
		apiv1.GET("/user/info", v1.UserInfo)
		apiv1.GET("/user/list", v1.UserList)

		apiv1.GET("/worker/list", v1.GetWorkerList)
	}

	return r
}
