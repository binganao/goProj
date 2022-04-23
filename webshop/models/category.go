package models

//UNCHANGED
type Category struct {
	Id       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	ParentId uint64 `gorm:"parent_id"`
	Level    uint   `gorm:"level"`
	Sort     uint   `gorm:"sort"`
	Created  string `gorm:"created"`
	Updated  string `gorm:"updated"`
}

type WebCategoryCreateParam struct {
	Name     string `json:"name"     binding:"required"`
	ParentId uint64 `json:"parentId" binding:"required,gt=0"`
	Level    uint   `json:"level"    binding:"required,oneof=1 2 3"`
	Sort     uint   `json:"sort"     binding:"required,gt=0"`
}

type WebCategoryDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

type WebCategoryUpdateParam struct {
	Id   uint64 `json:"id"       binding:"required,gt=0"`
	Name string `json:"name"     binding:"required"`
	Sort uint   `json:"sort"     binding:"required,gt=0"`
}

type WebCategoryQueryParam struct {
	Page     Page
	Id       uint64 `form:"id"       binding:"omitempty,gt=0"`
	Name     string `form:"name"     binding:"omitempty"`
	ParentId uint64 `form:"parentId" binding:"omitempty,gt=0"`
	Level    uint   `form:"level"    binding:"omitempty,oneof=1 2 3"`
}

type WebCategoryList struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	ParentId uint64 `json:"parentId"`
	Level    uint   `json:"level"`
	Sort     uint   `json:"sort"`
}

type WebCategoryOption struct {
	Value    uint64              `json:"value"`
	Label    string              `json:"label"`
	Children []WebCategoryOption `json:"children"`
}
