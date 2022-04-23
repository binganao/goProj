package models

//UNCHANGED
type User struct {
	Id           uint64 `gorm:"primaryKey"`
	Username     string `gorm:"username"`
	Password     string `gorm:"password"`
	Status       uint   `gorm:"status"`
	CaptchaId    string `gorm:"captchaId"`
	CaptchaValue string `gorm:"captchaValue"`
}

type WebUserLoginParam struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaId    string `json:"captchaId"`
	CaptchaValue string `json:"captchaValue"`
}

type WebUserInfo struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}
