package v1

import (
	"net/http"
	"nyx_api/pkg/app"
	"nyx_api/pkg/e"
	"nyx_api/pkg/log"
	wallet_service "nyx_api/service/wallet_service"

	"github.com/gin-gonic/gin"
)

// @Summary 用户钱包数据
// @Description 根据用户uuid获取用户的钱包数据
// @Tags 钱包管理
// @Produce  json
// @Param uuid body string true "用户uuid"
// @Success 200 {object} app.Response	"success"
// @Success 500 {object} app.Response	"用户钱包数据失败"
// @Router /api/v1/wallet/info [get]
func WalletInfoHandler(c *gin.Context) {
	var req wallet_service.WalletRequest
	appG := app.Gin{C: c}
	httpCode, errCode, err := app.BindQueryAndValid(c, &req)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
		return
	}

	log.Log.Debugf("wallet 请求参数 : %+v", req)

	wallet, err := wallet_service.WalletInfoService(req.Uuid)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, wallet)
}

// @Summary 交易记录
// @Description 根据用户uuid获取该用户的所有交易记录
// @Tags 钱包管理
// @Produce  json
// @Param uuid body string true "用户uuid"
// @Success 200 {object} app.Response	"success"
// @Success 500 {object} app.Response	"获取用户交易记录失败"
// @Router /api/v1/wallet/action [post]
func WalletActionListHandler(c *gin.Context) {
	var req wallet_service.WalletRequest
	appG := app.Gin{C: c}
	httpCode, errCode, err := app.BindJsonAndValid(c, &req)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, err.Error())
	}

	data, err := wallet_service.WalletActionListService(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
