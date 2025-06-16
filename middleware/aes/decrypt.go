package aes

import (
	"bytes"
	"io"
	"nyx_api/middleware/log"

	"github.com/gin-gonic/gin"
)

func DecryptMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过不需要解密的请求（如 OPTIONS 或非 POST/PUT 请求）
		log.Log.Info("请求url", c.Request.URL.Path)
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

		decryptedData, err := AesDecryptCBCBase64(string(body))
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to decrypt data", "details": err.Error()})
			return
		}
		log.Log.Info("解密数据：", decryptedData)
		// 重新设置请求体
		c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(decryptedData)))

		// 继续执行后续处理
		c.Next()
	}
}
