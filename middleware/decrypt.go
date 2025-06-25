package middleware

import (
	"bytes"
	"io"
	"net/http"
	"nyx_api/pkg/aes"
	"nyx_api/pkg/app"
	"nyx_api/pkg/e"

	"github.com/gin-gonic/gin"
)

func DecryptMiddleware() gin.HandlerFunc {
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

		// 设置需要解密的接口列表
		list := map[string]bool{
			"/api/v1/user/login": true,
		}
		// 获取请求路径
		path := c.Request.URL.Path
		if _, ok := list[path]; ok {
			// log.Log.Debugf("进入解密流程：%s", path)
			// 读取原始请求体
			appG := app.Gin{C: c}
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				appG.Response(http.StatusBadRequest, e.ERROR, "Failed to read request body")
				return
			}
			defer c.Request.Body.Close()

			// log.Log.Debugf("解密前数据：%s", string(body))
			decryptedData, err := aes.AesDecryptCBCBase64(string(body))
			if err != nil {
				appG.Response(http.StatusBadRequest, e.ERROR, "Failed to decrypt data")
				return
			}
			// log.Log.Debugf("解密后数据：%s", decryptedData)
			// 重新设置请求体
			c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(decryptedData)))

			// 继续执行后续处理
			c.Next()
			return
		}
		c.Next()
		return
	}
}
