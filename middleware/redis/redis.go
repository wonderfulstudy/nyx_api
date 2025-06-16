package redis

import (
	"nyx_api/middleware/log"
	"nyx_api/pkg/setting"

	"github.com/go-redis/redis"
)

// 声明一个全局的RDB变量
var RDB *redis.Client

// 初始化连接
func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     setting.RedisHost,
		Password: setting.RedisPassword,
		DB:       setting.RedisDatabase,
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		log.Log.Debugf("redis数据库连接失败: %s", err)
	}
}
