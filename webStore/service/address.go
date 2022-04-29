package service

import (
	"fmt"
	"mall/common"
	"mall/global"
	"mall/models"
)

type AddressService struct{}

func (a *AddressService) Add(param models.WebAddressAddParam) int64 {
	address := models.Address{
		UserId:          param.UserId,
		Name:            param.Name,
		Mobile:          param.Mobile,
		PostalCode:      param.PostalCode,
		Province:        param.Province,
		City:            param.City,
		District:        param.District,
		DetailedAddress: param.DetailedAddress,
		IsDefault:       param.IsDefault,
		Created:         common.NowTime(),
	}
	if address.IsDefault == 1 {
		var addressId uint
		row := global.Db.Table("address").Select("id").
			Where("is_default = ? and user_id = ?", 1, address.UserId).Take(&addressId).RowsAffected
		if row > 0 {
			global.Db.Table("address").Where("id = ?", addressId).
				Update("is_default", 2)
			return global.Db.Create(&address).RowsAffected
		}
		return global.Db.Create(&address).RowsAffected
	}
	return global.Db.Create(&address).RowsAffected
}

func (a *AddressService) Delete(id uint) int64 {
	return global.Db.Delete(&models.Address{}, id).RowsAffected
}

func (a *AddressService) Update(param models.WebAddressUpdateParam) int64 {
	address := models.Address{
		Id:              param.Id,
		UserId:          param.UserId,
		Name:            param.Name,
		Mobile:          param.Mobile,
		PostalCode:      param.PostalCode,
		Province:        param.Province,
		City:            param.City,
		District:        param.District,
		DetailedAddress: param.DetailedAddress,
		IsDefault:       param.IsDefault,
		Updated:         common.NowTime(),
	}
	if address.IsDefault == 1 {
		var addressId uint
		row := global.Db.Table("address").Select("id").
			Where("is_default = ? and user_id = ?", 1, address.UserId).Take(&addressId).RowsAffected
		fmt.Println(addressId)
		if row > 0 {
			global.Db.Table("address").Where("id = ?", addressId).
				Update("is_default", 2)
			return global.Db.Updates(&address).RowsAffected
		}
		return global.Db.Updates(&address).RowsAffected
	}
	return global.Db.Updates(&address).RowsAffected
}

func (a *AddressService) GetId(uid string) uint64 {
	var id uint64
	global.Db.Table("address").Select("id").
		Where("is_default = ? and user_id = ?", 1, uid).Take(&id)
	return id
}

func (a *AddressService) GetInfo(id uint) models.WebAddressUpdateInfo {
	var updateInfo models.WebAddressUpdateInfo
	global.Db.Table("address").First(&updateInfo, id)
	return updateInfo
}

func (a *AddressService) GetList(uid string) []models.WebAddressList {
	aList := make([]models.WebAddressList, 0)
	global.Db.Table("address").Where("user_id = ?", uid).Order("is_default").Find(&aList)
	return aList
}
