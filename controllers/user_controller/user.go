package user

import (
	"net/http"
	"nyx_api/middleware/log"
	"nyx_api/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
		})
	}

	validate := validator.New()
	err := validate.Struct(&loginRequest)
	if err != nil {
		log.Log.Info("=== error msg ====")
		log.Log.Info(err)

		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Log.Info(err)
			return
		}

		log.Log.Info("\r\n=========== error field info ====================")
		for _, err := range err.(validator.ValidationErrors) {
			// 列出效验出错字段的信息
			log.Log.Info("Namespace: ", err.Namespace())
			log.Log.Info("Fild: ", err.Field())
			log.Log.Info("StructNamespace: ", err.StructNamespace())
			log.Log.Info("StructField: ", err.StructField())
			log.Log.Info("Tag: ", err.Tag())
			log.Log.Info("ActualTag: ", err.ActualTag())
			log.Log.Info("Kind: ", err.Kind())
			log.Log.Info("Type: ", err.Type())
			log.Log.Info("Value: ", err.Value())
			log.Log.Info("Param: ", err.Param())
			log.Log.Info()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
}
