package user

import "github.com/gin-gonic/gin"

func SetupUserRoutes(r *gin.Engine) {
	r.GET("/user/:name", func(c *gin.Context) {
		// 处理逻辑
	})
}
