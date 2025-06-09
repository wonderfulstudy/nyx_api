package v1

import (
	"nyx_api/models"
	"nyx_api/pkg/e"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetWorkerList(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	valid := validation.Validation{}
	valid.Required(page, "page").Message("page 参数不能为空")
	valid.Required(limit, "limit").Message("limit 参数不能为空")
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	workers := models.GetWorkerList(pageInt, limitInt)
	c.JSON(e.SUCCESS, gin.H{
		"code": 20000,
		"data": gin.H{
			"total": models.GetWorkerCount(),
			"items": workers,
		},
	})
}
