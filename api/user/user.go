package user

import (
	"database/sql"
	"nyx_api/configs"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/v1/user")
	userGroup.POST("/login", handleLogin)
}