package routers

import (
	"nyx_api/pkg/setting"
	api "nyx_api/routers/api"
	v1 "nyx_api/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 			return
// 		}
// 		c.Next()
// 	}
// }

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	{
		// 获取用户信息
		apiv1.GET("/user", v1.GetUserBy)
	}

	return r
}
