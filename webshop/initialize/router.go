package initialize

import (
	"log"
	"mall/api"
	"mall/global"
	"mall/middleware"

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

		web.Use(middleware.JwtAuth())

		web.POST("/upload", api.WebFileUpload)

		web.POST("/category/create", api.WebCreateCategory)
		web.DELETE("/category/delete", api.WebDeleteCategory)
		web.PUT("/category/update", api.WebUpdateCategory)
		web.GET("/category/list", api.WebGetCategoryList)
		web.GET("/category/option", api.WebGetCategoryOption)

		web.POST("/product/create", api.WebCreateProduct)
		web.DELETE("/product/delete", api.WebDeleteProduct)
		web.PUT("/product/update", api.WebUpdateProduct)
		web.PUT("/product/status/update", api.WebUpdateProductStatus)
		web.GET("/product/info", api.WebGetProductInfo)
		web.GET("/product/list", api.WebGetProductList)
	}

	//...

	port := ":" + global.Config.Server.Port
	if err := engine.Run(port); err != nil {
		log.Panicln("server start", err)
	}
}
