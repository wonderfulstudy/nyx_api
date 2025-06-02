package validservice

import (
	"github.com/astaxie/beego/validation"
	"github.com/google/uuid"
)

func init() {
	validation.AddCustomFunc("validuuid", ValidUUID)
}

// 自定义验证UUID函数
func ValidUUID(v *validation.Validation, obj interface{}, key string) {
	str, ok := obj.(string)
	if !ok {
		v.SetError(key, "必须是字符串")
		return
	}

	if _, err := uuid.Parse(str); err != nil {
		v.SetError(key, "不是有效的 UUID")
	}
}
