package api

import (
	"mall/models"
	"mall/response"
	"mall/service"

	"github.com/gin-gonic/gin"
)

var address service.AddressService

func WebAddAddress(c *gin.Context) {
	var param *models.WebAddressAddParam
	act(param, func() int64 {
		return address.Add(*param)
	}, "添加", c)
}

func WebDeleteAddress(c *gin.Context) {
	var key *models.WebAddressDeleteParam
	act(key, func() int64 {
		return address.Delete(key.AddressId)
	}, "删除", c)
}

func WebUpdateAddress(c *gin.Context) {
	var param *models.WebAddressUpdateParam
	act(param, func() int64 {
		return address.Update(*param)
	}, "更新", c)
}

func WebGetAddressUpdateInfo(c *gin.Context) {
	var key models.WebAddressInfoParam
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	info := address.GetInfo(key.AddressId)
	response.Success("查询成功", info, c)
}

func WebGetAddressList(c *gin.Context) {
	var param models.WebAddressListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	addressList := address.GetList(param.UserId)
	response.Success("查询成功", addressList, c)
}
