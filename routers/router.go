package routers

import (
	_ "nyx_api/docs"
	"nyx_api/middleware"
	"nyx_api/pkg/setting"
	v1 "nyx_api/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.DecryptMiddleware())
	r.Use(middleware.AuthMiddleware())

	gin.SetMode(setting.RunMode)

	// r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		// 获取用户信息
		apiv1.GET("/user/info", v1.UserInfoHandler)
		apiv1.GET("/user/list", v1.UserListHandler)
		apiv1.POST("/user/login", v1.UserLoginHandler)
		apiv1.POST("/user/update", v1.UserUpdateHandler)
		apiv1.POST("/user/create", v1.UserCreateHandler)
		apiv1.POST("/user/delete", v1.UserDeleteHandler)
		apiv1.POST("/user/logout", v1.UserLogoutHandler)

		apiv1.GET("/worker/list", v1.WorkerListHandler)

		apiv1.GET("/wallet/info", v1.WalletInfoHandler)
		apiv1.GET("/wallet/action", v1.WalletActionListHandler)
	}

	return r
}
