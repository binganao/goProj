package initialize

import (
	"fmt"
	"time"

	"mall/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Mysql() {
	RetryModule(ConnectMysql, 0)
}

func ConnectMysql() {
	m := global.Config.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	global.Db = db
}
