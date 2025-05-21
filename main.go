package main

import (
	"nyx_api/api/pubkey"
	"nyx_api/api/user"
	"nyx_api/configs"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func main() {
	// 初始化数据库
	configs.InitDB()
	configs.InitLogger()

	r := gin.Default()
	r.Use(CorsMiddleware())

	user.SetupUserRoutes(r)
	pubkey.SetupPubkeyRoutes(r)

	r.Run(":9528")
}
