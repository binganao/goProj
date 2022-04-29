package global

import (
	"mall/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
)
