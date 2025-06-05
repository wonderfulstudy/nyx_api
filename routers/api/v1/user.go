package v1

import (
	"fmt"
	"net/http"
	"nyx_api/middleware/aes"
	"nyx_api/models"
	"nyx_api/pkg/e"
	"regexp"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 数据模型定义（统一放在文件顶部）
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	// 其他需要的字段...
}

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

			fmt.Println("--------加密开始")
			cipherText, err := aes.AesEncryptCBCBase64(user.Password) // 使用完整参数调用
			if err != nil {
				fmt.Println("--------加密失败")
			}
			user.Password = string(cipherText) // 赋值base64编码后的字符串
			fmt.Println("--------加密结束")

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

// 用户登录
func UserLogin(c *gin.Context) {
	// 定义接收结构体
	var loginReq LoginRequest
	// 绑定JSON数据
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	// 查询数据库
	var user models.User
	user = models.GetUserByName(loginReq.Username)

	// 验证密码
	if user.Password != loginReq.Password {
		// 添加调试日志
		fmt.Printf("数据库密码: [%s]", user.Password)
		fmt.Printf("输入密码: [%s]\n", loginReq.Password)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": e.ERROR_AUTH,
			"msg":  e.GetMsg(e.ERROR_AUTH),
		})
		return
	}

	c.JSON(e.SUCCESS, gin.H{
		"code": 20000,
		"data": gin.H{
			"token": user.Token,
		},
	})
}

// 用户信息
func UserInfo(c *gin.Context) {
	token := c.Query("token")
	fmt.Println("token: ", token)
	info := models.GetUserByToken(token)
	var roles []string
	roles = append(roles, models.GetRoleById(info.RoleId).KeyName)
	// TODO: 添加基于token的业务逻辑处理
	res := gin.H{
		"avatar":       info.Avatar,
		"name":         info.Name,
		"introduction": info.Introduction,
		"roles":        roles,
	}
	fmt.Println("返回数据", res)

	// 返回JSON响应
	c.JSON(e.SUCCESS, gin.H{
		"code": 20000,
		"data": res,
	})
}
