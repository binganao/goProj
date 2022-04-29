package common

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore
var driver = base64Captcha.NewDriverDigit(40, 120, 4, .7, 80)

func GenerateCaptcha() (id, b64s string, err error) {
	return base64Captcha.NewCaptcha(driver, store).Generate()
}

func VerifyCaptcha(id, value string) bool {
	return store.Verify(id, value, true)
}
