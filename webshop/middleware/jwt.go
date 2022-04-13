package middleware

import (
	"mall/common"
	"mall/response"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Failed("未登录或非法访问", c)
			c.Abort()
			return
		}
		if err := common.VerifyToken(token); err != nil {
			response.Failed("登录已过期", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
