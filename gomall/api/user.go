package api

import (
	"mall/common"
	"mall/models"
	"mall/response"
	"mall/service"

	"github.com/gin-gonic/gin"
)

var user service.WebUserService

func WebGetCaptcha(c *gin.Context) {
	id, b64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{
		"captchaId":  id,
		"captchaImg": b64s,
	}
	response.Success("操作成功", data, c)
}

func WebUserLogin(c *gin.Context){
	var param models.WebUserLoginParam
	if err:=c.ShouldBind(&param);err!=nil{
		response.Failed("参数无效", c)
		return
	}

	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue){
		response.Failed("验证码错误", c)
		return
	}

	uid:=user.Login(param)
	if uid>0{
		token, _:=common.
	}
}