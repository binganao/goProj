package api

import (
	"mall/models"
	"mall/response"
	"mall/service"

	"github.com/gin-gonic/gin"
)

var cart service.CartService

func WebAddCart(c *gin.Context) {
	var param *models.WebCartAddParam
	act(param, func() int64 {
		return cart.Add(*param)
	}, "添加", c)
}

func WebDeleteCart(c *gin.Context) {
	var param models.WebCartDeleteParam
	act(param, func() int64 {
		return cart.Delete(param)
	}, "删除", c)
}

func WebClearCart(c *gin.Context) {
	var param models.WebCartClearParam
	act(param, func() int64 {
		return cart.Clear(param)
	}, "清除", c)
}

func WebGetCartInfo(c *gin.Context) {
	var param models.WebCartQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	info := cart.GetInfo(param)
	if len(info.CartItem) == 0 {
		response.Success("购物车是空的", info, c)
	}
	response.Success("查询成功", info, c)
}
