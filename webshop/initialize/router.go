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

		web.POST("/cart/add", api.WebAddCart)
		web.DELETE("/cart/delete", api.WebDeleteCart)
		web.DELETE("/cart/clear", api.WebClearCart)
		web.GET("/cart/info", api.WebGetCartInfo)

		web.DELETE("/order/delete", api.WebDeleteOrder)
		web.PUT("/order/update", api.WebUpdateOrder)
		web.GET("/order/list", api.WebGetOrderList)
		web.GET("/order/detail", api.WebGetOrderInfo)

		web.POST("/address/add", api.WebAddAddress)
		web.DELETE("/address/delete", api.WebDeleteAddress)
		web.PUT("/address/update", api.WebUpdateAddress)
		web.GET("/address/info", api.WebGetAddressUpdateInfo)
		web.GET("/address/list", api.WebGetAddressList)
	}

	port := ":" + global.Config.Server.Port
	if err := engine.Run(port); err != nil {
		log.Panicln("server start", err)
	}
}
