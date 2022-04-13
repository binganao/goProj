package service

import (
	"log"
	"mall/common"
	"mall/global"
	"mall/models"
	"strconv"
)

type WebProductService struct{}

func (p *WebProductService) Create(param models.WebProductCreateParam) int64 {
	product := models.Product{
		CategoryId:        param.CategoryId,
		Title:             param.Title,
		Description:       param.Description,
		Price:             param.Price,
		Amount:            param.Amount,
		MainImage:         param.MainImage,
		Delivery:          param.Delivery,
		Assurance:         param.Assurance,
		Name:              param.Name,
		Weight:            param.Weight,
		Brand:             param.Brand,
		Origin:            param.Origin,
		ShelfLife:         param.ShelfLife,
		NetWeight:         param.NetWeight,
		UseWay:            param.UseWay,
		PackingWay:        param.PackingWay,
		StorageConditions: param.StorageConditions,
		DetailImage:       param.DetailImage,
		Status:            param.Status,
		Created:           common.NowTime(),
	}
	rows := global.Db.Create(&product).RowsAffected
	records := global.Db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(product.Id, 10)
		result, err := global.Es.Index().Index("product").Id(id).BodyJson(product).Do(ctx)
		if err != nil {
			log.Println(err)
		}
		return result.PrimaryTerm
	}
	return rows
}

func (p *WebProductService) Delete(param models.WebProductDeleteParam) int64 {
	rows := global.Db.Delete(&models.Product{}, param.Id).RowsAffected
	if rows > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Delete().Index("product").Id(id).Do(ctx)
		if err != nil {
			log.Println(err)
		}
	}
	return rows
}

func (p *WebProductService) Update(param models.WebProductUpdateParam) int64 {
	product := models.Product{
		Id:                param.Id,
		CategoryId:        param.CategoryId,
		Title:             param.Title,
		Description:       param.Description,
		Price:             param.Price,
		Amount:            param.Amount,
		MainImage:         param.MainImage,
		Delivery:          param.Delivery,
		Assurance:         param.Assurance,
		Name:              param.Name,
		Weight:            param.Weight,
		Brand:             param.Brand,
		Origin:            param.Origin,
		ShelfLife:         param.ShelfLife,
		NetWeight:         param.NetWeight,
		UseWay:            param.UseWay,
		PackingWay:        param.PackingWay,
		StorageConditions: param.StorageConditions,
		DetailImage:       param.DetailImage,
		Status:            param.Status,
		Updated:           common.NowTime(),
	}
	rows := global.Db.Model(&product).Updates(product).RowsAffected
	records := global.Db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Update().Index("product").Id(id).Doc(product).Do(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return rows
}

func (p *WebProductService) UpdateStatus(param models.WebProductStatusUpdateParam) int64 {
	product := models.Product{
		Id:     param.Id,
		Status: param.Status,
	}
	rows := global.Db.Model(&product).Update("status", product.Status).RowsAffected
	records := global.Db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Update().Index("product").Id(id).Doc(product).Do(ctx)
		if err != nil {
			log.Println(err)
		}
	}
	return rows
}

// GetInfo 后台管理前端，获取商品信息
func (p *WebProductService) GetInfo(param models.WebProductInfoParam) models.WebProductInfo {
	var product models.WebProductInfo
	global.Db.Table("product").First(&product, param.Id)
	return product
}

// GetList 后台管理前端，获取商品列表
func (p *WebProductService) GetList(param models.WebProductListParam) ([]models.WebProductList, int64) {
	query := &models.Product{
		Id:         param.Id,
		CategoryId: param.CategoryId,
		Title:      param.Title,
		Status:     param.Status,
	}
	productList := make([]models.WebProductList, 0)
	rows := common.RestPage(param.Page, "product", query, &productList, &[]models.Product{})
	return productList, rows
}
