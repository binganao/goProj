package service

import (
	"dServer/settings"

	"dServer/middleware"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(middleware.LogMiddleware())
	r.Use(gin.Recovery())

	if settings.Debug {
		pprof.Register(r)
	}
	//r.GET("/favicon.ico", GetFavicon)

	base := r.Group(settings.Path)
	base.GET("/", ParseGet)
	base.POST("/", ParsePost)
	base.PUT("/", ParsePut)

	return r
}
