package models

type Address struct {
	Id              uint   `gorm:"primaryKey"`
	UserId          string `gorm:"user_id"`
	Name            string `gorm:"name"`
	Mobile          string `gorm:"mobile"`
	PostalCode      int    `gorm:"postal_code"`
	Province        string `gorm:"province"`
	City            string `gorm:"city"`
	District        string `gorm:"district"`
	DetailedAddress string `gorm:"detailed_address"`
	IsDefault       int    `gorm:"is_default"`
	Created         string `gorm:"created"`
	Updated         string `gorm:"updated"`
}

type WebAddressAddParam struct {
	UserId          string `form:"userId"`
	Name            string `form:"name"`
	Mobile          string `form:"mobile"`
	PostalCode      int    `form:"postalCode"`
	Province        string `form:"province"`
	City            string `form:"city"`
	District        string `form:"district"`
	DetailedAddress string `form:"detailedAddress"`
	IsDefault       int    `form:"isDefault"`
}

type WebAddressUpdateParam struct {
	Id              uint   `form:"id"`
	UserId          string `form:"userId"`
	Name            string `form:"name"`
	Mobile          string `form:"mobile"`
	PostalCode      int    `form:"postalCode"`
	Province        string `form:"province"`
	City            string `form:"city"`
	District        string `form:"district"`
	DetailedAddress string `form:"detailedAddress"`
	IsDefault       int    `form:"isDefault"`
}

type WebAddressDeleteParam struct {
	AddressId uint `form:"addressId"`
}

type WebAddressInfoParam struct {
	AddressId uint `form:"addressId"`
}

type WebAddressListParam struct {
	UserId string `form:"userId" json:"userId"`
}

type WebAddressList struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	PostalCode      int    `json:"postalCode"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`
	IsDefault       int    `json:"isDefault"`
}

type WebAddressUpdateInfo struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	PostalCode      int    `json:"postalCode"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`
	IsDefault       int    `json:"isDefault"`
}
