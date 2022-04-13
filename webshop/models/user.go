package models

//UNCHANGED
// 数据库，用户数据映射模型
type User struct {
	Id           uint64 `gorm:"primaryKey"`
	Username     string `gorm:"username"`
	Password     string `gorm:"password"`
	Status       uint   `gorm:"status"`
	CaptchaId    string `gorm:"captchaId"`
	CaptchaValue string `gorm:"captchaValue"`
}

// 后台管理前端，用户登录参数模型
type WebUserLoginParam struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaId    string `json:"captchaId"`
	CaptchaValue string `json:"captchaValue"`
}

// 后台管理前端，用户信息传输模型
type WebUserInfo struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}
