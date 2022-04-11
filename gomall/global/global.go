package global

import (
	"mall/config"

	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
	Es     *elastic.Client
)
