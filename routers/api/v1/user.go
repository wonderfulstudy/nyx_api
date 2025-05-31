package v1

import (
	"nyx_api/models"

	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetUser(c *gin.Context) {
	db.First(&models.User{})
}

// 新增用户
func AddUser(c *gin.Context) {
	user := models.User{Name: "test", Phone: "12345678901", Uid: "12345678901"}
	result := models.DB.Create(&user)

	result.Error
	result.RowsAffected
}

// 修改用户信息
func UpdateUser(c *gin.Context) {
}

// 删除用户
func DeleteUser(c *gin.Context) {
}
