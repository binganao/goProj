package initialize

import (
	"./global"
	"github.com/go-redis/redis/v8"
)

func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	global.Rdb = rdb
}
