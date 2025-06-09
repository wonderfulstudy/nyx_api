package util

import (
	"github.com/google/uuid"
)

func GenerateStringUUID() string {
	return uuid.New().String()
}

// 生成数字型uuid
func GenerateIntUUID() uint32 {
	u, _ := uuid.NewRandom()
	return u.ID()
}
