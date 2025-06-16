package routers

import (
	"nyx_api/middleware"
	"nyx_api/middleware/aes"
	"nyx_api/middleware/log"
	"nyx_api/pkg/setting"
	v1 "nyx_api/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(log.LoggerMiddleware())
	r.Use(gin.Recovery())
	r.Use(middleware.CorsMiddleware())
	r.Use(aes.DecryptMiddleware())

	gin.SetMode(setting.RunMode)

	// r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	{
		// 获取用户信息
		apiv1.POST("/user/update", v1.UpdateUser)
		apiv1.POST("/user/create", v1.CreateUser)
		apiv1.POST("/user/delete", v1.DeleteUser)
		apiv1.POST("/user/login", v1.UserLogin)
		apiv1.GET("/user/info", v1.UserInfo)
		apiv1.GET("/user/list", v1.UserList)
		apiv1.POST("/user/logout", v1.LoginOut)

		apiv1.GET("/worker/list", v1.GetWorkerList)

		apiv1.GET("/wallet/info", v1.GetWallet)
		apiv1.GET("/wallet/action", v1.GetAction)
	}

	return r
}
