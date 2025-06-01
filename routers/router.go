package routers

import (
	"nyx_api/pkg/setting"
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

	apiv1 := r.Group("/api/v1")
	{
		// 获取用户信息
		apiv1.GET("/user", v1.GetUser)
		// 新增用户
		apiv1.POST("/user", v1.AddUser)

		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
