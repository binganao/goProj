package initialize

import (
	"./global"
	"github.com/gin-gonic/gin"
)

func Router() {
	engine := gin.Default()

	//engine.Use()

	engine.Static("/image", global.Config.Upload.SavePath)

	web := engine.Group("/web")
	{
		web.GET("/captcha", api)
	}

}
