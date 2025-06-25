package redis

import (
	"context"
	"nyx_api/pkg/log"
	"nyx_api/pkg/setting"
	"time"

	"github.com/go-redis/redis/v8"
)

// 声明一个全局的RDB变量
var (
	RDB *redis.Client
	CTX = context.Background()
)

// 初始化连接
func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     setting.RedisHost,
		Password: setting.RedisPassword,
		DB:       setting.RedisDatabase,
		PoolSize: 100,
	})

	// 需要使用context库
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Log.Debugf("redis数据库连接失败: %s", err)
	}
}
