package service

import (
	"mall/common"
	"mall/global"
	"mall/models"
)

type WebCategoryService struct{}

func (c *WebCategoryService) Create(param models.WebCategoryCreateParam) uint64 {
	var category models.Category
	result := global.Db.Where("name = ?", param.Name).First(&category)
	if result.RowsAffected > 0 {
		return category.Id
	}
	category = models.Category{
		Name:     param.Name,
		ParentId: param.ParentId,
		Level:    param.Level,
		Sort:     param.Sort,
		Created:  common.NowTime(),
	}
	global.Db.Create(&category)
	return category.Id
}

func (c *WebCategoryService) Delete(param models.WebCategoryDeleteParam) int64 {
	var pid2, pid3 models.Category
	global.Db.Where("parent_id = ?", param.Id).First(&pid2)
	global.Db.Where("parent_id = ?", pid3.Id).First(&pid3)
	//TODO ???why
	return global.Db.Delete(&models.Category{}, []uint64{param.Id, pid2.Id, pid3.Id}).RowsAffected
}

func (c *WebCategoryService) Update(param models.WebCategoryUpdateParam) int64 {
	category := models.Category{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&category).Updates(category).RowsAffected
}

func (c *WebCategoryService) GetList(param models.WebCategoryQueryParam) ([]models.WebCategoryList, int64) {
	list := make([]models.WebCategoryList, 0)
	query := &models.Category{
		Id:       param.Id,
		Name:     param.Name,
		Level:    param.Level,
		ParentId: param.ParentId,
	}
	rows := common.RestPage(param.Page, "category", query, &list, &[]models.Category{})
	return list, rows
}

func (c *WebCategoryService) GetOption() []models.WebCategoryOption {
	selectList := make([]models.WebCategoryList, 0)
	global.Db.Table("category").Find(&selectList)
	return getTreeOptions(1, selectList)
}

func getTreeOptions(id uint64, list []models.WebCategoryList) []models.WebCategoryOption {
	optionList := make([]models.WebCategoryOption, 0)
	for _, opt := range list {
		if opt.ParentId == id && (opt.Level == 1 || opt.Level == 2) {
			option := models.WebCategoryOption{
				Value:    opt.Id,
				Label:    opt.Name,
				Children: getTreeOptions(opt.Id, list),
			}
			if opt.Level == 2 {
				option.Children = nil
			}
			optionList = append(optionList, option)
		}
	}
	return optionList
}
