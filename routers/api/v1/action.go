package v1

import (
	"net/http"
	"nyx_api/pkg/app"
	"nyx_api/pkg/e"
	wallet_service "nyx_api/service/wallet_service"

	"github.com/gin-gonic/gin"
)

// @Summary 交易记录
// @Description 根据用户uuid获取该用户的所有交易记录
// @Tags 交易管理
// @Produce  json
// @Param uuid body string true "用户uuid"
// @Success 200 {object} app.Response	"success"
// @Success 500 {object} app.Response	"获取用户交易记录失败"
// @Router /api/v1/action/list [get]
func ActionListHandler(c *gin.Context) {
	var req wallet_service.WalletRequest
	appG := app.Gin{C: c}
	httpCode, errCode, err := app.BindJsonAndValid(c, &req)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
	}

	action, err := wallet_service.WalletActionListService(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, action)
}
