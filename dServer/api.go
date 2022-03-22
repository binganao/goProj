package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	STATUS_OK = 200
)

func Ok(c *gin.Context) {
	c.JSON(STATUS_OK, gin.H{
		"test": "ok",
	})
}

func GetHistory(c *gin.Context) {
	c.HTML(STATUS_OK, strings.Join(History, ""))
}
