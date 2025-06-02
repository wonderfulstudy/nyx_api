package v1

import (
	"net/http"
	"nyx_api/models"
	"nyx_api/pkg/e"
	"regexp"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetUserBy(c *gin.Context) {
	uuid := c.Query("uuid")
	name := c.Query("name")
	phone := c.Query("phone")

	// 验证uuid合法性
	valid := validation.Validation{}
	code := e.INVALID_PARAMS
	if uuid == "" && name == "" && phone == "" {
		valid.SetError("params", "必须提供uuid/name/phone中的至少一个非空参数")
	} else if uuid != "" && name == "" && phone == "" {
		// 添加UUID格式验证
		uuidPattern := "^[A-Za-z0-9]{8}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{12}$"
		matched, _ := regexp.MatchString(uuidPattern, uuid)
		if !matched {
			valid.SetError("uuid", "uuid格式不正确")
			code = e.INVALID_PARAMS
		} else {
			code = e.SUCCESS
			var user models.User
			user = models.GetUserByUuid(uuid)
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": user,
			})
		}
	} else if uuid == "" && name != "" && phone == "" {
		code = e.SUCCESS
	} else if uuid == "" && name == "" && phone != "" {
		code = e.SUCCESS
	} else {
		valid.SetError("params", "参数错误")
		code = e.INVALID_PARAMS
	}
}

// 新增用户
func AddUser(c *gin.Context) {
}

// 修改用户信息
func UpdateUser(c *gin.Context) {
}

// 删除用户
func DeleteUser(c *gin.Context) {
}
