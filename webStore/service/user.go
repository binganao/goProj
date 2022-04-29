package service

import (
	"mall/global"
	"mall/models"
)

type WebUserService struct{}

func (u *WebUserService) Login(param models.WebUserLoginParam) uint64 {
	var user models.User
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&user)
	return user.Id
}
