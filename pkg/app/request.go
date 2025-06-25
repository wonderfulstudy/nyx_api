package app

import (
	"nyx_api/pkg/log"

	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Log.Info(err.Key, err.Message)
	}

	return
}
