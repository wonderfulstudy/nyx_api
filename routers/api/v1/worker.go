package v1

import (
	"net/http"
	"nyx_api/pkg/app"
	"nyx_api/pkg/e"

	worker_service "nyx_api/service/worker_service"

	"github.com/gin-gonic/gin"
)

// @Summary 获取矿机列表
// @Description 获取所有矿机数据
// @Tags 机器管理
// @Produce  json
// @Param page query string true "分页查询第几页"
// @Param limit query string true "每页限制查询的数据数量"
// @Success 200 {object} app.Response	"success"
// @Failure 500 {object} app.Response	"获取矿机列表失败"
// @Router /api/v1/user/list [get]
func WorkerListHandler(c *gin.Context) {
	var req worker_service.WorkerListRequest
	appG := app.Gin{C: c}
	httpCode, errCode, err := app.BindQueryAndValid(appG.C, &req)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
		return
	}

	data, err := worker_service.ListService(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
