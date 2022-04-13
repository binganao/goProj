package api

import (
	"mall/models"
	"mall/response"
	"mall/service"

	"github.com/gin-gonic/gin"
)

var webCategory service.WebCategoryService

func WebCreateCategory(c *gin.Context) {
	var param models.WebCategoryCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	if count := webCategory.Create(param); count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

func WebDeleteCategory(c *gin.Context) {
	var param models.WebCategoryDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	if count := webCategory.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

func WebUpdateCategory(c *gin.Context) {
	var param models.WebCategoryUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	if count := webCategory.Update(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
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
