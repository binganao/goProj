package api

import (
	"mall/models"
	"mall/response"
	"mall/service"

	"github.com/gin-gonic/gin"
)

var webProduct service.WebProductService

func WebCreateProduct(c *gin.Context) {
	var param *models.WebProductCreateParam
	act(param, func() int64 {
		return webProduct.Create(*param)
	}, "创建", c)
}

func WebDeleteProduct(c *gin.Context) {
	var param *models.WebProductDeleteParam
	act(param, func() int64 {
		return webProduct.Delete(*param)
	}, "删除", c)
}

func WebUpdateProduct(c *gin.Context) {
	var param *models.WebProductUpdateParam
	act(param, func() int64 {
		return webProduct.Update(*param)
	}, "更新", c)
}

func WebUpdateProductStatus(c *gin.Context) {
	var param *models.WebProductStatusUpdateParam
	act(param, func() int64 {
		return webProduct.UpdateStatus(*param)
	}, "更新", c)
}

// WebGetProductInfo 后台管理前端，获取商品信息
func WebGetProductInfo(c *gin.Context) {
	var param models.WebProductInfoParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productInfo := webProduct.GetInfo(param)
	response.Success("查询成功", productInfo, c)
}

// WebGetProductList 后台管理前端，获取商品列表
func WebGetProductList(c *gin.Context) {
	var param models.WebProductListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productList, rows := webProduct.GetList(param)
	response.SuccessPage("查询成功", productList, rows, c)
}
