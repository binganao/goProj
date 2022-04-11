package initialize

import (
	"log"
	"mall/api"
	"mall/global"

	"github.com/gin-gonic/gin"
)

func Router() {
	engine := gin.Default()

	//TODO: Cors

	engine.Static("/image", global.Config.Upload.SavePath)

	web := engine.Group("/web")
	{
		web.GET("/captcha", api.WebGetCaptcha)
		web.POST("/login", api.WebUserLogin)
		//TODO: JwtAuth
		web.POST("/upload")
	}

	//...

	port := ":" + global.Config.Server.Port
	if err := engine.Run(port); err != nil {
		log.Panicln("server start", err)
	}
}
