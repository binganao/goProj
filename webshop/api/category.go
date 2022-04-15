package api

import (
	"mall/models"
	"mall/response"
	"mall/service"

	"github.com/gin-gonic/gin"
)

var webCategory service.WebCategoryService

func WebCreateCategory(c *gin.Context) {
	var param *models.WebCategoryCreateParam
	act(param, func() int64 {
		return webCategory.Create(*param)
	}, "创建", c)
}

func WebDeleteCategory(c *gin.Context) {
	var param *models.WebCategoryDeleteParam
	act(param, func() int64 {
		return webCategory.Delete(*param)
	}, "删除", c)
}

func WebUpdateCategory(c *gin.Context) {
	var param *models.WebCategoryUpdateParam
	act(param, func() int64 {
		return webCategory.Update(*param)
	}, "更新", c)
}

func WebGetCategoryList(c *gin.Context) {
	var param models.WebCategoryQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	cateList, rows := webCategory.GetList(param)
	response.SuccessPage("查询成功", cateList, rows, c)
}

func WebGetCategoryOption(c *gin.Context) {
	option := webCategory.GetOption()
	response.Success("查询成功", option, c)
}

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
