package initialize

import (
	"mall/global"

	"github.com/go-redis/redis/v8"
)

func Redis() {
	RetryModule(ConnectRedis, 0)
}

func ConnectRedis() {
	m := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     m.Addr,
		Password: "",
		DB:       0,
	})
	global.Rdb = rdb
}
