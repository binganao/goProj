package api

import (
	"mall/response"

	"github.com/gin-gonic/gin"
)

func act(param interface{}, do func() int64, name string, c *gin.Context) {
	if err := c.ShouldBind(param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := do(); count > 0 {
		response.Success(name+"成功", count, c)
		return
	}
	response.Failed(name+"失败", c)
}
