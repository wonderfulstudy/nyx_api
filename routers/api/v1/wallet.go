package v1

import (
	"fmt"
	"net/http"
	"nyx_api/models"
	"nyx_api/pkg/e"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type WalletRequest struct {
	Uuid string `json:"uuid" valid:"Required"`
}

func GetWallet(c *gin.Context) {
	uuid := c.Query("uuid")

	valid := validation.Validation{}
	code := e.INVALID_PARAMS
	if uuid == "" {
		valid.SetError("uuid", "uuid不能为空")
	} else if uuid != "" {
		code = e.SUCCESS
		wallet := models.GetWalletByUuid(uuid)
		c.JSON(e.SUCCESS, gin.H{
			"code": 20000,
			"msg":  e.GetMsg(code),
			"data": wallet,
		})
	} else {
		valid.SetError("params", "参数错误")
		code = e.INVALID_PARAMS
	}
}

func GetAction(c *gin.Context) {
	uuid := c.Query("uuid")

	fmt.Println("uuid", uuid)

	valid := validation.Validation{}
	code := e.INVALID_PARAMS
	if uuid == "" {
		valid.SetError("uuid", "uuid不能为空")
	} else if uuid != "" {
		code = e.SUCCESS
		walletAction := models.GetActionByUuid(uuid)
		fmt.Println(walletAction[0].UpdatedAt)
		c.JSON(e.SUCCESS, gin.H{
			"code": 20000,
			"msg":  e.GetMsg(code),
			"data": walletAction,
		})
		return
	} else {
		valid.SetError("params", "参数错误")
		code = e.INVALID_PARAMS
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
