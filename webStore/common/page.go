package common

import (
	"mall/global"
	"mall/models"
)

func RestPage(page models.Page, name string, query interface{}, dest interface{}, bind interface{}) int64 {
	offset := (page.PageNum - 1) * page.PageSize
	global.Db.Offset(offset).Limit(page.PageSize).Table(name).Where(query).Find(dest)
	return global.Db.Table(name).Where(query).Find(bind).RowsAffected
}
