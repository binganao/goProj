package api

import (
	"mall/models"
	"mall/response"
	"mall/service"

	"github.com/gin-gonic/gin"
)

var webOrder service.WebOrderService

func WebCreateOrder(c *gin.Context) {
	var param *models.WebOrderCreateParam
	act(param, func() int64 {
		return webOrder.Create(*param)
	}, "创建", c)
}

func WebDeleteOrder(c *gin.Context) {
	var param *models.WebOrderDeleteParam
	act(param, func() int64 {
		return webOrder.Delete(*param)
	}, "删除", c)
}

func WebUpdateOrder(c *gin.Context) {
	var param *models.WebOrderUpdateParam
	act(param, func() int64 {
		return webOrder.Update(*param)
	}, "更新", c)
}

func WebGetOrderInfo(c *gin.Context) {
	var param models.WebOrderInfoParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productInfo := webOrder.GetInfo(param)
	response.Success("查询成功", productInfo, c)
}

func WebGetOrderList(c *gin.Context) {
	var param models.WebOrderListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productList, rows := webOrder.GetList(param)
	response.SuccessPage("查询成功", productList, rows, c)
}
