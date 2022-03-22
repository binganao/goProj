package main

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", ParseGet)
	r.POST("/", ParsePost)
	r.PUT("/", ParsePut)
	r.GET("/favicon.ico", GetFavicon)

	return r
}
