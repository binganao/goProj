package service

import (
	"mall/common"
	"mall/global"
	"mall/models"
	"strconv"
	"strings"
)

var cart CartService
var address AddressService

type WebOrderService struct{}

func (o *WebOrderService) Create(param models.WebOrderCreateParam) int64 {
	info := cart.GetInfo(models.WebCartQueryParam{UserId: param.UserId})
	pids := make([]string, 0)
	for _, item := range info.CartItem {
		pids = append(pids, strconv.Itoa(int(item.Id)))
	}
	pidsItem := strings.Join(pids, ",")
	order := models.Order{
		ProductItem: pidsItem,
		TotalPrice:  info.TotalPrice,
		Status:      param.Status,
		AddressId:   address.GetId(param.UserId),
		UserId:      param.UserId,
		NickName:    param.NickName,
		Created:     common.NowTime(),
	}
	return global.Db.Create(&order).RowsAffected
}

func (o *WebOrderService) Delete(param models.WebOrderDeleteParam) int64 {
	return global.Db.Delete(&models.Order{}, param.Id).RowsAffected
}

func (o *WebOrderService) Update(param models.WebOrderUpdateParam) int64 {
	order := models.Order{
		Id:      param.Id,
		Status:  param.Status,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&order).Updates(order).RowsAffected
}

func (o *WebOrderService) GetList(param models.WebOrderListParam) ([]models.WebOrderList, int64) {
	orderList := make([]models.WebOrderList, 0)
	query := &models.Order{
		Id:     param.Id,
		Status: param.Status,
	}
	rows := common.RestPage(param.Page, "order", query, &orderList, &[]models.Order{})
	return orderList, rows
}

func (o *WebOrderService) GetInfo(param models.WebOrderInfoParam) (od models.WebOrderInfo) {
	var order models.Order
	var address models.Address
	var productItem []models.WebProductItem

	global.Db.First(&order, param.Id)
	global.Db.First(&address, order.AddressId)

	idList := strings.Split(order.ProductItem, ",")
	productIdList := make([]uint64, 0)
	for _, id := range idList {
		pid, _ := strconv.Atoi(id)
		if pid != 0 {
			productIdList = append(productIdList, uint64(pid))
		}
	}
	global.Db.Table("product").Find(&productItem, productIdList)
	orderInfo := models.WebOrderInfo{
		Id:              order.Id,
		Created:         order.Created,
		NickName:        order.NickName,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		Name:            address.Name,
		Mobile:          address.Mobile,
		PostalCode:      address.PostalCode,
		Province:        address.Province,
		City:            address.City,
		District:        address.District,
		DetailedAddress: address.DetailedAddress,
		ProductItem:     productItem,
	}
	return orderInfo
}
