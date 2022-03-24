package main

import (
	"dServer/settings"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.GET("/favicon.ico", GetFavicon)

	base := r.Group(settings.Path)
	base.GET("/", ParseGet)
	base.POST("/", ParsePost)
	base.PUT("/", ParsePut)

	return r
}
