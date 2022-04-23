package models

type Order struct {
	Id          uint64  `gorm:"primaryKey"`
	ProductItem string  `gorm:"product_item"`
	TotalPrice  float64 `gorm:"total_price"`
	Status      string  `gorm:"status"`
	AddressId   uint64  `gorm:"address_id"`
	UserId      string  `gorm:"user_id"`
	NickName    string  `gorm:"nick_name"`
	Created     string  `gorm:"created"`
	Updated     string  `gorm:"updated"`
}

type WebOrderCreateParam struct {
	UserId   string `form:"userId"    json:"userId"`
	NickName string `form:"nickName"  json:"nickName"`
	Status   string `form:"status"    json:"status"`
}

type WebOrderDeleteParam struct {
	Id uint64 `form:"id"`
}

type WebOrderUpdateParam struct {
	Id     uint64 `json:"id"`
	Status string `json:"status"`
}

type WebOrderListParam struct {
	Page   Page
	Id     uint64 `form:"id"`
	Status string `form:"status"`
}

type WebOrderInfoParam struct {
	Id uint64 `form:"id"`
}

type WebOrderList struct {
	Id         uint64  `json:"id"`
	NickName   string  `json:"nickName"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`
	Created    string  `json:"created"`
}

type WebOrderInfo struct {
	Id         uint64  `json:"id"`
	Created    string  `json:"created"`
	NickName   string  `json:"nickName"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`

	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	PostalCode      int    `json:"postalCode"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`

	ProductItem []WebProductItem `json:"productItem"`
}

type WebOrderQueryParam struct {
	UserId string `form:"userId"    json:"userId"`
	Status string `form:"status"    json:"status"`
}

// 微信小程序，订单列表传输模型
type WebOrderListInfo struct {
	Id          uint64           `json:"id"`
	Status      string           `json:"status"`
	TotalPrice  float64          `json:"totalPrice"`
	ProductItem []WebProductItem `json:"productItem"`
}
