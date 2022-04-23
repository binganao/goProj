package models

type WebCartAddParam struct {
	ProductId    uint   `form:"productId"`
	ProductCount uint   `form:"productCount"`
	UserId       string `form:"userId"`
}

type WebCartDeleteParam struct {
	ProductId string `form:"productId"`
	UserId    string `form:"userId"`
}

type WebCartClearParam struct {
	UserId string `form:"userId"`
}

type WebCartQueryParam struct {
	ProductId uint   `form:"productId"`
	UserId    string `form:"userId"`
}

type WebCartItem struct {
	Id        uint64  `gorm:"primaryKey" json:"id"`
	MainImage string  `gorm:"image_url"  json:"mainImage"`
	Title     string  `gorm:"title"      json:"title"`
	Price     float64 `gorm:"price"      json:"price"`
	Count     int     `gorm:"count"      json:"count"`
}

type WebCartInfo struct {
	CartItem   []WebCartItem `json:"cartItem"`
	TotalPrice float64       `json:"totalPrice"`
}
