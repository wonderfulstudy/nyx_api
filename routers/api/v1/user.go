package v1

import (
	"net/http"
	"nyx_api/pkg/app"
	"nyx_api/pkg/e"
	"nyx_api/pkg/log"
	user_service "nyx_api/service/user_service"

	"github.com/gin-gonic/gin"
)

// @Summary 创建用户
// @Description 创建一个用户，并附带默认配置，默认密码为dephy@2025.com
// @Tags 用户管理
// @Produce  json
// @Param username body string true "用户名"
// @Param phone body string true "用户电话号码"
// @Success 200 {object} app.Response	"success"
// @Success 500 {object} app.Response	"创建用户失败"
// @Router /api/v1/user/create [post]
func UserCreateHandler(c *gin.Context) {
	var req user_service.CreateRequest
	appG := app.Gin{C: c}
	httpCode, errCode, err := app.BindJsonAndValid(c, &req)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
		return
	}

	if err := user_service.CreateUserService(req); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, e.GetMsg(e.SUCCESS))
}

// @Summary 更新用户
// @Description 更新一个用户的所有信息，根据uuid搜索用户
// @Tags 用户管理
// @Produce  json
// @Param uuid body string true "用户uuid"
// @Param name body string true "修改后的用户真实姓名"
// @Param avatar body string true "用户头像url"
// @Param introduction body string true "用户个人简介描述信息"
// @Param phone body string true "用户电话号码"
// @Param address body string true "用户提币地址"
// @Success 200 {object} app.Response	"success"
// @Failure 404 {object} app.Response	"用户未找到"
// @Failure 500 {object} app.Response	"更新用户信息失败"
// @Router /api/v1/user/update [post]
func UserUpdateHandler(c *gin.Context) {
	var req user_service.UpdateRequest
	appG := app.Gin{C: c}
	httpCode, errCode, err := app.BindJsonAndValid(c, &req)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
		return
	}

	if err := user_service.UpdateUserService(req); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, e.GetMsg(e.SUCCESS))
}

// @Summary 删除用户
// @Description 根据用户uuid参数删除当前用户
// @Tags 用户管理
// @Produce  json
// @Param uuid query string true "用户uuid"
// @Success 200 {object} app.Response	"success"
// @Failure 500 {object} app.Response	"删除用户失败"
// @Router /api/v1/user/delete [get]
func UserDeleteHandler(c *gin.Context) {
	var req user_service.DeleteRequest
	appG := app.Gin{C: c}
	httpCode, errCode, err := app.BindJsonAndValid(c, &req)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
		return
	}

	if err := user_service.DeleteUserService(req); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, e.GetMsg(e.SUCCESS))
}

// @Summary 登录接口
// @Description 用户登录接口
// @Tags 用户管理
// @Produce  json
// @Param username query string true "用户名"
// @Param phone query string true "用户电话号码"
// @Param password query string true "用户密码"
// @Success 200 {object} app.Response	"success"
// @Failure 500 {object} app.Response	"登录失败"
// @Router /api/v1/user/login [post]
func UserLoginHandler(c *gin.Context) {
	var req user_service.LoginRequest
	appG := app.Gin{C: c}
	httpCode, errCode, err := app.BindJsonAndValid(c, &req)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
		return
	}

	token, err := user_service.LoginService(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{"token": token})
}

// @Summary 获取用户列表
// @Description 获取当前数据库中所有用户信息数据
// @Tags 用户管理
// @Produce  json
// @Param page query string true "分页查询第几页"
// @Param limit query string true "每页限制查询的数据数量"
// @Success 200 {object} app.Response	"success"
// @Failure 500 {object} app.Response	"获取用户列表数据失败"
// @Router /api/v1/user/list [get]
func UserListHandler(c *gin.Context) {
	var req user_service.UserListRequest
	appG := app.Gin{C: c}
	httpCode, errCode, err := app.BindQueryAndValid(c, &req)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
		return
	}

	data, err := user_service.ListService(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary 获取用户信息
// @Description 返回用户的详细信息
// @Tags 用户管理
// @Produce  json
// @Success 200 {object} app.Response	"success"
// @Failure 500 {object} app.Response	"获取用户信息失败"
// @Router /api/v1/user/info [get]
func UserInfoHandler(c *gin.Context) {
	appG := app.Gin{C: c}
	uuid, isExist := appG.C.Get("uuid")
	if !isExist {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_VALUE, e.GetMsg(e.ERROR_AUTH_VALUE))
		return
	}

	uuidStr, ok := uuid.(string)
	if !ok {
		appG.Response(http.StatusInternalServerError, e.ERROR, e.GetMsg(e.ERROR))
		return
	}

	response, err := user_service.InfoService(uuidStr)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, response)
}

// @Summary 用户登出
// @Description 当前用户登出操作
// @Tags 用户管理
// @Produce  json
// @Param token query string true "用户token"
// @Success 200 {object} app.Response	"success"
// @Failure 500 {object} app.Response	"用户登出失败"
// @Router /api/v1/user/logout [post]
func UserLogoutHandler(c *gin.Context) {
	usernameC, _ := c.Get("username")
	log.Log.Debugf("上下文username: %v", usernameC)
	appG := app.Gin{C: c}
	uername, _ := appG.C.Get("username")
	log.Log.Debugf("UserLogoutHandler called for user: %v", uername)
	if err := user_service.LogoutService(c); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, e.GetMsg(e.SUCCESS))
}
