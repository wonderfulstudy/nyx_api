package app

import (
	"errors"
	"net/http"
	"nyx_api/pkg/e"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func validateForm(c *gin.Context, form interface{}) error {
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return errors.New("执行验证表单数据函数失败：" + err.Error())
	}
	if !check {
		MarkErrors(valid.Errors)
		return errors.New("表单数据验证失败")
	}

	return nil
}

func BindJsonAndValid(c *gin.Context, form interface{}) (int, int, error) {
	err := c.ShouldBind(form)
	if err != nil {
		return http.StatusBadRequest, e.ERROR_BIND, err
	}

	if err := validateForm(c, form); err != nil {
		return http.StatusInternalServerError, e.INVALID_PARAMS, err
	}

	return http.StatusOK, e.SUCCESS, nil
}

func BindUrlAndValid(c *gin.Context, form interface{}) (int, int, error) {
	err := c.ShouldBindUri(form)
	if err != nil {
		return http.StatusBadRequest, e.ERROR_BIND, err
	}

	if err := validateForm(c, form); err != nil {
		return http.StatusInternalServerError, e.INVALID_PARAMS, err
	}

	return http.StatusOK, e.SUCCESS, nil
}

func BindQueryAndValid(c *gin.Context, form interface{}) (int, int, error) {
	err := c.ShouldBindQuery(form)
	if err != nil {
		return http.StatusBadRequest, e.ERROR_BIND, err
	}

	if err := validateForm(c, form); err != nil {
		return http.StatusInternalServerError, e.INVALID_PARAMS, err
	}

	return http.StatusOK, e.SUCCESS, nil
}
